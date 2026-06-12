package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/limes-cloud/configure/internal/core"
	kmodel "github.com/limes-cloud/kratosx/model"
	"github.com/limes-cloud/kratosx/pkg/crypto"

	pb "github.com/limes-cloud/configure/api/env"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/service"
	"github.com/limes-cloud/configure/internal/infra/dbs"
	"github.com/limes-cloud/configure/internal/types"
)

type Env struct {
	pb.UnimplementedEnvServer
	srv *service.Env
}

func NewEnv() *Env {
	return &Env{
		srv: service.NewEnv(dbs.NewEnv()),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewEnv()
		pb.RegisterEnvHTTPServer(hs, srv)
		pb.RegisterEnvServer(gs, srv)
	})
}

// ListEnv 获取环境信息列表
func (app *Env) ListEnv(c context.Context, req *pb.ListEnvRequest) (*pb.ListEnvResponse, error) {
	list, total, err := app.srv.ListEnv(core.MustContext(c), &types.ListEnvRequest{
		Keyword: req.Keyword,
		Name:    req.Name,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	reply := pb.ListEnvResponse{Total: total}
	for _, item := range list {
		reply.List = append(reply.List, &pb.ListEnvResponse_Env{
			Id:          item.Id,
			Keyword:     item.Keyword,
			Name:        item.Name,
			Status:      item.Status,
			Description: item.Description,
			CreatedAt:   uint32(item.CreatedAt),
			UpdatedAt:   uint32(item.UpdatedAt),
		})
	}
	return &reply, nil
}

// CreateEnv 创建环境信息
func (app *Env) CreateEnv(c context.Context, req *pb.CreateEnvRequest) (*pb.CreateEnvResponse, error) {
	id, err := app.srv.CreateEnv(core.MustContext(c), &entity.Env{
		Keyword:     req.Keyword,
		Name:        req.Name,
		Status:      req.Status,
		Description: req.Description,
		Token:       crypto.MD5ToUpper([]byte(uuid.NewString())),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateEnvResponse{Id: id}, nil
}

// UpdateEnv 更新环境信息
func (app *Env) UpdateEnv(c context.Context, req *pb.UpdateEnvRequest) (*pb.UpdateEnvResponse, error) {
	if err := app.srv.UpdateEnv(core.MustContext(c), &entity.Env{
		BaseModel:   kmodel.BaseModel{Id: req.Id},
		Keyword:     req.Keyword,
		Name:        req.Name,
		Status:      req.Status,
		Description: req.Description,
	}); err != nil {
		return nil, err
	}
	return &pb.UpdateEnvResponse{}, nil
}

// DeleteEnv 删除环境信息
func (app *Env) DeleteEnv(c context.Context, req *pb.DeleteEnvRequest) (*pb.DeleteEnvResponse, error) {
	if err := app.srv.DeleteEnv(core.MustContext(c), req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteEnvResponse{}, nil
}

// GetEnvToken 获取环境token
func (app *Env) GetEnvToken(c context.Context, req *pb.GetEnvTokenRequest) (*pb.GetEnvTokenResponse, error) {
	token, err := app.srv.GetEnvToken(core.MustContext(c), req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetEnvTokenResponse{Token: token}, nil
}

// ResetEnvToken 重置token密码
func (app *Env) ResetEnvToken(c context.Context, req *pb.ResetEnvTokenRequest) (*pb.ResetEnvTokenResponse, error) {
	in := entity.Env{
		BaseModel: kmodel.BaseModel{Id: req.Id},
		Token:     crypto.MD5ToUpper([]byte(uuid.NewString())),
	}
	return &pb.ResetEnvTokenResponse{Token: in.Token}, app.srv.UpdateEnv(core.MustContext(c), &in)
}
