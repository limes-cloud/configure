package model

import "github.com/limes-cloud/kratos"

type Server struct {
	BaseModel
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create 新建资源
func (s *Server) Create(ctx kratos.Context) error {
	return ctx.DB().Model(s).Create(s).Error
}

// OneByID 通过关键词查找指定资源
func (s *Server) OneByID(ctx kratos.Context, id uint32) error {
	return ctx.DB().First(s, "id = ?", id).Error
}

func (s *Server) OneByKeyword(ctx kratos.Context, keyword string) error {
	return ctx.DB().First(s, "keyword = ?", keyword).Error
}

// Page 查询分页资源
func (s *Server) Page(ctx kratos.Context, options *PageOptions) ([]*Server, uint32, error) {
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
func (s *Server) All(ctx kratos.Context, scopes Scopes) ([]*Server, error) {
	var list []*Server

	db := ctx.DB().Model(s)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源
func (s *Server) Update(ctx kratos.Context) error {
	return ctx.DB().Model(s).Updates(s).Error
}

// DeleteByID 删除指定id的资源
func (s *Server) DeleteByID(ctx kratos.Context, id uint32) error {
	return ctx.DB().Model(s).Delete(s, "id = ?", id).Error
}
