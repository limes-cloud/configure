package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "github.com/limes-cloud/configure/api/configure/configure/v1"
	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/service"
	"github.com/limes-cloud/configure/internal/infra/dbs"
	"github.com/limes-cloud/configure/internal/infra/rpc"
	"github.com/limes-cloud/configure/internal/types"
)

type ConfigureApp struct {
	pb.UnimplementedConfigureServer
	srv *service.ConfigureService
}

func NewConfigureApp(conf *conf.Config) *ConfigureApp {
	return &ConfigureApp{
		srv: service.NewConfigureService(
			conf,
			dbs.NewConfigureInfra(),
			dbs.NewServerInfra(),
			dbs.NewEnvInfra(),
			dbs.NewBusinessInfra(),
			dbs.NewResourceInfra(),
			dbs.NewTemplateInfra(),
			rpc.NewPermissionInfra(),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewConfigureApp(c)
		pb.RegisterConfigureHTTPServer(hs, srv)
		pb.RegisterConfigureServer(gs, srv)
	})
}

func (s *ConfigureApp) CompareConfigure(c context.Context, in *pb.CompareConfigureRequest) (*pb.CompareConfigureReply, error) {
	list, err := s.srv.CompareConfigure(kratosx.MustContext(c), &types.CompareConfigureRequest{
		EnvId:    in.EnvId,
		ServerId: in.ServerId,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.CompareConfigureReply{}
	if err := valx.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *ConfigureApp) GetConfigure(ctx context.Context, in *pb.GetConfigureRequest) (*pb.GetConfigureReply, error) {
	res, err := s.srv.GetConfigureByEnvAndSrv(kratosx.MustContext(ctx), in.EnvId, in.ServerId)
	if err != nil {
		return nil, err
	}
	reply := pb.GetConfigureReply{}
	if err := valx.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *ConfigureApp) UpdateConfigure(ctx context.Context, in *pb.UpdateConfigureRequest) (*pb.UpdateConfigureReply, error) {
	return &pb.UpdateConfigureReply{}, s.srv.UpdateConfigure(kratosx.MustContext(ctx), &entity.Configure{
		ServerId:    in.ServerId,
		EnvId:       in.EnvId,
		Description: &in.Description,
	})
}

func (s *ConfigureApp) WatchConfigure(in *pb.WatchConfigureRequest, reply pb.Configure_WatchConfigureServer) error {
	return s.srv.Watch(kratosx.MustContext(reply.Context()), &types.WatcherConfigRequest{
		Server: in.Server,
		Token:  in.Token,
	}, func(data *types.WatcherConfigureReply) error {
		return reply.Send(&pb.WatchConfigureReply{
			Format:  data.Format,
			Content: data.Content,
		})
	})
}
