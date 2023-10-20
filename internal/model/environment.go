package model

import (
	"github.com/limes-cloud/kratos"
)

type Environment struct {
	BaseModel
	Keyword     string `json:"keyword"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Token       string `json:"token,omitempty"`
	Status      *bool  `json:"status"`
}

// Create 新建环境
func (e *Environment) Create(ctx kratos.Context) error {
	return ctx.DB().Model(e).Create(e).Error
}

// OneByKeyword 通过关键词查找指定环境
func (e *Environment) OneByKeyword(ctx kratos.Context, keyword string) error {
	return ctx.DB().First(e, "keyword = ?", keyword).Error
}

// OneByToken 通过token查找指定环境
func (e *Environment) OneByToken(ctx kratos.Context, token string) error {
	return ctx.DB().First(e, "token = ?", token).Error
}

// OneByID 通过id查找指定环境
func (e *Environment) OneByID(ctx kratos.Context, id int64) error {
	return ctx.DB().First(e, "id = ?", id).Error
}

// All 查询所有环境
func (e *Environment) All(ctx kratos.Context) ([]*Environment, error) {
	var list []*Environment
	return list, ctx.DB().Model(e).Find(&list).Error
}

// UpdateByID 更新指定id的环境
func (e *Environment) UpdateByID(ctx kratos.Context, id int64) error {
	return ctx.DB().Model(e).Where("id = ?", id).Updates(e).Error
}

// DeleteByID 删除指定id的环境
func (e *Environment) DeleteByID(ctx kratos.Context, id int64) error {
	return ctx.DB().Model(e).Delete(e, "id = ?", id).Error
}
