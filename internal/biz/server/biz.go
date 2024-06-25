package server

import (
	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/kratosx"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// GetServer 获取指定的服务信息
func (u *UseCase) GetServer(ctx kratosx.Context, req *GetServerRequest) (*Server, error) {
	var (
		res *Server
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetServer(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetServerByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListServer 获取服务信息列表
func (u *UseCase) ListServer(ctx kratosx.Context, req *ListServerRequest) ([]*Server, uint32, error) {
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
	if err := u.repo.UpdateServer(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteServer 删除服务信息
func (u *UseCase) DeleteServer(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteServer(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// UpdateServerStatus 更新服务信息状态
func (u *UseCase) UpdateServerStatus(ctx kratosx.Context, id uint32, status bool) error {
	if err := u.repo.UpdateServerStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
