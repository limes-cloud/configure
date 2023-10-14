package model

import "github.com/limes-cloud/kratos"

type Server struct {
	BaseModel
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
}

// Create 新建资源
func (e *Server) Create(ctx kratos.Context) error {
	return ctx.DB().Model(e).Create(e).Error
}

// OneByID 通过关键词查找指定资源
func (e *Server) OneByID(ctx kratos.Context, id int64) error {
	return ctx.DB().First(e, "id = ?", id).Error
}

func (e *Server) OneByKeyword(ctx kratos.Context, keyword string) error {
	return ctx.DB().First(e, "keyword = ?", keyword).Error
}

// Page 查询分页资源
func (e *Server) Page(ctx kratos.Context, options *PageOptions) ([]*Server, int64, error) {
	var list []*Server
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

// All 查询全部资源
func (e *Server) All(ctx kratos.Context, scopes Scopes) ([]*Server, error) {
	var list []*Server

	db := ctx.DB().Model(e)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源
func (e *Server) Update(ctx kratos.Context) error {
	return ctx.DB().Model(e).Updates(e).Error
}

// DeleteByID 删除指定id的资源
func (e *Server) DeleteByID(ctx kratos.Context, id int64) error {
	return ctx.DB().Model(e).Delete(e, "id = ?", id).Error
}
