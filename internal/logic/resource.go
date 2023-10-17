package logic

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/md"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type Resource struct {
	conf *config.Config
}

func NewResource(conf *config.Config) *Resource {
	return &Resource{
		conf: conf,
	}
}

// Page 分页查询资源
func (r *Resource) Page(ctx kratos.Context, in *v1.PageResourceRequest) (*v1.PageResourceReply, error) {
	resource := model.Resource{}
	list, total, err := resource.Page(ctx, &model.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.Keyword != nil {
				db = db.Where("keyword like ?", "%"+*in.Keyword+"%")
			}
			if in.Tag != nil {
				db = db.Where("tag like ?", "%"+*in.Tag+"%")
			}
			return db
		},
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.PageResourceReply{
		Total: total,
	}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Add 添加资源
func (r *Resource) Add(ctx kratos.Context, in *v1.AddResourceRequest) (*emptypb.Empty, error) {
	resource := model.Resource{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.Create(ctx)
}

// Update 更新资源
func (r *Resource) Update(ctx kratos.Context, in *v1.UpdateResourceRequest) (*emptypb.Empty, error) {
	resource := model.Resource{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.Update(ctx)
}

// Delete 删除资源
func (r *Resource) Delete(ctx kratos.Context, in *v1.DeleteResourceRequest) (*emptypb.Empty, error) {
	resource := model.Resource{}
	return nil, resource.DeleteByID(ctx, in.Id)
}
