package service

import (
	"context"

	v1 "github.com/limes-cloud/configure/api/v1"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	token, err := s.UserUseCase.Login(kratosx.MustContext(ctx), in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{Token: token}, nil
}

func (s *Service) RefreshToken(ctx context.Context, in *emptypb.Empty) (*v1.RefreshTokenReply, error) {
	token, err := s.UserUseCase.RefreshToken(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	return &v1.RefreshTokenReply{Token: token}, nil
}
