package env

import "github.com/limes-cloud/kratosx"

type Repo interface {
	GetEnv(ctx kratosx.Context, id uint32) (*Env, error)
	GetEnvByKeyword(ctx kratosx.Context, keyword string) (*Env, error)
	GetEnvByToken(ctx kratosx.Context, keyword string) (*Env, error)

	// GetEnvByIds(ctx kratosx.Context, ids []uint32) ([]*Env, error)
	AllEnv(ctx kratosx.Context, scope ...string) ([]*Env, error)

	AddEnv(ctx kratosx.Context, c *Env) (uint32, error)
	UpdateEnv(ctx kratosx.Context, c *Env) error
	DeleteEnv(ctx kratosx.Context, uint322 uint32) error
}
