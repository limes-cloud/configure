package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type Server struct {
	types.BaseModel
	Keyword     string `json:"keyword" gorm:"not null;type:char(32) binary;comment:服务标识"`
	Name        string `json:"name" gorm:"not null;size:64;comment:服务名称"`
	Description string `json:"description" gorm:"not null;size:128;comment:服务描述"`
}

// Create 新建资源
func (s *Server) Create(ctx kratosx.Context) error {
	return ctx.DB().Model(s).Create(s).Error
}

// OneByID 通过关键词查找指定资源
func (s *Server) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(s, "id = ?", id).Error
}

func (s *Server) OneByKeyword(ctx kratosx.Context, keyword string) error {
	return ctx.DB().First(s, "keyword = ?", keyword).Error
}

// Page 查询分页资源
func (s *Server) Page(ctx kratosx.Context, options *types.PageOptions) ([]*Server, uint32, error) {
	var list []*Server
	total := int64(0)

	db := ctx.DB().Model(s)
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
func (s *Server) All(ctx kratosx.Context, scopes types.Scopes) ([]*Server, error) {
	var list []*Server

	db := ctx.DB().Model(s)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源
func (s *Server) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(s).Updates(s).Error
}

// DeleteByID 删除指定id的资源
func (s *Server) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Model(s).Delete(s, "id = ?", id).Error
}
