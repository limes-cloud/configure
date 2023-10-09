package handler

import (
	"context"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/logic"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *Service) AllEnvironment(ctx context.Context, in *emptypb.Empty) (*v1.AllEnvironmentReply, error) {
	return s.logic.AllEnvironment(kratos.MustContext(ctx), in)
}

func (s *Service) AddEnvironment(ctx context.Context, in *v1.AddEnvironmentRequest) (*emptypb.Empty, error) {
	return s.logic.AddEnvironment(kratos.MustContext(ctx), in)
}

func (s *Service) UpdateEnvironment(ctx context.Context, in *v1.UpdateEnvironmentRequest) (*emptypb.Empty, error) {
	return s.logic.UpdateEnvironment(kratos.MustContext(ctx), in)
}

func (s *Service) DeleteEnvironment(ctx context.Context, in *v1.DeleteEnvironmentRequest) (*emptypb.Empty, error) {
	return s.logic.DeleteEnvironment(kratos.MustContext(ctx), in)
}

func (s *Service) GetEnvironmentToken(ctx context.Context, in *v1.GetEnvironmentTokenRequest) (*v1.GetEnvironmentTokenReply, error) {
	return s.logic.GetEnvironmentToken(kratos.MustContext(ctx), in)
}

func (s *Service) ResetEnvironmentToken(ctx context.Context, in *v1.ResetEnvironmentTokenRequest) (*emptypb.Empty, error) {
	return s.logic.ResetEnvironmentToken(kratos.MustContext(ctx), in)
}
