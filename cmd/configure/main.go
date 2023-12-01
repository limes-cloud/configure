package main

import (
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	v1 "github.com/limes-cloud/configure/api/v1"
	ccf "github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/handler"
	"github.com/limes-cloud/kratos"
	"github.com/limes-cloud/kratos/config"
	"github.com/limes-cloud/kratos/config/file"
	"github.com/limes-cloud/kratos/log"
	"github.com/limes-cloud/kratos/middleware/tracing"
	"github.com/limes-cloud/kratos/transport/grpc"
	thttp "github.com/limes-cloud/kratos/transport/http"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string

	id, _ = os.Hostname()
)

func main() {
	app := kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Config(file.NewSource("config/config.yaml")),
		kratos.RegistrarServer(RegisterServer),
		kratos.LoggerWith(kratos.LogField{
			"id":      id,
			"name":    Name,
			"version": Version,
			"trace":   tracing.TraceID(),
			"span":    tracing.SpanID(),
		}),
	)

	if err := app.Run(); err != nil {
		log.Errorf("run service fail: %v", err)
	}
}

func RegisterServer(hs *thttp.Server, gs *grpc.Server, c config.Config) {
	ccfIns := &ccf.Config{}
	if err := c.ScanKey("business", ccfIns); err != nil {
		panic("author config format error:" + err.Error())
	}
	go RegisterWebServer(ccfIns)
	srv := handler.New(ccfIns)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}

func RegisterWebServer(config *ccf.Config) {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("web/dist/", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile("web/dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				_, _ = c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			_, _ = c.Writer.Write((content))
			c.Writer.Flush()
		}
	})
	if err := r.Run(config.WebUI.Addr); err != nil {
		panic("start web ui error:" + err.Error())
	}
}
