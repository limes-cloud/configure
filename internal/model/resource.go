package model

import (
	"github.com/limes-cloud/kratos"
)

type Resource struct {
	BaseModel
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Fields      string `json:"fields"`
	Tag         string `json:"tag"`
}

// Create 新建资源
func (r *Resource) Create(ctx kratos.Context) error {
	return ctx.DB().Model(r).Create(r).Error
}

// OneByID 通过关键词查找指定资源
func (r *Resource) OneByID(ctx kratos.Context, id int64) error {
	return ctx.DB().First(r, "id = ?", id).Error
}

// Page 查询分页资源
func (r *Resource) Page(ctx kratos.Context, options *PageOptions) ([]*Resource, int64, error) {
	var list []*Resource
	total := int64(0)

	db := ctx.DB().Model(r)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, total, err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, total, db.Find(&list).Error
}

// All 查询全部资源
func (r *Resource) All(ctx kratos.Context, scopes Scopes) ([]*Resource, error) {
	var list []*Resource

	db := ctx.DB().Model(r)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源
func (r *Resource) Update(ctx kratos.Context) error {
	return ctx.DB().Model(r).Updates(r).Error
}

// DeleteByID 删除指定id的资源
func (r *Resource) DeleteByID(ctx kratos.Context, id int64) error {
	return ctx.DB().Model(r).Delete(r, "id = ?", id).Error
}
