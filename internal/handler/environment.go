package handler

import (
	"context"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) AllEnvironment(ctx context.Context, in *emptypb.Empty) (*v1.AllEnvironmentReply, error) {
	return s.Environment.All(kratosx.MustContext(ctx), in)
}

func (s *Service) AddEnvironment(ctx context.Context, in *v1.AddEnvironmentRequest) (*emptypb.Empty, error) {
	return s.Environment.Add(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateEnvironment(ctx context.Context, in *v1.UpdateEnvironmentRequest) (*emptypb.Empty, error) {
	return s.Environment.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) DeleteEnvironment(ctx context.Context, in *v1.DeleteEnvironmentRequest) (*emptypb.Empty, error) {
	return s.Environment.Delete(kratosx.MustContext(ctx), in)
}

func (s *Service) GetEnvironmentToken(ctx context.Context, in *v1.GetEnvironmentTokenRequest) (*v1.GetEnvironmentTokenReply, error) {
	return s.Environment.GetToken(kratosx.MustContext(ctx), in)
}

func (s *Service) ResetEnvironmentToken(ctx context.Context, in *v1.ResetEnvironmentTokenRequest) (*v1.ResetEnvironmentTokenReply, error) {
	return s.Environment.ResetToken(kratosx.MustContext(ctx), in)
}
