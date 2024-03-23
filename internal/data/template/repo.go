package template

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/configure/internal/biz/template"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

// AddTemplate 新建模板
func (s *repo) AddTemplate(ctx kratosx.Context, template *biz.Template) (uint32, error) {
	return template.ID, ctx.DB().Create(template).Error
}

// GetTemplate 通过ID查找指定模板
func (s *repo) GetTemplate(ctx kratosx.Context, id uint32) (*biz.Template, error) {
	var template biz.Template
	return &template, ctx.DB().First(&template, "id=?", id).Error
}

// CurrentTemplate 获取当前模板
func (s *repo) CurrentTemplate(ctx kratosx.Context, id uint32) (*biz.Template, error) {
	var template biz.Template
	return &template, ctx.DB().First(&template, "is_use=true and server_id=?", id).Error
}

func (s *repo) UseTemplate(ctx kratosx.Context, srvId, tpId uint32) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(biz.Template{}).Where("server_id=? and id=?", srvId, tpId).Update("is_use", true).Error; err != nil {
			return err
		}
		return tx.Model(biz.Template{}).Where("server_id=? and id!=?", srvId, tpId).Update("is_use", false).Error
	})
}

// GetTemplateByVersion 通过版本号查找指定模板
func (s *repo) GetTemplateByVersion(ctx kratosx.Context, keyword string) (*biz.Template, error) {
	var template biz.Template
	return &template, ctx.DB().First(&template, "version=?", keyword).Error
}

// PageTemplate 查询分页模板
func (s *repo) PageTemplate(ctx kratosx.Context, req *biz.PageTemplateRequest) ([]*biz.Template, uint32, error) {
	var list []*biz.Template
	var total int64

	db := ctx.DB().Model(biz.Template{}).Where("server_id=?", req.ServerId)

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Order("created_at desc").Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// UpdateTemplate 更新指定id的模板
func (s *repo) UpdateTemplate(ctx kratosx.Context, template *biz.Template) error {
	return ctx.DB().Where("id=?", template.ID).Updates(template).Error
}

// DeleteTemplate 删除指定id的模板
func (s *repo) DeleteTemplate(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Template{}, "id = ?", id).Error
}
