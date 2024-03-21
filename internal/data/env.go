package data

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	"github.com/limes-cloud/configure/internal/biz"
)

type envRepo struct {
}

func NewEnvRepo() biz.EnvRepo {
	return &envRepo{}
}

func (u *envRepo) Create(ctx kratosx.Context, env *biz.Env) (uint32, error) {
	return env.ID, ctx.DB().Create(env).Error
}

func (u *envRepo) Update(ctx kratosx.Context, env *biz.Env) error {
	return ctx.DB().Where("id=?", env.ID).Updates(env).Error
}

func (u *envRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Env{}, "id=?", id).Error
}

func (u *envRepo) Get(ctx kratosx.Context, id uint32) (*biz.Env, error) {
	var ins biz.Env
	return &ins, ctx.DB().First(&ins, "id=?", id).Error
}

func (u *envRepo) GetByKeyword(ctx kratosx.Context, keyword string) (*biz.Env, error) {
	var ins biz.Env
	return &ins, ctx.DB().First(&ins, "keyword=?", keyword).Error
}

func (u *envRepo) GetByToken(ctx kratosx.Context, token string) (*biz.Env, error) {
	var ins biz.Env
	return &ins, ctx.DB().First(&ins, "token=?", token).Error
}

func (u *envRepo) All(ctx kratosx.Context, scope ...string) ([]*biz.Env, error) {
	var list []*biz.Env
	return list, ctx.DB().Model(biz.Env{}).Scopes(func(db *gorm.DB) *gorm.DB {
		if len(scope) != 0 {
			return db.Where("id in ?", scope)
		}
		return db
	}).Find(&list).Error
}

func (u *envRepo) GetByIds(ctx kratosx.Context, ids []uint32) ([]*biz.Env, error) {
	var list []*biz.Env
	return list, ctx.DB().Model(biz.Env{}).Where("id in ?", ids).Find(&list).Error
}
