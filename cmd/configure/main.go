package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	thttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/config"
	"github.com/limes-cloud/kratosx/pkg/printx"
	"github.com/limes-cloud/kratosx/pkg/webserver"
	_ "go.uber.org/automaxprocs"

	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/service"
)

const (
	AppName = "Configure"
)

func main() {
	app := kratosx.New(
		kratosx.RegistrarServer(RegisterServer),
		kratosx.Options(kratos.AfterStart(func(_ context.Context) error {
			printx.ArtFont(fmt.Sprintf("Hello %s !", AppName))
			return nil
		})),
	)

	if err := app.Run(); err != nil {
		log.Fatal("run service fail", err.Error())
	}
}

func RegisterServer(c config.Config, hs *thttp.Server, gs *grpc.Server) {
	cfg := &conf.Config{}

	// 监听服务
	c.ScanWatch("business", func(value config.Value) {
		if err := value.Scan(cfg); err != nil {
			log.Printf("business配置变更失败：%s", err.Error())
		}
	})
	go webserver.Run("web/dist/", cfg.WebUI.Addr, map[string]any{
		"Port": c.App().Server.Http.Port,
	})
	service.New(cfg, hs, gs)
}
