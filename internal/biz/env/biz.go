package env

import (
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	"github.com/limes-cloud/kratosx/types"
	"github.com/limes-cloud/manager/api/auth"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/config"
)

type UseCase struct {
	config *config.Config
	repo   Repo
}

func NewUseCase(config *config.Config, repo Repo) *UseCase {
	return &UseCase{config: config, repo: repo}
}

// AddEnv 新建环境
func (e *UseCase) AddEnv(ctx kratosx.Context, env *Env) (uint32, error) {
	id, err := e.repo.AddEnv(ctx, env)
	if err != nil {
		return 0, errors.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// GetEnv 查询指定的服务
func (e *UseCase) GetEnv(ctx kratosx.Context, id uint32) (*Env, error) {
	env, err := e.repo.GetEnv(ctx, id)
	if err != nil {
		return nil, errors.NotRecordError()
	}
	return env, nil
}

// GetEnvByKeyword 通过关键词查找指定环境
func (e *UseCase) GetEnvByKeyword(ctx kratosx.Context, keyword string) (*Env, error) {
	env, err := e.repo.GetEnvByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.DatabaseErrorFormat(err.Error())
	}
	return env, nil
}

// GetEnvByToken 通过token查找指定环境
func (e *UseCase) GetEnvByToken(ctx kratosx.Context, token string) (*Env, error) {
	env, err := e.repo.GetEnvByToken(ctx, token)
	if err != nil {
		return nil, errors.DatabaseErrorFormat(err.Error())
	}
	return env, nil
}

// AllEnv 查询所有环境
func (e *UseCase) AllEnv(ctx kratosx.Context) ([]*Env, error) {
	var scope []string
	if info, err := auth.Get(ctx); err == nil && info.Scope["env_scope"] != nil {
		scope = info.Scope["env_scope"].List
	}

	envs, err := e.repo.AllEnv(ctx, scope...)
	if err != nil {
		return nil, errors.DatabaseErrorFormat(err.Error())
	}
	return envs, nil
}

// // GetEnvByIds 查询所有环境
// func (e *UseCase) GetEnvByIds(ctx kratosx.Context, ids []uint32) ([]*Env, error) {
//	envs, err := e.repo.GetEnvByIds(ctx, ids)
//	if err != nil {
//		return nil, errors.DatabaseErrorFormat(err.Error())
//	}
//	return envs, nil
// }

// UpdateEnv 更新环境
func (e *UseCase) UpdateEnv(ctx kratosx.Context, env *Env) error {
	if err := e.repo.UpdateEnv(ctx, env); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// DeleteEnv 删除指定id的环境
func (e *UseCase) DeleteEnv(ctx kratosx.Context, id uint32) error {
	if err := e.repo.DeleteEnv(ctx, id); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// ResetEnvToken 重置环境token
func (e *UseCase) ResetEnvToken(ctx kratosx.Context, id uint32) (string, error) {
	env := &Env{
		BaseModel: types.BaseModel{ID: id},
		Token:     util.MD5ToUpper([]byte(uuid.NewString())),
	}
	if err := e.repo.UpdateEnv(ctx, env); err != nil {
		return "", errors.DatabaseErrorFormat(err.Error())
	}
	return env.Token, nil
}
