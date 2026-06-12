package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/kratosx/pkg/value"

	pb "github.com/limes-cloud/configure/api/configure"
	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/service"
	"github.com/limes-cloud/configure/internal/infra/dbs"
	"github.com/limes-cloud/configure/internal/types"
)

type Configure struct {
	pb.UnimplementedConfigureServer
	srv *service.Configure
}

func NewConfigure() *Configure {
	return &Configure{
		srv: service.NewConfigure(
			dbs.NewConfigure(),
			dbs.NewEnv(),
			dbs.NewBusiness(),
			dbs.NewResource(),
			dbs.NewTemplate(),
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewConfigure()
		pb.RegisterConfigureHTTPServer(hs, srv)
		pb.RegisterConfigureServer(gs, srv)
	})
}

func (s *Configure) CompareConfigure(c context.Context, in *pb.CompareConfigureRequest) (*pb.CompareConfigureResponse, error) {
	list, err := s.srv.CompareConfigure(core.MustContext(c), &types.CompareConfigureRequest{
		EnvId: in.EnvId,
		App:   in.App,
	})
	if err != nil {
		return nil, err
	}
	reply := pb.CompareConfigureResponse{}
	if err := value.Transform(list, &reply.List); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Configure) GetConfigure(ctx context.Context, in *pb.GetConfigureRequest) (*pb.GetConfigureResponse, error) {
	res, err := s.srv.GetConfigureByEnvAndApp(core.MustContext(ctx), in.EnvId, in.App)
	if err != nil {
		return nil, err
	}
	reply := pb.GetConfigureResponse{}
	if err := value.Transform(res, &reply); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *Configure) UpdateConfigure(ctx context.Context, in *pb.UpdateConfigureRequest) (*pb.UpdateConfigureResponse, error) {
	return &pb.UpdateConfigureResponse{}, s.srv.UpdateConfigure(core.MustContext(ctx), &entity.Configure{
		App:         in.App,
		EnvId:       in.EnvId,
		Description: &in.Description,
	})
}

func (s *Configure) WatchConfigure(in *pb.WatchConfigureRequest, reply pb.Configure_WatchConfigureServer) error {
	return s.srv.Watch(core.MustContext(reply.Context()), &types.WatcherConfigRequest{
		App:   in.App,
		Token: in.Token,
	}, func(data *types.WatcherConfigureReply) error {
		return reply.Send(&pb.WatchConfigureResponse{
			Format:  data.Format,
			Content: data.Content,
		})
	})
}
