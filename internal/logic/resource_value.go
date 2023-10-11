package logic

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// AllResourceValue 分页查询服务
func (l *Logic) AllResourceValue(ctx kratos.Context, in *v1.AllResourceValueRequest) (*v1.AllResourceValueReply, error) {
	resource := model.ResourceValue{}
	list, err := resource.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("resource_id =?", in.ResourceId)
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.AllResourceValueReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// AddResourceValue 添加服务
func (l *Logic) AddResourceValue(ctx kratos.Context, in *v1.AddResourceValueRequest) (*emptypb.Empty, error) {
	resource := model.ResourceValue{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.Create(ctx)
}

// UpdateResourceValue 更新服务
func (l *Logic) UpdateResourceValue(ctx kratos.Context, in *v1.UpdateResourceValueRequest) (*emptypb.Empty, error) {
	resource := model.ResourceValue{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.Update(ctx)
}
