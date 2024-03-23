package env

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/configure/internal/biz/env"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

func (u *repo) AddEnv(ctx kratosx.Context, env *biz.Env) (uint32, error) {
	return env.ID, ctx.DB().Create(env).Error
}

func (u *repo) UpdateEnv(ctx kratosx.Context, env *biz.Env) error {
	return ctx.DB().Where("id=?", env.ID).Updates(env).Error
}

func (u *repo) DeleteEnv(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Env{}, "id=?", id).Error
}

func (u *repo) GetEnv(ctx kratosx.Context, id uint32) (*biz.Env, error) {
	var ins biz.Env
	return &ins, ctx.DB().First(&ins, "id=?", id).Error
}

func (u *repo) GetEnvByKeyword(ctx kratosx.Context, keyword string) (*biz.Env, error) {
	var ins biz.Env
	return &ins, ctx.DB().First(&ins, "keyword=?", keyword).Error
}

func (u *repo) GetEnvByToken(ctx kratosx.Context, token string) (*biz.Env, error) {
	var ins biz.Env
	return &ins, ctx.DB().First(&ins, "token=?", token).Error
}

func (u *repo) AllEnv(ctx kratosx.Context, scope ...string) ([]*biz.Env, error) {
	var list []*biz.Env
	return list, ctx.DB().Model(biz.Env{}).Scopes(func(db *gorm.DB) *gorm.DB {
		if scope != nil {
			return db.Where("id in ?", scope)
		}
		return db
	}).Find(&list).Error
}

// func (u *repo) GetEnvByIds(ctx kratosx.Context, ids []uint32) ([]*biz.Env, error) {
//	var list []*biz.Env
//	return list, ctx.DB().Model(biz.Env{}).Where("id in ?", ids).Find(&list).Error
// }
