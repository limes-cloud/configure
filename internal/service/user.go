package service

import (
	"context"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/configure/api/user/v1"
	biz "github.com/limes-cloud/configure/internal/biz/user"
	"github.com/limes-cloud/configure/internal/config"
)

type UserService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase
}

func NewUserService(conf *config.Config) *UserService {
	return &UserService{
		uc: biz.NewUseCase(conf),
	}
}

func (s *UserService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	token, err := s.uc.Login(kratosx.MustContext(ctx), in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	return &v1.LoginReply{Token: token}, nil
}

func (s *UserService) RefreshToken(ctx context.Context, in *emptypb.Empty) (*v1.RefreshTokenReply, error) {
	token, err := s.uc.RefreshToken(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	return &v1.RefreshTokenReply{Token: token}, nil
}
