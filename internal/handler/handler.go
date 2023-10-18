package handler

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/logic"
)

// Service is a greeter service.
type Service struct {
	v1.UnimplementedServiceServer
	Environment    *logic.Environment
	Server         *logic.Server
	Resource       *logic.Resource
	ResourceServer *logic.ResourceServer
	ResourceValue  *logic.ResourceValue
	Business       *logic.Business
	BusinessValue  *logic.BusinessValue
	Template       *logic.Template
	Configure      *logic.Configure
}

func New(config *config.Config) *Service {
	return &Service{
		Environment:    logic.NewEnvironment(config),
		Server:         logic.NewServer(config),
		Resource:       logic.NewResource(config),
		ResourceServer: logic.NewResourceServer(config),
		ResourceValue:  logic.NewResourceValue(config),
		Business:       logic.NewBusiness(config),
		BusinessValue:  logic.NewBusinessValue(config),
		Template:       logic.NewTemplate(config),
		Configure:      logic.NewConfigure(config),
	}
}
