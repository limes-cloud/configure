package model

import "github.com/limes-cloud/kratosx"

type Resource struct {
	BaseModel
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Fields      string `json:"fields"`
	Tag         string `json:"tag"`
	Private     *bool  `json:"private"`
}

// Create 新建资源
func (r *Resource) Create(ctx kratosx.Context) error {
	return ctx.DB().Model(r).Create(r).Error
}

// OneByID 通过关键词查找指定资源
func (r *Resource) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(r, "id = ?", id).Error
}

// Page 查询分页资源
func (r *Resource) Page(ctx kratosx.Context, options *PageOptions) ([]*Resource, uint32, error) {
	var list []*Resource
	total := int64(0)

	db := ctx.DB().Model(r)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// All 查询全部资源
func (r *Resource) All(ctx kratosx.Context, scopes Scopes) ([]*Resource, error) {
	var list []*Resource

	db := ctx.DB().Model(r)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源
func (r *Resource) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(r).Updates(r).Error
}

// DeleteByID 删除指定id的资源
func (r *Resource) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Model(r).Delete(r, "id = ?", id).Error
}
