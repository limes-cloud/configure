package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "github.com/limes-cloud/configure/api/configure/business/v1"
	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/biz/business"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/data"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer
	uc *business.UseCase
}

func NewBusinessService(conf *conf.Config) *BusinessService {
	return &BusinessService{
		uc: business.NewUseCase(conf, data.NewBusinessRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewBusinessService(c)
		pb.RegisterBusinessHTTPServer(hs, srv)
		pb.RegisterBusinessServer(gs, srv)
	})
}

// GetBusiness 获取指定的业务配置信息
func (s *BusinessService) GetBusiness(c context.Context, req *pb.GetBusinessRequest) (*pb.GetBusinessReply, error) {
	var (
		in  = business.GetBusinessRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetBusiness(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetBusinessReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListBusiness 获取业务配置信息列表
func (s *BusinessService) ListBusiness(c context.Context, req *pb.ListBusinessRequest) (*pb.ListBusinessReply, error) {
	var (
		in  = business.ListBusinessRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListBusiness(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListBusinessReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateBusiness 创建业务配置信息
func (s *BusinessService) CreateBusiness(c context.Context, req *pb.CreateBusinessRequest) (*pb.CreateBusinessReply, error) {
	var (
		in  = business.Business{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateBusiness(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBusinessReply{Id: id}, nil
}

// UpdateBusiness 更新业务配置信息
func (s *BusinessService) UpdateBusiness(c context.Context, req *pb.UpdateBusinessRequest) (*pb.UpdateBusinessReply, error) {
	var (
		in  = business.Business{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateBusiness(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateBusinessReply{}, nil
}

// DeleteBusiness 删除业务配置信息
func (s *BusinessService) DeleteBusiness(c context.Context, req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessReply, error) {
	total, err := s.uc.DeleteBusiness(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBusinessReply{Total: total}, nil
}

// ListBusinessValue 获取业务配置值信息列表
func (s *BusinessService) ListBusinessValue(c context.Context, req *pb.ListBusinessValueRequest) (*pb.ListBusinessValueReply, error) {
	var (
		ctx = kratosx.MustContext(c)
	)
	result, total, err := s.uc.ListBusinessValue(kratosx.MustContext(c), req.BusinessId)
	if err != nil {
		return nil, err
	}

	reply := pb.ListBusinessValueReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// UpdateBusinessValue 更新业务配置值信息
func (s *BusinessService) UpdateBusinessValue(c context.Context, req *pb.UpdateBusinessValueRequest) (*pb.UpdateBusinessValueReply, error) {
	var (
		in  []*business.BusinessValue
		ctx = kratosx.MustContext(c)
	)

	for _, item := range req.List {
		in = append(in, &business.BusinessValue{
			BusinessId: req.BusinessId,
			EnvId:      item.EnvId,
			Value:      item.Value,
		})
	}

	if err := s.uc.UpdateBusinessValue(ctx, in); err != nil {
		return nil, err
	}

	return &pb.UpdateBusinessValueReply{}, nil
}
