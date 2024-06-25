package env

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

// GetEnv 获取指定的环境信息
func (u *UseCase) GetEnv(ctx kratosx.Context, req *GetEnvRequest) (*Env, error) {
	var (
		res *Env
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetEnv(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetEnvByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListEnv 获取环境信息列表
func (u *UseCase) ListEnv(ctx kratosx.Context, req *ListEnvRequest) ([]*Env, uint32, error) {
	list, total, err := u.repo.ListEnv(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateEnv 创建环境信息
func (u *UseCase) CreateEnv(ctx kratosx.Context, req *Env) (uint32, error) {
	id, err := u.repo.CreateEnv(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateEnv 更新环境信息
func (u *UseCase) UpdateEnv(ctx kratosx.Context, req *Env) error {
	if err := u.repo.UpdateEnv(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteEnv 删除环境信息
func (u *UseCase) DeleteEnv(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteEnv(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// UpdateEnvStatus 更新环境信息状态
func (u *UseCase) UpdateEnvStatus(ctx kratosx.Context, id uint32, status bool) error {
	if err := u.repo.UpdateEnvStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
