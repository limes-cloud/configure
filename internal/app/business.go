package app

import (
	"context"

	kmodel "github.com/limes-cloud/kratosx/model"

	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pb "github.com/limes-cloud/configure/api/business"
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/service"
	"github.com/limes-cloud/configure/internal/infra/dbs"
)

type Business struct {
	pb.UnimplementedBusinessServer
	srv *service.Business
}

func NewBusiness() *Business {
	return &Business{
		srv: service.NewBusiness(dbs.NewBusiness()),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewBusiness()
		pb.RegisterBusinessHTTPServer(hs, srv)
		pb.RegisterBusinessServer(gs, srv)
	})
}

// ListBusiness 获取业务配置信息列表
func (s *Business) ListBusiness(c context.Context, req *pb.ListBusinessRequest) (*pb.ListBusinessResponse, error) {
	list, total, err := s.srv.ListBusiness(core.MustContext(c), &types.ListBusinessRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		App:      req.App,
		Keyword:  req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListBusinessResponse{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListBusinessResponse_Business{
			Id:          item.Id,
			App:         item.App,
			Keyword:     item.Keyword,
			Type:        item.Type,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateBusiness 创建业务配置信息
func (s *Business) CreateBusiness(c context.Context, req *pb.CreateBusinessRequest) (*pb.CreateBusinessResponse, error) {
	id, err := s.srv.CreateBusiness(core.MustContext(c), &entity.Business{
		App:         req.App,
		Keyword:     req.Keyword,
		Type:        req.Type,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateBusinessResponse{Id: id}, nil
}

// UpdateBusiness 更新业务配置信息
func (s *Business) UpdateBusiness(c context.Context, req *pb.UpdateBusinessRequest) (*pb.UpdateBusinessResponse, error) {
	if err := s.srv.UpdateBusiness(core.MustContext(c), &entity.Business{
		BaseModel:   kmodel.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Type:        req.Type,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateBusinessResponse{}, nil
}

// DeleteBusiness 删除业务配置信息
func (s *Business) DeleteBusiness(c context.Context, req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessResponse, error) {
	if err := s.srv.DeleteBusiness(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteBusinessResponse{}, nil
}

// ListBusinessValue 获取业务配置值信息列表
func (s *Business) ListBusinessValue(c context.Context, req *pb.ListBusinessValueRequest) (*pb.ListBusinessValueResponse, error) {
	list, err := s.srv.ListBusinessValue(core.MustContext(c), req.BusinessId)
	if err != nil {
		return nil, err
	}
	reply := pb.ListBusinessValueResponse{
		Total: uint32(len(list)),
	}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListBusinessValueResponse_BusinessValue{
			Id:         item.Id,
			EnvId:      item.EnvId,
			BusinessId: item.BusinessId,
			Value:      item.Value,
			CreatedAt:  uint32(item.CreatedAt),
			UpdatedAt:  uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// UpdateBusinessValue 更新业务配置值信息
func (s *Business) UpdateBusinessValue(c context.Context, req *pb.UpdateBusinessValueRequest) (*pb.UpdateBusinessValueResponse, error) {
	var list []*entity.BusinessValue
	for _, item := range req.List {
		list = append(list, &entity.BusinessValue{
			BusinessId: req.BusinessId,
			EnvId:      item.EnvId,
			Value:      item.Value,
		})
	}

	if err := s.srv.UpdateBusinessValue(core.MustContext(c), list); err != nil {
		return nil, err
	}
	return &pb.UpdateBusinessValueResponse{}, nil
}
