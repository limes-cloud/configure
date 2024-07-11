package data

import (
	"errors"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	biz "github.com/limes-cloud/configure/internal/biz/env"
	"github.com/limes-cloud/configure/internal/data/model"
)

type envRepo struct {
}

func NewEnvRepo() biz.Repo {
	return &envRepo{}
}

// ToEnvEntity model转entity
func (r envRepo) ToEnvEntity(m *model.Env) *biz.Env {
	e := &biz.Env{}
	_ = valx.Transform(m, e)
	return e
}

// ToEnvModel entity转model
func (r envRepo) ToEnvModel(e *biz.Env) *model.Env {
	m := &model.Env{}
	_ = valx.Transform(e, m)
	return m
}

// GetEnvByKeyword 获取指定数据
func (r envRepo) GetEnvByKeyword(ctx kratosx.Context, keyword string) (*biz.Env, error) {
	var (
		m  = model.Env{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToEnvEntity(&m), nil
}

// GetEnv 获取指定的数据
func (r envRepo) GetEnv(ctx kratosx.Context, id uint32) (*biz.Env, error) {
	var (
		m  = model.Env{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToEnvEntity(&m), nil
}

// ListEnv 获取列表
func (r envRepo) ListEnv(ctx kratosx.Context, req *biz.ListEnvRequest) ([]*biz.Env, uint32, error) {
	var (
		bs    []*biz.Env
		ms    []*model.Env
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Env{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	// 资源权限限制
	all, scopes, err := GetEnvPermission(ctx)
	if err != nil {
		return nil, 0, errors.New(err.Error())
	}
	if !all {
		db = db.Where("id  in ?", scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToEnvEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateEnv 创建数据
func (r envRepo) CreateEnv(ctx kratosx.Context, req *biz.Env) (uint32, error) {
	m := r.ToEnvModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateEnv 更新数据
func (r envRepo) UpdateEnv(ctx kratosx.Context, req *biz.Env) error {
	return ctx.DB().Updates(r.ToEnvModel(req)).Error
}

// DeleteEnv 删除数据
func (r envRepo) DeleteEnv(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Env{})
	return uint32(db.RowsAffected), db.Error
}

// UpdateEnvStatus 更新数据状态
func (r envRepo) UpdateEnvStatus(ctx kratosx.Context, id uint32, status bool) error {
	return ctx.DB().Model(model.Env{}).Where("id=?", id).Update("status", status).Error
}
