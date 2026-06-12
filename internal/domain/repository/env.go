package repository

import (
	"github.com/limes-cloud/configure/internal/core"

	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type Env interface {
	// GetEnvByToken 获取指定的环境信息
	GetEnvByToken(ctx core.Context, token string) (*entity.Env, error)

	// GetEnvByKeyword 获取指定的环境信息
	GetEnvByKeyword(ctx core.Context, keyword string) (*entity.Env, error)

	// GetEnv 获取指定的环境信息
	GetEnv(ctx core.Context, id uint32) (*entity.Env, error)

	// ListEnv 获取环境信息列表
	ListEnv(ctx core.Context, req *types.ListEnvRequest) ([]*entity.Env, uint32, error)

	// CreateEnv 创建环境信息
	CreateEnv(ctx core.Context, env *entity.Env) (uint32, error)

	// UpdateEnv 更新环境信息
	UpdateEnv(ctx core.Context, env *entity.Env) error

	// DeleteEnv 删除环境信息
	DeleteEnv(ctx core.Context, ids uint32) error
}
