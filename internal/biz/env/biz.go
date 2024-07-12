package env

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

// ListEnv 获取环境信息列表
func (u *UseCase) ListEnv(ctx kratosx.Context, req *ListEnvRequest) ([]*Env, uint32, error) {
	all, scopes, err := permission.GetEnv(ctx)
	if err != nil {
		return nil, 0, err
	}
	if !all {
		req.Ids = scopes
	}
	list, total, err := u.repo.ListEnv(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// GetEnvToken 获取环境token
func (u *UseCase) GetEnvToken(ctx kratosx.Context, id uint32) (string, error) {
	if !permission.HasEnv(ctx, id) {
		return "", errors.NotPermissionError()
	}

	env, err := u.repo.GetEnv(ctx, id)
	if err != nil {
		return "", errors.CreateError(err.Error())
	}
	return env.Token, nil
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
	if !permission.HasEnv(ctx, req.Id) {
		return errors.NotPermissionError()
	}

	if err := u.repo.UpdateEnv(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteEnv 删除环境信息
func (u *UseCase) DeleteEnv(ctx kratosx.Context, id uint32) error {
	if !permission.HasEnv(ctx, id) {
		return errors.NotPermissionError()
	}

	if err := u.repo.DeleteEnv(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// UpdateEnvStatus 更新环境信息状态
func (u *UseCase) UpdateEnvStatus(ctx kratosx.Context, id uint32, status bool) error {
	if !permission.HasEnv(ctx, id) {
		return errors.NotPermissionError()
	}

	if err := u.repo.UpdateEnvStatus(ctx, id, status); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
