package model

import (
	"github.com/limes-cloud/kratos"
	"gorm.io/gorm"
)

type Template struct {
	BaseModel
	ServerID    int64   `json:"server_id"`
	Content     string  `json:"content"`
	Version     string  `json:"version"`
	IsUse       bool    `json:"is_use"`
	Format      string  `json:"format"`
	Description *string `json:"description"`
	Operator    string  `json:"operator"`
	OperatorID  int64   `json:"operator_id"`
}

func (t *Template) Create(ctx kratos.Context) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(t).Create(&t).Error; err != nil {
			return err
		}
		temp := Template{}
		return tx.Model(&temp).Where("server_id=? and id!=?", t.ServerID, t.ID).Update("is_use", false).Error
	})
}

func (t *Template) OneById(ctx kratos.Context, id int64) error {
	return ctx.DB().First(t, "id = ?", id).Error
}

func (t *Template) Current(ctx kratos.Context, srvId int64) error {
	return ctx.DB().First(t, "is_use = true and server_id = ?", srvId).Error
}

// Page 查询分页资源
func (t *Template) Page(ctx kratos.Context, options *PageOptions) ([]*Template, int64, error) {
	var list []*Template
	total := int64(0)

	db := ctx.DB().Model(t)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, total, db.Find(&list).Error
}

func (t *Template) UseVersionByID(ctx kratos.Context, srvId, id int64) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("server_id=? and id=?", srvId, t.ID).Update("is_use", true).Error; err != nil {
			return err
		}
		return tx.Where("server_id=? and id!=?", t.ID).Update("is_use", false).Error
	})
}
