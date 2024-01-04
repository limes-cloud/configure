package logic

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

type ResourceServer struct {
	conf *config.Config
}

func NewResourceServer(conf *config.Config) *ResourceServer {
	return &ResourceServer{
		conf: conf,
	}
}

// PageResource 查询指定服务的所有资源
func (rs *ResourceServer) PageResource(ctx kratosx.Context, in *v1.PageServerResourceRequest) (*v1.PageServerResourceReply, error) {
	resource := model.Resource{}
	list, total, err := resource.Page(ctx, &types.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			return db.Where("id in (select resource_id from resource_server where server_id=?) or private=false", in.ServerId)
		},
	})

	if err != nil {
		return nil, v1.DatabaseError()
	}
	reply := v1.PageServerResourceReply{
		Total: total,
	}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

// AllServer 查询指定资源的所有服务
func (rs *ResourceServer) AllServer(ctx kratosx.Context, in *v1.AllResourceServerRequest) (*v1.AllResourceServerReply, error) {
	resource := model.Server{}
	list, err := resource.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in (select server_id from resource_server where resource_id=?)", in.ResourceId)
	})

	if err != nil {
		return nil, v1.DatabaseError()
	}
	reply := v1.AllResourceServerReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}
