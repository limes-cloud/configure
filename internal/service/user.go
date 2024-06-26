package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"

	pb "github.com/limes-cloud/configure/api/configure/user/v1"
	"github.com/limes-cloud/configure/internal/biz/user"
	"github.com/limes-cloud/configure/internal/conf"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *user.UseCase
}

func NewUserService(conf *conf.Config) *UserService {
	return &UserService{
		uc: user.NewUseCase(conf),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewUserService(c)
		pb.RegisterUserHTTPServer(hs, srv)
		pb.RegisterUserServer(gs, srv)
	})
}

func (s *UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := s.uc.Login(kratosx.MustContext(ctx), in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: token}, nil
}

func (s *UserService) RefreshToken(ctx context.Context, _ *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	token, err := s.uc.RefreshToken(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	return &pb.RefreshTokenReply{Token: token}, nil
}
