package logic

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// AllResourceServer 分页查询服务
func (l *Logic) AllResourceServer(ctx kratos.Context, in *v1.AllResourceServerRequest) (*v1.AllResourceServerReply, error) {
	resource := model.ResourceServer{}
	list, err := resource.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("resource_id =?", in.ResourceId)
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.AllResourceServerReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// AddResourceServer 添加服务
func (l *Logic) AddResourceServer(ctx kratos.Context, in *v1.AddResourceServerRequest) (*emptypb.Empty, error) {
	resource := model.ResourceServer{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.Create(ctx)
}

// DeleteResourceServer 删除服务
func (l *Logic) DeleteResourceServer(ctx kratos.Context, in *v1.DeleteResourceServerRequest) (*emptypb.Empty, error) {
	resource := model.ResourceServer{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.DeleteByID(ctx, in.Id)
}
