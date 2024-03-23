package service

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	businessV1 "github.com/limes-cloud/configure/api/business/v1"
	configureV1 "github.com/limes-cloud/configure/api/configure/v1"
	envV1 "github.com/limes-cloud/configure/api/env/v1"
	resourceV1 "github.com/limes-cloud/configure/api/resource/v1"
	serverV1 "github.com/limes-cloud/configure/api/server/v1"
	templateV1 "github.com/limes-cloud/configure/api/template/v1"
	userV1 "github.com/limes-cloud/configure/api/user/v1"
	"github.com/limes-cloud/configure/internal/config"
)

func New(c *config.Config, hs *http.Server, gs *grpc.Server) {
	envService := NewEnvService(c)
	envV1.RegisterServiceHTTPServer(hs, envService)
	envV1.RegisterServiceServer(gs, envService)

	serverService := NewServerService(c)
	serverV1.RegisterServiceHTTPServer(hs, serverService)
	serverV1.RegisterServiceServer(gs, serverService)

	businessService := NewBusinessService(c)
	businessV1.RegisterServiceHTTPServer(hs, businessService)
	businessV1.RegisterServiceServer(gs, businessService)

	resourceService := NewResourceService(c)
	resourceV1.RegisterServiceHTTPServer(hs, resourceService)
	resourceV1.RegisterServiceServer(gs, resourceService)

	templateService := NewTemplateService(c)
	templateV1.RegisterServiceHTTPServer(hs, templateService)
	templateV1.RegisterServiceServer(gs, templateService)

	configureService := NewConfigureService(c, templateService.uc, envService.uc, serverService.uc)
	configureV1.RegisterServiceHTTPServer(hs, configureService)
	configureV1.RegisterServiceServer(gs, configureService)

	userService := NewUserService(c)
	userV1.RegisterServiceHTTPServer(hs, userService)
	userV1.RegisterServiceServer(gs, userService)
}
