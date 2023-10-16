package model

import (
	"github.com/limes-cloud/kratos"
)

type Business struct {
	BaseModel
	ServerID    int64  `json:"server_id"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
}

// Create 新建业务字段
func (e *Business) Create(ctx kratos.Context) error {
	return ctx.DB().Model(e).Create(e).Error
}

// OneByID 通过关键词查找指定业务字段
func (e *Business) OneByID(ctx kratos.Context, id int64) error {
	return ctx.DB().First(e, "id = ?", id).Error
}

// Page 查询分页资源
func (e *Business) Page(ctx kratos.Context, options *PageOptions) ([]*Business, int64, error) {
	var list []*Business
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

// All 查询所有业务字段
func (e *Business) All(ctx kratos.Context, scopes Scopes) ([]*Business, error) {
	var list []*Business

	db := ctx.DB().Model(e)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的业务字段
func (e *Business) Update(ctx kratos.Context) error {
	return ctx.DB().Model(e).Updates(e).Error
}

// DeleteByID 删除指定id的业务字段
func (e *Business) DeleteByID(ctx kratos.Context, id int64) error {
	return ctx.DB().Model(e).Delete(e, "id = ?", id).Error
}
