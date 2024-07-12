package server

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/pkg/permission"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// ListServer 获取服务信息列表
func (u *UseCase) ListServer(ctx kratosx.Context, req *ListServerRequest) ([]*Server, uint32, error) {
	all, scopes, err := permission.GetServer(ctx)
	if err != nil {
		return nil, 0, err
	}
	if !all {
		req.Ids = scopes
	}
	list, total, err := u.repo.ListServer(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateServer 创建服务信息
func (u *UseCase) CreateServer(ctx kratosx.Context, req *Server) (uint32, error) {
	id, err := u.repo.CreateServer(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateServer 更新服务信息
func (u *UseCase) UpdateServer(ctx kratosx.Context, req *Server) error {
	if !permission.HasServer(ctx, req.Id) {
		return errors.NotPermissionError()
	}

	if err := u.repo.UpdateServer(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteServer 删除服务信息
func (u *UseCase) DeleteServer(ctx kratosx.Context, id uint32) error {
	if !permission.HasServer(ctx, id) {
		return errors.NotPermissionError()
	}

	if err := u.repo.DeleteServer(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// UpdateServerStatus 更新服务信息状态
func (u *UseCase) UpdateServerStatus(ctx kratosx.Context, id uint32, status bool) error {
	if !permission.HasServer(ctx, id) {
		return errors.NotPermissionError()
	}

	if err := u.repo.UpdateServerStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
