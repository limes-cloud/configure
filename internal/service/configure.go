package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "github.com/limes-cloud/configure/api/configure/configure/v1"
	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/biz/configure"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/data"
	"github.com/limes-cloud/configure/internal/factory"
)

type ConfigureService struct {
	pb.UnimplementedConfigureServer
	uc *configure.UseCase
}

func NewConfigureService(conf *conf.Config) *ConfigureService {
	return &ConfigureService{
		uc: configure.NewUseCase(
			conf,
			data.NewConfigureRepo(),
			factory.New(data.NewBusinessRepo(), data.NewResourceRepo()),
		),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewConfigureService(c)
		pb.RegisterConfigureHTTPServer(hs, srv)
		pb.RegisterConfigureServer(gs, srv)
	})
}

func (s *ConfigureService) CompareConfigure(c context.Context, in *pb.CompareConfigureRequest) (*pb.CompareConfigureReply, error) {
	list, err := s.uc.CompareConfigure(kratosx.MustContext(c), &configure.CompareConfigureRequest{
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

func (s *ConfigureService) GetConfigure(ctx context.Context, in *pb.GetConfigureRequest) (*pb.GetConfigureReply, error) {
	res, err := s.uc.GetConfigureByEnvAndSrv(kratosx.MustContext(ctx), in.EnvId, in.ServerId)
	if err != nil {
		return nil, err
	}
	reply := pb.GetConfigureReply{}
	if err := valx.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *ConfigureService) UpdateConfigure(ctx context.Context, in *pb.UpdateConfigureRequest) (*pb.UpdateConfigureReply, error) {
	return &pb.UpdateConfigureReply{}, s.uc.UpdateConfigure(kratosx.MustContext(ctx), &configure.Configure{
		ServerId:    in.ServerId,
		EnvId:       in.EnvId,
		Description: &in.Description,
	})
}

func (s *ConfigureService) WatchConfigure(in *pb.WatchConfigureRequest, reply pb.Configure_WatchConfigureServer) error {
	return s.uc.Watch(kratosx.MustContext(reply.Context()), &configure.WatcherConfigRequest{
		Server: in.Server,
		Token:  in.Token,
	}, func(data *configure.WatcherConfigureReply) error {
		return reply.Send(&pb.WatchConfigureReply{
			Format:  data.Format,
			Content: data.Content,
		})
	})
}
