package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	thttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/manager/api/authorize"
	_ "go.uber.org/automaxprocs"

	"github.com/limes-cloud/configure/internal/app"
)

const (
	AppName = "Configure"
)

func main() {
	server := kratosx.New(
		kratosx.WithRegistrarServer(func(hs *thttp.Server, gs *grpc.Server) {
			app.New(hs, gs)
		}),
		kratosx.WithMiddleware(authorize.AuthMiddleware()),
		kratosx.WithKratosOptions(kratos.AfterStart(func(_ context.Context) error {
			fmt.Printf("Hello %s !\n", AppName)
			return nil
		})),
	)

	if err := server.App().Run(); err != nil {
		log.Fatal("run service fail", err.Error())
	}
}
