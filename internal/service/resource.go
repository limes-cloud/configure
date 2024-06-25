package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/configure/api/configure/errors"
	pb "github.com/limes-cloud/configure/api/configure/resource/v1"
	"github.com/limes-cloud/configure/internal/biz/resource"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/data"
)

type ResourceService struct {
	pb.UnimplementedResourceServer
	uc *resource.UseCase
}

func NewResourceService(conf *conf.Config) *ResourceService {
	return &ResourceService{
		uc: resource.NewUseCase(conf, data.NewResourceRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewResourceService(c)
		pb.RegisterResourceHTTPServer(hs, srv)
		pb.RegisterResourceServer(gs, srv)
	})
}

// GetResource 获取指定的资源配置信息
func (s *ResourceService) GetResource(c context.Context, req *pb.GetResourceRequest) (*pb.GetResourceReply, error) {
	var (
		in  = resource.GetResourceRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetResource(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetResourceReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListResource 获取资源配置信息列表
func (s *ResourceService) ListResource(c context.Context, req *pb.ListResourceRequest) (*pb.ListResourceReply, error) {
	var (
		in  = resource.ListResourceRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListResource(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListResourceReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateResource 创建资源配置信息
func (s *ResourceService) CreateResource(c context.Context, req *pb.CreateResourceRequest) (*pb.CreateResourceReply, error) {
	var (
		in  = resource.Resource{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	for _, id := range req.ServerIds {
		in.ResourceServers = append(in.ResourceServers, &resource.ResourceServer{
			ServerId: id,
		})
	}

	id, err := s.uc.CreateResource(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResourceReply{Id: id}, nil
}

// UpdateResource 更新资源配置信息
func (s *ResourceService) UpdateResource(c context.Context, req *pb.UpdateResourceRequest) (*pb.UpdateResourceReply, error) {
	var (
		in  = resource.Resource{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	for _, id := range req.ServerIds {
		in.ResourceServers = append(in.ResourceServers, &resource.ResourceServer{
			ServerId: id,
		})
	}

	if err := s.uc.UpdateResource(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateResourceReply{}, nil
}

// DeleteResource 删除资源配置信息
func (s *ResourceService) DeleteResource(c context.Context, req *pb.DeleteResourceRequest) (*pb.DeleteResourceReply, error) {
	total, err := s.uc.DeleteResource(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResourceReply{Total: total}, nil
}

// ListResourceValue 获取业务配置值信息列表
func (s *ResourceService) ListResourceValue(c context.Context, req *pb.ListResourceValueRequest) (*pb.ListResourceValueReply, error) {
	var (
		ctx = kratosx.MustContext(c)
	)
	result, total, err := s.uc.ListResourceValue(kratosx.MustContext(c), req.ResourceId)
	if err != nil {
		return nil, err
	}

	reply := pb.ListResourceValueReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// UpdateResourceValue 更新业务配置值信息
func (s *ResourceService) UpdateResourceValue(c context.Context, req *pb.UpdateResourceValueRequest) (*pb.UpdateResourceValueReply, error) {
	var (
		in  []*resource.ResourceValue
		ctx = kratosx.MustContext(c)
	)

	for _, item := range req.List {
		in = append(in, &resource.ResourceValue{
			ResourceId: req.ResourceId,
			EnvId:      item.EnvId,
			Value:      item.Value,
		})
	}

	if err := s.uc.UpdateResourceValue(ctx, in); err != nil {
		return nil, err
	}

	return &pb.UpdateResourceValueReply{}, nil
}
