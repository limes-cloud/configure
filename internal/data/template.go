package data

import (
	"github.com/limes-cloud/configure/internal/biz/types"
	"gorm.io/gorm"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type templateRepo struct {
}

func NewTemplateRepo() biz.TemplateRepo {
	return &templateRepo{}
}

// Create 新建模板
func (s *templateRepo) Create(ctx kratosx.Context, template *biz.Template) (uint32, error) {
	return template.ID, ctx.DB().Create(template).Error
}

// Get 通过ID查找指定模板
func (s *templateRepo) Get(ctx kratosx.Context, id uint32) (*biz.Template, error) {
	var template biz.Template
	return &template, ctx.DB().First(&template, "id=?", id).Error
}

// Current 获取当前模板
func (s *templateRepo) Current(ctx kratosx.Context, id uint32) (*biz.Template, error) {
	var template biz.Template
	return &template, ctx.DB().First(&template, "is_use=true and server_id=?", id).Error
}

func (s *templateRepo) Use(ctx kratosx.Context, srvId, tpId uint32) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(biz.Template{}).Where("server_id=? and id=?", srvId, tpId).Update("is_use", true).Error; err != nil {
			return err
		}
		return tx.Model(biz.Template{}).Where("server_id=? and id!=?", srvId, tpId).Update("is_use", false).Error
	})
}

// GetByVersion 通过版本号查找指定模板
func (s *templateRepo) GetByVersion(ctx kratosx.Context, keyword string) (*biz.Template, error) {
	var template biz.Template
	return &template, ctx.DB().First(&template, "version=?", keyword).Error
}

// Page 查询分页模板
func (s *templateRepo) Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*biz.Template, uint32, error) {
	var list []*biz.Template
	var total int64

	db := ctx.DB().Model(biz.Template{})
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Order("created_at desc").Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (s templateRepo) PageServerTemplate(ctx kratosx.Context, req *types.PageTemplateRequest) ([]*biz.Template, uint32, error) {
	return s.Page(ctx, &ktypes.PageOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			return db.Where("server_id=?", req.ServerID)
		},
	})
}

// All 查询全部模板
func (s *templateRepo) All(ctx kratosx.Context, scopes ktypes.Scopes) ([]*biz.Template, error) {
	var list []*biz.Template

	db := ctx.DB().Model(biz.Template{})
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的模板
func (s *templateRepo) Update(ctx kratosx.Context, template *biz.Template) error {
	return ctx.DB().Where("id=?", template.ID).Updates(template).Error
}

// Delete 删除指定id的模板
func (s *templateRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Template{}, "id = ?", id).Error
}
