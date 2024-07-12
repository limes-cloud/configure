package env

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// GetEnv 获取指定的环境信息
	GetEnv(ctx kratosx.Context, id uint32) (*Env, error)

	// ListEnv 获取环境信息列表
	ListEnv(ctx kratosx.Context, req *ListEnvRequest) ([]*Env, uint32, error)

	// CreateEnv 创建环境信息
	CreateEnv(ctx kratosx.Context, req *Env) (uint32, error)

	// UpdateEnv 更新环境信息
	UpdateEnv(ctx kratosx.Context, req *Env) error

	// DeleteEnv 删除环境信息
	DeleteEnv(ctx kratosx.Context, ids uint32) error

	// GetEnvByKeyword 获取指定的环境信息
	GetEnvByKeyword(ctx kratosx.Context, keyword string) (*Env, error)

	// UpdateEnvStatus 更新环境信息状态
	UpdateEnvStatus(ctx kratosx.Context, id uint32, status bool) error
}
