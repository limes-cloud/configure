package model

import (
	"github.com/limes-cloud/kratos"
)

type Resource struct {
	BaseModel
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Fields      string `json:"fields"`
	Type        string `json:"type"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
}

// Create 新建资源
func (e *Resource) Create(ctx kratos.Context) error {
	return ctx.DB().Model(e).Create(e).Error
}

// OneByID 通过关键词查找指定资源
func (e *Resource) OneByID(ctx kratos.Context, id int64) error {
	return ctx.DB().First(e, "id = ?", id).Error
}

// Page 查询分页资源
func (e *Resource) Page(ctx kratos.Context, options *PageOptions) ([]*Resource, error) {
	var list []*Resource

	db := ctx.DB().Model(e)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, db.Find(&list).Error
}

// All 查询全部资源
func (e *Resource) All(ctx kratos.Context, scopes Scopes) ([]*Resource, error) {
	var list []*Resource

	db := ctx.DB().Model(e)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源
func (e *Resource) Update(ctx kratos.Context) error {
	return ctx.DB().Model(e).Updates(e).Error
}

// DeleteByID 删除指定id的资源
func (e *Resource) DeleteByID(ctx kratos.Context, id int64) error {
	return ctx.DB().Model(e).Delete(e, "id = ?", id).Error
}
