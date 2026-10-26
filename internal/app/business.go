package app

import (
	"context"

	kmodel "github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/pkg/value"

	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pb "github.com/limes-cloud/configure/api/business"
	"github.com/limes-cloud/configure/api/errors"
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
		App:      resolveApp(core.MustContext(c), req.GetApp()),
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
		App:         resolveApp(core.MustContext(c), req.GetApp()),
		Keyword:     req.Keyword,
		Type:        req.Type,
		Extra:       req.GetExtra(),
		Description: value.String(value.Value(req.Description)), // 描述列 NOT NULL，nil 兜底空串指针
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateBusinessResponse{Id: id}, nil
}

// UpdateBusiness 更新业务配置信息
func (s *Business) UpdateBusiness(c context.Context, req *pb.UpdateBusinessRequest) (*pb.UpdateBusinessResponse, error) {
	if err := s.checkAppPermission(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	if err := s.srv.UpdateBusiness(core.MustContext(c), &entity.Business{
		BaseModel:   kmodel.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Type:        req.Type,
		Extra:       req.GetExtra(),
		Description: value.String(value.Value(req.Description)), // 描述列 NOT NULL，nil 兜底空串指针
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateBusinessResponse{}, nil
}

// DeleteBusiness 删除业务配置信息
func (s *Business) DeleteBusiness(c context.Context, req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessResponse, error) {
	if err := s.checkAppPermission(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
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

// resolveApp 解析业务配置所属应用：
// 未传 app（页面直接调接口）时默认取当前应用 AppKey（X-App-Key 请求头）；
// 仅 manager 允许显式指定其他 app；非 manager 传其他 app 一律强制回 AppKey（防越权）。
// AppKey 为空（未鉴权的内部调用）时回退到请求参数，保持向后兼容。
func resolveApp(ctx core.Context, app string) string {
	appKey := ctx.AppKey()
	if appKey == "" {
		return app
	}
	if appKey != "manager" && app != "" && app != appKey {
		return appKey
	}
	if app == "" {
		return appKey
	}
	return app
}

// checkAppPermission 校验业务配置归属：非 manager 只能操作自己应用（AppKey）下的业务配置，越权返回 NotPermissionError
func (s *Business) checkAppPermission(ctx core.Context, id uint32) error {
	appKey := ctx.AppKey()
	if appKey == "" || appKey == "manager" {
		return nil
	}
	business, err := s.srv.GetBusiness(ctx, id)
	if err != nil {
		return err
	}
	if business.App != appKey {
		return errors.NotPermissionError()
	}
	return nil
}
