package service

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/internal/data"
)

type Service struct {
	v1.UnimplementedServiceServer
	UserUseCase      *biz.UserUseCase
	EnvUseCase       *biz.EnvUseCase
	ServerUseCase    *biz.ServerUseCase
	ResourceUseCase  *biz.ResourceUseCase
	BusinessUseCase  *biz.BusinessUseCase
	TemplateUseCase  *biz.TemplateUseCase
	ConfigureUseCase *biz.ConfigureUseCase
}

func New(config *config.Config) *Service {
	envRepo := data.NewEnvRepo()
	serverRepo := data.NewServerRepo()
	bsRepo := data.NewBusinessRepo()
	rsRepo := data.NewResourceRepo()
	tpRepo := data.NewTemplateRepo()
	cfRepo := data.NewConfigureRepo()
	return &Service{
		UserUseCase:      biz.NewUserUseCase(config),
		EnvUseCase:       biz.NewEnvUseCase(config, envRepo),
		ServerUseCase:    biz.NewServerUseCase(config, serverRepo),
		BusinessUseCase:  biz.NewBusinessUseCase(config, bsRepo),
		ResourceUseCase:  biz.NewResourceUseCase(config, rsRepo),
		ConfigureUseCase: biz.NewConfigureUseCase(config, cfRepo),
		TemplateUseCase:  biz.NewTemplateUseCase(config, tpRepo, rsRepo, bsRepo),
	}
}
