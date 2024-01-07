package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"
)

type Template struct {
	types.BaseModel
	ServerID    uint32  `json:"server_id" gorm:"uniqueIndex:sv;not null;comment:服务id"`
	Content     string  `json:"content" gorm:"not null;type:text;comment:模板内容"`
	Version     string  `json:"version" gorm:"uniqueIndex:sv;not null;size:32;comment:模板版本"`
	IsUse       bool    `json:"is_use" gorm:"default:false;comment:是否使用"`
	Format      string  `json:"format" gorm:"not null;size:32;comment:模板格式"`
	Description string  `json:"description" gorm:"not null;size:128;comment:模板描述"`
	Compare     string  `json:"compare" gorm:"not null;type:text;comment:变更详情"`
	Server      *Server `json:"server" gorm:"constraint:onDelete:cascade"`
}

func (t *Template) Create(ctx kratosx.Context) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(t).Create(&t).Error; err != nil {
			return err
		}
		temp := Template{}
		return tx.Model(&temp).Where("server_id=? and id!=?", t.ServerID, t.ID).Update("is_use", false).Error
	})
}

func (t *Template) OneById(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(t, "id = ?", id).Error
}

func (t *Template) Current(ctx kratosx.Context, srvId uint32) error {
	return ctx.DB().First(t, "is_use = true and server_id = ?", srvId).Error
}

// Page 查询分页资源
func (t *Template) Page(ctx kratosx.Context, options *types.PageOptions) ([]*Template, uint32, error) {
	var list []*Template
	total := int64(0)

	db := ctx.DB().Model(t)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Order("created_at desc").Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

func (t *Template) UseVersionByID(ctx kratosx.Context, srvId, id uint32) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(Template{}).Where("server_id=? and id=?", srvId, t.ID).Update("is_use", true).Error; err != nil {
			return err
		}
		return tx.Model(Template{}).Where("server_id=? and id!=?", srvId, t.ID).Update("is_use", false).Error
	})
}
