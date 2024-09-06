package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type EnvInfra struct {
}

var (
	envIns  *EnvInfra
	envOnce sync.Once
)

func NewEnvInfra() *EnvInfra {
	envOnce.Do(func() {
		envIns = &EnvInfra{}
	})
	return envIns
}

// GetEnvByToken 获取指定数据
func (r EnvInfra) GetEnvByToken(ctx kratosx.Context, token string) (*entity.Env, error) {
	var (
		env = entity.Env{}
		fs  = []string{"*"}
	)
	return &env, ctx.DB().Select(fs).Where("token = ?", token).First(&env).Error
}

// GetEnvByKeyword 获取指定数据
func (r EnvInfra) GetEnvByKeyword(ctx kratosx.Context, keyword string) (*entity.Env, error) {
	var (
		env = entity.Env{}
		fs  = []string{"*"}
	)
	return &env, ctx.DB().Select(fs).Where("keyword = ?", keyword).First(&env).Error
}

// GetEnv 获取指定的数据
func (r EnvInfra) GetEnv(ctx kratosx.Context, id uint32) (*entity.Env, error) {
	var (
		env = entity.Env{}
		fs  = []string{"*"}
	)
	return &env, ctx.DB().Select(fs).First(&env, id).Error
}

// ListEnv 获取列表
func (r EnvInfra) ListEnv(ctx kratosx.Context, req *types.ListEnvRequest) ([]*entity.Env, uint32, error) {
	var (
		envs  []*entity.Env
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.Env{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Ids != nil {
		db = db.Where("id in ?", req.Ids)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return envs, uint32(total), db.Find(&envs).Error
}

// CreateEnv 创建环境数据
func (r EnvInfra) CreateEnv(ctx kratosx.Context, env *entity.Env) (uint32, error) {
	return env.Id, ctx.DB().Create(env).Error
}

// UpdateEnv 更新环境数据
func (r EnvInfra) UpdateEnv(ctx kratosx.Context, env *entity.Env) error {
	return ctx.DB().Where("id = ?", env.Id).Updates(env).Error
}

// DeleteEnv 删除环境数据
func (r EnvInfra) DeleteEnv(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id=?", id).Delete(&entity.Env{}).Error
}
