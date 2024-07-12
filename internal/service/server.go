package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/configure/api/configure/errors"
	pb "github.com/limes-cloud/configure/api/configure/server/v1"
	"github.com/limes-cloud/configure/internal/biz/server"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/data"
)

type ServerService struct {
	pb.UnimplementedServerServer
	uc *server.UseCase
}

func NewServerService(conf *conf.Config) *ServerService {
	return &ServerService{
		uc: server.NewUseCase(conf, data.NewServerRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewServerService(c)
		pb.RegisterServerHTTPServer(hs, srv)
		pb.RegisterServerServer(gs, srv)
	})
}

// ListServer 获取服务信息列表
func (s *ServerService) ListServer(c context.Context, req *pb.ListServerRequest) (*pb.ListServerReply, error) {
	var (
		in  = server.ListServerRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListServer(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListServerReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateServer 创建服务信息
func (s *ServerService) CreateServer(c context.Context, req *pb.CreateServerRequest) (*pb.CreateServerReply, error) {
	var (
		in  = server.Server{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateServer(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateServerReply{Id: id}, nil
}

// UpdateServer 更新服务信息
func (s *ServerService) UpdateServer(c context.Context, req *pb.UpdateServerRequest) (*pb.UpdateServerReply, error) {
	var (
		in  = server.Server{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateServer(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateServerReply{}, nil
}

// DeleteServer 删除服务信息
func (s *ServerService) DeleteServer(c context.Context, req *pb.DeleteServerRequest) (*pb.DeleteServerReply, error) {
	if err := s.uc.DeleteServer(kratosx.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteServerReply{}, nil
}

// UpdateServerStatus 更新服务信息状态
func (s *ServerService) UpdateServerStatus(c context.Context, req *pb.UpdateServerStatusRequest) (*pb.UpdateServerStatusReply, error) {
	return &pb.UpdateServerStatusReply{}, s.uc.UpdateServerStatus(kratosx.MustContext(c), req.Id, req.Status)
}
