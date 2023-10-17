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

type ResourceServer struct {
	conf *config.Config
}

func NewResourceServer(conf *config.Config) *ResourceServer {
	return &ResourceServer{
		conf: conf,
	}
}

// AllResource 查询指定服务的所有资源
func (rs *ResourceServer) AllResource(ctx kratos.Context, in *v1.AllServerResourceRequest) (*v1.AllServerResourceReply, error) {
	resource := model.ResourceServer{}
	list, err := resource.All(ctx, func(db *gorm.DB) *gorm.DB {
		db.Preload("Resource")
		return db.Where("server_id =?", in.ServerId)
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.AllServerResourceReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// AllServer 查询指定资源的所有服务
func (rs *ResourceServer) AllServer(ctx kratos.Context, in *v1.AllResourceServerRequest) (*v1.AllResourceServerReply, error) {
	resource := model.ResourceServer{}
	list, err := resource.All(ctx, func(db *gorm.DB) *gorm.DB {
		db.Preload("Server")
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

// Add 添加资源-服务
func (rs *ResourceServer) Add(ctx kratos.Context, in *v1.AddResourceServerRequest) (*emptypb.Empty, error) {
	resource := model.ResourceServer{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.Create(ctx)
}

// Delete 删除资源-服务
func (rs *ResourceServer) Delete(ctx kratos.Context, in *v1.DeleteResourceServerRequest) (*emptypb.Empty, error) {
	resource := model.ResourceServer{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.DeleteByID(ctx, in.Id)
}
