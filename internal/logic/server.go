package logic

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

type Server struct {
	conf *config.Config
}

func NewServer(conf *config.Config) *Server {
	return &Server{
		conf: conf,
	}
}

// Get 分页查询服务
func (s *Server) Get(ctx kratosx.Context, in *v1.GetServerRequest) (*v1.GetServerReply, error) {
	var err error
	var server model.Server

	if in.Id == nil && in.Keyword == nil {
		return nil, v1.ParamsError()
	}

	if in.Id != nil {
		err = server.OneByID(ctx, *in.Id)
	} else {
		err = server.OneByKeyword(ctx, *in.Keyword)
	}

	if err != nil {
		return nil, v1.NotRecordError()
	}
	reply := v1.GetServerReply{}
	if util.Transform(server, &reply.Server) != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

// Page 分页查询服务
func (s *Server) Page(ctx kratosx.Context, in *v1.PageServerRequest) (*v1.PageServerReply, error) {
	server := model.Server{}
	list, total, err := server.Page(ctx, &types.PageOptions{
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
		return nil, v1.DatabaseError()
	}
	reply := v1.PageServerReply{
		Total: total,
	}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

// Add 添加服务
func (s *Server) Add(ctx kratosx.Context, in *v1.AddServerRequest) (*emptypb.Empty, error) {
	server := model.Server{}

	if server.OneByKeyword(ctx, in.Keyword) == nil {
		return nil, v1.AlreadyExistsError()
	}

	if util.Transform(in, &server) != nil {
		return nil, v1.TransformError()
	}

	return nil, server.Create(ctx)
}

// Update 更新服务
func (s *Server) Update(ctx kratosx.Context, in *v1.UpdateServerRequest) (*emptypb.Empty, error) {
	server := model.Server{}
	if util.Transform(in, &server) != nil {
		return nil, v1.TransformError()
	}

	return nil, server.Update(ctx)
}

// Delete 删除服务
func (s *Server) Delete(ctx kratosx.Context, in *v1.DeleteServerRequest) (*emptypb.Empty, error) {
	server := model.Server{}
	return nil, server.DeleteByID(ctx, in.Id)
}
