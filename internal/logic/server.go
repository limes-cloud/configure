package logic

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// PageServer 分页查询服务
func (l *Logic) PageServer(ctx kratos.Context, in *v1.PageServerRequest) (*v1.PageServerReply, error) {
	server := model.Server{}
	list, err := server.Page(ctx, &model.PageOptions{
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
	reply := v1.PageServerReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}

	return &reply, nil
}

// AddServer 添加服务
func (l *Logic) AddServer(ctx kratos.Context, in *v1.AddServerRequest) (*emptypb.Empty, error) {
	server := model.Server{}
	if util.Transform(in, &server) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, server.Create(ctx)
}

// UpdateServer 更新服务
func (l *Logic) UpdateServer(ctx kratos.Context, in *v1.UpdateServerRequest) (*emptypb.Empty, error) {
	server := model.Server{}
	if util.Transform(in, &server) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, server.Update(ctx)
}

// DeleteServer 删除服务
func (l *Logic) DeleteServer(ctx kratos.Context, in *v1.DeleteServerRequest) (*emptypb.Empty, error) {
	server := model.Server{}
	return nil, server.DeleteByID(ctx, in.Id)
}