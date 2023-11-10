package logic

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/proto"
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
		Total: uint32(total),
	}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Add 添加资源
func (r *Resource) Add(ctx kratos.Context, in *v1.AddResourceRequest) (*emptypb.Empty, error) {
	if in.Private == nil {
		in.Private = proto.Bool(false)
	}

	resource := model.Resource{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	if err := resource.Create(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 私有变量进行所属服务写入
	if *in.Private {
		rs := model.ResourceServer{}
		if err := rs.CreateBySrvIds(ctx, resource.ID, in.Servers); err != nil {
			return nil, v1.ErrorDatabaseFormat(err.Error())
		}
	}

	return nil, nil
}

// Update 更新资源
func (r *Resource) Update(ctx kratos.Context, in *v1.UpdateResourceRequest) (*emptypb.Empty, error) {
	if in.Private == nil {
		in.Private = proto.Bool(false)
	}

	resource := model.Resource{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	if err := resource.Update(ctx); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 私有变量进行所属服务写入
	if *in.Private {
		rs := model.ResourceServer{}
		if err := rs.CreateBySrvIds(ctx, resource.ID, in.Servers); err != nil {
			return nil, v1.ErrorDatabaseFormat(err.Error())
		}
	}
	return nil, nil
}

// Delete 删除资源
func (r *Resource) Delete(ctx kratos.Context, in *v1.DeleteResourceRequest) (*emptypb.Empty, error) {
	resource := model.Resource{}
	return nil, resource.DeleteByID(ctx, in.Id)
}
