package model

import (
	"github.com/limes-cloud/kratos"
)

type Business struct {
	BaseModel
	ServerID    uint32 `json:"server_id"`
	Keyword     string `json:"keyword"`
	Type        string `json:"type"`
	Option      string `json:"option"`
	Description string `json:"description"`
}

// Create 新建业务字段
func (b *Business) Create(ctx kratos.Context) error {
	return ctx.DB().Model(b).Create(b).Error
}

// OneByID 通过关键词查找指定业务字段
func (b *Business) OneByID(ctx kratos.Context, id uint32) error {
	return ctx.DB().First(b, "id = ?", id).Error
}

// Page 查询分页资源
func (b *Business) Page(ctx kratos.Context, options *PageOptions) ([]*Business, uint32, error) {
	var list []*Business
	total := int64(0)

	db := ctx.DB().Model(b)
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// All 查询所有业务字段
func (b *Business) All(ctx kratos.Context, scopes Scopes) ([]*Business, error) {
	var list []*Business

	db := ctx.DB().Model(b)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的业务字段
func (b *Business) Update(ctx kratos.Context) error {
	return ctx.DB().Model(b).Updates(b).Error
}

// DeleteByID 删除指定id的业务字段
func (b *Business) DeleteByID(ctx kratos.Context, id uint32) error {
	return ctx.DB().Model(b).Delete(b, "id = ?", id).Error
}
