package biz

import (
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"github.com/limes-cloud/manager/api/auth"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/pkg/util"
)

type Env struct {
	types.BaseModel
	Keyword     string `json:"keyword" gorm:"not null;type:char(32) binary;comment:环境标识"`
	Name        string `json:"name" gorm:"not null;size:64;comment:环境名称"`
	Description string `json:"description" gorm:"not null;size:128;comment:环境描述"`
	Token       string `json:"token,omitempty" gorm:"not null;size:128;comment:环境token"`
	Status      *bool  `json:"status" gorm:"default:false;comment:环境状态"`
}

type EnvRepo interface {
	Get(ctx kratosx.Context, id uint32) (*Env, error)
	GetByKeyword(ctx kratosx.Context, keyword string) (*Env, error)
	GetByToken(ctx kratosx.Context, keyword string) (*Env, error)
	GetByIds(ctx kratosx.Context, ids []uint32) ([]*Env, error)
	All(ctx kratosx.Context, scope ...string) ([]*Env, error)
	Create(ctx kratosx.Context, c *Env) (uint32, error)
	Update(ctx kratosx.Context, c *Env) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}

type EnvUseCase struct {
	config *config.Config
	repo   EnvRepo
}

func NewEnvUseCase(config *config.Config, repo EnvRepo) *EnvUseCase {
	return &EnvUseCase{config: config, repo: repo}
}

// Add 新建环境
func (e *EnvUseCase) Add(ctx kratosx.Context, env *Env) (uint32, error) {
	id, err := e.repo.Create(ctx, env)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Get 查询指定的服务
func (e *EnvUseCase) Get(ctx kratosx.Context, id uint32) (*Env, error) {
	env, err := e.repo.Get(ctx, id)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return env, nil
}

// GetByKeyword 通过关键词查找指定环境
func (e *EnvUseCase) GetByKeyword(ctx kratosx.Context, keyword string) (*Env, error) {
	env, err := e.repo.GetByKeyword(ctx, keyword)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return env, nil
}

// GetByToken 通过token查找指定环境
func (e *EnvUseCase) GetByToken(ctx kratosx.Context, token string) (*Env, error) {
	env, err := e.repo.GetByToken(ctx, token)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return env, nil
}

// All 查询所有环境
func (e *EnvUseCase) All(ctx kratosx.Context) ([]*Env, error) {
	var scope []string
	if info, err := auth.Get(ctx); err == nil && info.Scope["env_scope"] != nil {
		scope = info.Scope["env_scope"].List
	}

	envs, err := e.repo.All(ctx, scope...)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return envs, nil
}

// GetByIds 查询所有环境
func (e *EnvUseCase) GetByIds(ctx kratosx.Context, ids []uint32) ([]*Env, error) {
	envs, err := e.repo.GetByIds(ctx, ids)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return envs, nil
}

// Update 更新环境
func (e *EnvUseCase) Update(ctx kratosx.Context, env *Env) error {
	if err := e.repo.Update(ctx, env); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除指定id的环境
func (e *EnvUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := e.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// ResetToken 重置环境token
func (e *EnvUseCase) ResetToken(ctx kratosx.Context, id uint32) (string, error) {
	env := &Env{
		BaseModel: types.BaseModel{ID: id},
		Token:     util.MD5ToUpper([]byte(uuid.NewString())),
	}
	if err := e.repo.Update(ctx, env); err != nil {
		return "", v1.DatabaseErrorFormat(err.Error())
	}
	return env.Token, nil
}
