package service

import (
	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/configure/internal/types"
	"github.com/limes-cloud/kratosx"
)

type Server struct {
	conf       *conf.Config
	repo       repository.ServerRepository
	permission repository.PermissionRepository
}

func NewServer(
	conf *conf.Config,
	repo repository.ServerRepository,
	permission repository.PermissionRepository,
) *Server {
	return &Server{
		conf:       conf,
		repo:       repo,
		permission: permission,
	}
}

// ListServer 获取服务信息列表
func (u *Server) ListServer(ctx kratosx.Context, req *types.ListServerRequest) ([]*entity.Server, uint32, error) {
	// 获取服务权限id列表
	all, scopes, err := u.permission.GetServer(ctx)
	if err != nil {
		return nil, 0, err
	}
	if !all {
		req.Ids = scopes
	}

	// 获取列表
	list, total, err := u.repo.ListServer(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateServer 创建服务信息
func (u *Server) CreateServer(ctx kratosx.Context, req *entity.Server) (uint32, error) {
	id, err := u.repo.CreateServer(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateServer 更新服务信息
func (u *Server) UpdateServer(ctx kratosx.Context, req *entity.Server) error {
	// 服务鉴权
	if !u.permission.HasServer(ctx, req.Id) {
		return errors.NotPermissionError()
	}

	// 更新服务
	if err := u.repo.UpdateServer(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteServer 删除服务信息
func (u *Server) DeleteServer(ctx kratosx.Context, id uint32) error {
	// 服务鉴权
	if !u.permission.HasServer(ctx, id) {
		return errors.NotPermissionError()
	}

	// 删除服务
	if err := u.repo.DeleteServer(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
