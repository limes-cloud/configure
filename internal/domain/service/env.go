package service

import (
	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/configure/internal/types"
)

type Env struct {
	repo repository.Env
}

func NewEnv(repo repository.Env) *Env {
	return &Env{repo: repo}
}

// ListEnv 获取环境信息列表
func (u *Env) ListEnv(ctx core.Context, req *types.ListEnvRequest) ([]*entity.Env, uint32, error) {
	list, total, err := u.repo.ListEnv(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// GetEnvToken 获取环境token
func (u *Env) GetEnvToken(ctx core.Context, id uint32) (string, error) {
	env, err := u.repo.GetEnv(ctx, id)
	if err != nil {
		return "", errors.GetError(err.Error())
	}
	return env.Token, nil
}

// CreateEnv 创建环境信息
func (u *Env) CreateEnv(ctx core.Context, req *entity.Env) (uint32, error) {
	id, err := u.repo.CreateEnv(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateEnv 更新环境信息
func (u *Env) UpdateEnv(ctx core.Context, req *entity.Env) error {
	if err := u.repo.UpdateEnv(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteEnv 删除环境信息
func (u *Env) DeleteEnv(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteEnv(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
