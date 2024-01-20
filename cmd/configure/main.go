package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/limes-cloud/configure/internal/initiator"

	"github.com/limes-cloud/configure/internal/service"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	_ "go.uber.org/automaxprocs"

	v1 "github.com/limes-cloud/configure/api/v1"
	systemConfig "github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/consts"
	"github.com/limes-cloud/configure/pkg/pt"
)

const (
	AppName = "Configure"
)

func main() {
	app := kratosx.New(
		kratosx.Config(file.NewSource("config/config.yaml")),
		kratosx.RegistrarServer(RegisterServer),
		kratosx.Options(kratos.AfterStart(func(_ context.Context) error {
			pt.ArtFont(fmt.Sprintf("Hello %s !", AppName))
			return nil
		})),
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

	go RegisterWebServer(c, conf)
	srv := service.New(conf)
	v1.RegisterServiceHTTPServer(hs, srv)
	v1.RegisterServiceServer(gs, srv)

	// 初始化逻辑
	ior := initiator.New(conf)
	if err := ior.Run(srv); err != nil {
		panic("initiator error:" + err.Error())
	}
}

func RegisterWebServer(c config.Config, config *systemConfig.Config) {
	r := gin.Default()

	mode := c.App().Env
	if c.App().Env != consts.ENV_TEST {
		gin.SetMode(gin.ReleaseMode)
	}

	rootPath := fmt.Sprintf("web/dist/%s/", strings.ToLower(mode))
	r.Use(static.Serve("/", static.LocalFile(rootPath, true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile(rootPath + "index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				_, _ = c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			_, _ = c.Writer.Write(content)
			c.Writer.Flush()
		}
	})
	if err := r.Run(config.WebUI.Addr); err != nil {
		panic("start web ui error:" + err.Error())
	}
}
