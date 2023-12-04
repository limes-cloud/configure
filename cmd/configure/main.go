package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/limes-cloud/configure/api/v1"
	systemConfig "github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/handler"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
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

	app := kratosx.New(
		kratosx.Config(file.NewSource("config/config.yaml")),
		kratosx.RegistrarServer(RegisterServer),
	)

	if err := app.Run(); err != nil {
		log.Fatal("run service fail", err.Error())
	}
}

func RegisterServer(c config.Config, hs *http.Server, gs *grpc.Server) {
	conf := &systemConfig.Config{}

	// 配置初始化
	if err := c.Value("business").Scan(conf); err != nil {
		panic("author config format error:" + err.Error())
	}

	// 监听服务
	c.Watch("business", func(value config.Value) {
		if err := value.Scan(conf); err != nil {
			log.Printf("business配置变更失败：%s", err.Error())
		}
	})

	go RegisterWebServer(conf)
	srv := handler.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)
}

func RegisterWebServer(config *systemConfig.Config) {
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
