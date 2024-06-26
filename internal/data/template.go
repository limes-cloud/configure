package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/configure/internal/biz/template"
	"github.com/limes-cloud/configure/internal/data/model"
)

type templateRepo struct {
}

func NewTemplateRepo() biz.Repo {
	return &templateRepo{}
}

// ToTemplateEntity model转entity
func (r templateRepo) ToTemplateEntity(m *model.Template) *biz.Template {
	e := &biz.Template{}
	_ = valx.Transform(m, e)
	return e
}

// ToTemplateModel entity转model
func (r templateRepo) ToTemplateModel(e *biz.Template) *model.Template {
	m := &model.Template{}
	_ = valx.Transform(e, m)
	return m
}

// GetTemplateByVersion 获取指定版本的模板
func (r templateRepo) GetTemplateByVersion(ctx kratosx.Context, version string) (*biz.Template, error) {
	var (
		m  = model.Template{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("version = ?", version).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToTemplateEntity(&m), nil
}

// GetTemplate 获取指定的数据
func (r templateRepo) GetTemplate(ctx kratosx.Context, id uint32) (*biz.Template, error) {
	var (
		m  = model.Template{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return r.ToTemplateEntity(&m), nil
}

// CurrentTemplate 获取当前正在使用的模板
func (r templateRepo) CurrentTemplate(ctx kratosx.Context, sid uint32) (*biz.Template, error) {
	var (
		m  = model.Template{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, "is_use=true and server_id=?", sid).Error; err != nil {
		return nil, err
	}
	return r.ToTemplateEntity(&m), nil
}

// UseTemplate 使用模板
func (r templateRepo) UseTemplate(ctx kratosx.Context, srvId, tpId uint32) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(model.Template{}).Where("server_id=? and id=?", srvId, tpId).Update("is_use", true).Error; err != nil {
			return err
		}
		return tx.Model(model.Template{}).Where("server_id=? and id!=?", srvId, tpId).Update("is_use", false).Error
	})
}

// ListTemplate 获取列表
func (r templateRepo) ListTemplate(ctx kratosx.Context, req *biz.ListTemplateRequest) ([]*biz.Template, uint32, error) {
	var (
		bs    []*biz.Template
		ms    []*model.Template
		total int64
		fs    = []string{"id", "is_use", "version", "compare", "description", "created_at", "updated_at"}
	)

	db := ctx.DB().Model(model.Template{}).Select(fs).Where("server_id=?", req.ServerId)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToTemplateEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateTemplate 创建数据
func (r templateRepo) CreateTemplate(ctx kratosx.Context, req *biz.Template) (uint32, error) {
	m := r.ToTemplateModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateTemplate 更新数据
func (r templateRepo) UpdateTemplate(ctx kratosx.Context, req *biz.Template) error {
	return ctx.DB().Updates(r.ToTemplateModel(req)).Error
}

// DeleteTemplate 删除数据
func (r templateRepo) DeleteTemplate(ctx kratosx.Context, id uint32) error {
	db := ctx.DB().Where("id=?", id).Delete(&model.Template{})
	return db.Error
}
