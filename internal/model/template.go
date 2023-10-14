package model

import (
	"github.com/limes-cloud/kratos"
	"gorm.io/gorm"
)

type Template struct {
	BaseModel
	ServerID    string  `json:"server_id"`
	Content     string  `json:"content"`
	Version     string  `json:"version"`
	IsUse       bool    `json:"is_use"`
	Description *string `json:"description"`
	Operator    string  `json:"operator"`
	OperatorId  int64   `json:"operator_id"`
}

func (u *Template) Create(ctx kratos.Context) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&u).Error; err != nil {
			return err
		}
		return tx.Where("server_id=? and id!=?", u.ServerID, u.ID).Update("is_use", false).Error
	})
}

func (u *Template) OneById(ctx kratos.Context, id int64) error {
	return ctx.DB().First(u, "id = ?", id).Error
}

func (u *Template) One(ctx kratos.Context, conds ...interface{}) error {
	return ctx.DB().First(u, conds...).Error
}

func (u *Template) OneBy(ctx kratos.Context, conds ...interface{}) error {
	return ctx.DB().First(u, conds...).Error
}

// Page 查询分页资源
func (e *Template) Page(ctx kratos.Context, options *PageOptions) ([]*Template, int64, error) {
	var list []*Template
	total := int64(0)

	db := ctx.DB().Model(e)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, total, db.Find(&list).Error
}

func (u *Template) UseVersionByID(ctx kratos.Context, srvId, id int64) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("server_id=? and id=?", srvId, u.ID).Update("is_use", true).Error; err != nil {
			return err
		}
		return tx.Where("server_id=? and id!=?", u.ID).Update("is_use", false).Error
	})
}
