package logic

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// PageResource 分页查询服务
func (l *Logic) PageResource(ctx kratos.Context, in *v1.PageResourceRequest) (*v1.PageResourceReply, error) {
	resource := model.Resource{}
	list, err := resource.Page(ctx, &model.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.Keyword != nil {
				db = db.Where("keyword like ?", "%"+*in.Keyword+"%")
			}
			return db
		},
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.PageResourceReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// AddResource 添加服务
func (l *Logic) AddResource(ctx kratos.Context, in *v1.AddResourceRequest) (*emptypb.Empty, error) {
	resource := model.Resource{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.Create(ctx)
}

// UpdateResource 更新服务
func (l *Logic) UpdateResource(ctx kratos.Context, in *v1.UpdateResourceRequest) (*emptypb.Empty, error) {
	resource := model.Resource{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.Update(ctx)
}

// DeleteResource 删除服务
func (l *Logic) DeleteResource(ctx kratos.Context, in *v1.DeleteResourceRequest) (*emptypb.Empty, error) {
	resource := model.Resource{}
	return nil, resource.DeleteByID(ctx, in.Id)
}
