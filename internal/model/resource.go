package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type Resource struct {
	types.BaseModel
	Keyword        string            `json:"keyword" gorm:"not null;type:char(32) binary;comment:变量标识"`
	Description    string            `json:"description" gorm:"not null;size:128;comment:变量描述"`
	Fields         string            `json:"fields" gorm:"not null;size:256;comment:变量字段集合"`
	Tag            string            `json:"tag" gorm:"not null;size:32;comment:变量标签"`
	Private        *bool             `json:"private" gorm:"default:false;comment:是否私有"`
	ResourceServer []*ResourceServer `json:"-" gorm:"-"`
	ResourceValue  []*ResourceValue  `json:"-" gorm:"-"`
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
func (r *Resource) Page(ctx kratosx.Context, options *types.PageOptions) ([]*Resource, uint32, error) {
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
func (r *Resource) All(ctx kratosx.Context, scopes types.Scopes) ([]*Resource, error) {
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
