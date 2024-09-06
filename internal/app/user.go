package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	pb "github.com/limes-cloud/configure/api/configure/user/v1"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/domain/service"
	"github.com/limes-cloud/kratosx"
)

type UserApp struct {
	pb.UnimplementedUserServer
	srv *service.UserService
}

func NewUserApp(conf *conf.Config) *UserApp {
	return &UserApp{
		srv: service.NewUserService(conf),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewUserApp(c)
		pb.RegisterUserHTTPServer(hs, srv)
		pb.RegisterUserServer(gs, srv)
	})
}

func (s *UserApp) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	token, err := s.srv.Login(kratosx.MustContext(ctx), in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: token}, nil
}

func (s *UserApp) RefreshToken(ctx context.Context, _ *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	token, err := s.srv.RefreshToken(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	return &pb.RefreshTokenReply{Token: token}, nil
}
