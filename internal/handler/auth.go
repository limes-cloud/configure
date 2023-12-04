package handler

import (
	"context"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	return s.Auth.Login(kratosx.MustContext(ctx), in)
}

func (s *Service) RefreshToken(ctx context.Context, in *emptypb.Empty) (*v1.RefreshTokenReply, error) {
	return s.Auth.RefreshToken(kratosx.MustContext(ctx))
}
