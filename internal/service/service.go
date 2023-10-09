package service

import (
	"context"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/logic"
	"github.com/limes-cloud/kratos"
)

// Service is a greeter service.
type Service struct {
	v1.UnimplementedServiceServer
	logic *logic.Logic
}

func New(conf *config.Config) *Service {
	return &Service{
		logic: logic.NewLogic(conf),
	}
}

// AllEnvironment 获取全部环境
func (s *Service) AllEnvironment(ctx context.Context, in *v1.AllEnvironmentRequest) (*v1.AllEnvironmentResponse, error) {
	return s.logic.AllEnvironment(kratos.MustContext(ctx), in)
}
