package dbs

import (
	"sync"

	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type TemplateInfra struct {
}

var (
	templateIns  *TemplateInfra
	templateOnce sync.Once
)

func NewTemplateInfra() *TemplateInfra {
	templateOnce.Do(func() {
		templateIns = &TemplateInfra{}
	})
	return templateIns
}

// GetTemplateByVersion 获取指定版本的模板
func (r TemplateInfra) GetTemplateByVersion(ctx kratosx.Context, version string) (*entity.Template, error) {
	var (
		template = entity.Template{}
		fs       = []string{"*"}
	)
	return &template, ctx.DB().Select(fs).Where("version = ?", version).First(&template).Error
}

// GetTemplate 获取指定的数据
func (r TemplateInfra) GetTemplate(ctx kratosx.Context, id uint32) (*entity.Template, error) {
	var (
		template = entity.Template{}
		fs       = []string{"*"}
	)
	return &template, ctx.DB().Select(fs).First(&template, id).Error
}

// CurrentTemplate 获取当前正在使用的模板
func (r TemplateInfra) CurrentTemplate(ctx kratosx.Context, sid uint32) (*entity.Template, error) {
	var (
		template = entity.Template{}
		fs       = []string{"*"}
	)
	return &template, ctx.DB().Select(fs).First(&template, "is_use=true and server_id=?", sid).Error
}

// UseTemplate 使用模板
func (r TemplateInfra) UseTemplate(ctx kratosx.Context, srvId, tpId uint32) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(entity.Template{}).Where("server_id=? and id=?", srvId, tpId).Update("is_use", true).Error; err != nil {
			return err
		}
		return tx.Model(entity.Template{}).Where("server_id=? and id!=?", srvId, tpId).Update("is_use", false).Error
	})
}

// ListTemplate 获取列表
func (r TemplateInfra) ListTemplate(ctx kratosx.Context, req *types.ListTemplateRequest) ([]*entity.Template, uint32, error) {
	var (
		list  []*entity.Template
		total int64
		fs    = []string{"id", "is_use", "version", "compare", "description", "created_at", "updated_at"}
	)

	db := ctx.DB().Model(entity.Template{}).Select(fs).Where("server_id=?", req.ServerId)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// CreateTemplate 创建数据
func (r TemplateInfra) CreateTemplate(ctx kratosx.Context, template *entity.Template) (uint32, error) {
	return template.Id, ctx.DB().Create(template).Error
}

// UpdateTemplate 更新数据
func (r TemplateInfra) UpdateTemplate(ctx kratosx.Context, template *entity.Template) error {
	return ctx.DB().Where("id = ?", template.Id).Updates(template).Error
}

// DeleteTemplate 删除数据
func (r TemplateInfra) DeleteTemplate(ctx kratosx.Context, id uint32) error {
	db := ctx.DB().Where("id=?", id).Delete(&entity.Template{})
	return db.Error
}
