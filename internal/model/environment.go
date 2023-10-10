package model

import (
	"github.com/limes-cloud/kratos"
)

type Environment struct {
	BaseModel
	Keyword     string `json:"keyword" gorm:"unique;not null;size:32;comment:环境关键字"`
	Title       string `json:"title" gorm:"not null;size:32;comment:环境标题"`
	Description string `json:"description" gorm:"not null;size:128;comment:环境描述"`
	Token       string `json:"token,omitempty" gorm:"not null;size:32;comment:连接token"`
	Status      *bool  `json:"status" gorm:"not null;comment:启用状态"`
	Operator    string `json:"operator,omitempty" gorm:"not null;size:32;comment:操作人"`
	OperatorID  int64  `json:"operator_id,omitempty" gorm:"not null;comment:操作人id"`
}

// Create 新建环境
func (e *Environment) Create(ctx kratos.Context) error {
	return ctx.DB().Model(e).Create(e).Error
}

// OneByID 通过关键词查找指定环境
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
