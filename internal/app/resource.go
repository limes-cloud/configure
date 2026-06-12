package app

import (
	"context"

	kmodel "github.com/limes-cloud/kratosx/model"

	"github.com/limes-cloud/configure/internal/types"

	"github.com/limes-cloud/configure/internal/domain/entity"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/configure/api/errors"
	pb "github.com/limes-cloud/configure/api/resource"
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/service"
	"github.com/limes-cloud/configure/internal/infra/dbs"
)

type Resource struct {
	pb.UnimplementedResourceServer
	srv *service.Resource
}

func NewResource() *Resource {
	return &Resource{
		srv: service.NewResource(dbs.NewResource()),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewResource()
		pb.RegisterResourceHTTPServer(hs, srv)
		pb.RegisterResourceServer(gs, srv)
	})
}

// GetResource 获取指定的资源配置信息
func (s *Resource) GetResource(c context.Context, req *pb.GetResourceRequest) (*pb.GetResourceResponse, error) {
	var (
		res *entity.Resource
		err error
	)
	switch req.Params.(type) {
	case *pb.GetResourceRequest_Id:
		res, err = s.srv.GetResource(core.MustContext(c), req.GetId())
	case *pb.GetResourceRequest_Keyword:
		res, err = s.srv.GetResourceByKeyword(core.MustContext(c), req.GetKeyword())
	default:
		err = errors.ParamsError()
	}
	if err != nil {
		return nil, err
	}

	reply := pb.GetResourceResponse{
		Id:          res.Id,
		Keyword:     res.Keyword,
		Fields:      res.Fields,
		Tag:         res.Tag,
		Private:     res.Private,
		Description: res.Description,
		CreatedAt:   uint32(res.CreatedAt),
		UpdatedAt:   uint32(res.UpdatedAt),
	}
	for _, item := range res.ResourceApps {
		reply.Apps = append(reply.Apps, item.App)
	}
	return &reply, nil
}

// ListResource 获取资源配置信息列表
func (s *Resource) ListResource(c context.Context, req *pb.ListResourceRequest) (*pb.ListResourceResponse, error) {
	list, total, err := s.srv.ListResource(core.MustContext(c), &types.ListResourceRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
		Keyword:  req.Keyword,
		Tag:      req.Tag,
		Private:  req.Private,
		App:      req.App,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListResourceResponse{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListResourceResponse_Resource{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Fields:      item.Fields,
			Tag:         item.Tag,
			Private:     item.Private,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateResource 创建资源配置信息
func (s *Resource) CreateResource(c context.Context, req *pb.CreateResourceRequest) (*pb.CreateResourceResponse, error) {
	var res = entity.Resource{
		Keyword:     req.Keyword,
		Fields:      req.Fields,
		Tag:         req.Tag,
		Private:     req.Private,
		Description: req.Description,
	}
	for _, app := range req.Apps {
		res.ResourceApps = append(res.ResourceApps, &entity.ResourceApp{
			App: app,
		})
	}
	id, err := s.srv.CreateResource(core.MustContext(c), &res)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResourceResponse{Id: id}, nil
}

// UpdateResource 更新资源配置信息
func (s *Resource) UpdateResource(c context.Context, req *pb.UpdateResourceRequest) (*pb.UpdateResourceResponse, error) {
	var res = entity.Resource{
		BaseModel:   kmodel.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Fields:      req.Fields,
		Tag:         req.Tag,
		Private:     req.Private,
		Description: req.Description,
	}
	for _, app := range req.Apps {
		res.ResourceApps = append(res.ResourceApps, &entity.ResourceApp{
			App: app,
		})
	}

	if err := s.srv.UpdateResource(core.MustContext(c), &res); err != nil {
		return nil, err
	}
	return &pb.UpdateResourceResponse{}, nil
}

// DeleteResource 删除资源配置信息
func (s *Resource) DeleteResource(c context.Context, req *pb.DeleteResourceRequest) (*pb.DeleteResourceResponse, error) {
	err := s.srv.DeleteResource(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResourceResponse{}, nil
}

// ListResourceValue 获取业务配置值信息列表
func (s *Resource) ListResourceValue(c context.Context, req *pb.ListResourceValueRequest) (*pb.ListResourceValueResponse, error) {
	list, err := s.srv.ListResourceValue(core.MustContext(c), req.ResourceId)
	if err != nil {
		return nil, err
	}

	reply := pb.ListResourceValueResponse{Total: uint32(len(list))}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListResourceValueResponse_ResourceValue{
			Id:         item.Id,
			EnvId:      item.EnvId,
			ResourceId: item.ResourceId,
			Value:      item.Value,
			CreatedAt:  uint32(item.CreatedAt),
			UpdatedAt:  uint32(item.UpdatedAt),
		})
	}

	return &reply, nil
}

// UpdateResourceValue 更新业务配置值信息
func (s *Resource) UpdateResourceValue(c context.Context, req *pb.UpdateResourceValueRequest) (*pb.UpdateResourceValueResponse, error) {
	var list []*entity.ResourceValue
	for _, item := range req.List {
		list = append(list, &entity.ResourceValue{
			ResourceId: req.ResourceId,
			EnvId:      item.EnvId,
			Value:      item.Value,
		})
	}

	if err := s.srv.UpdateResourceValue(core.MustContext(c), list); err != nil {
		return nil, err
	}

	return &pb.UpdateResourceValueResponse{}, nil
}
