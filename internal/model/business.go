package model

import "github.com/limes-cloud/kratos"

type Business struct {
	BaseModel
	ServerID    string `json:"server_id"`
	Keyword     string `json:"keyword"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Operator    string `json:"operator"`
	OperatorID  int64  `json:"operator_id"`
}

// Create 新建业务字段
func (e *Business) Create(ctx kratos.Context) error {
	return ctx.DB().Model(e).Create(e).Error
}

// OneByID 通过关键词查找指定业务字段
func (e *Business) OneByID(ctx kratos.Context, id int64) error {
	return ctx.DB().First(e, "id = ?", id).Error
}

// All 查询所有业务字段
func (e *Business) All(ctx kratos.Context) ([]*Business, error) {
	var list []*Business
	return list, ctx.DB().Model(e).Find(&list).Error
}

// UpdateByID 更新指定id的业务字段
func (e *Business) UpdateByID(ctx kratos.Context, id int64) error {
	return ctx.DB().Model(e).Where("id = ?", id).Updates(e).Error
}

// DeleteByID 删除指定id的业务字段
func (e *Business) DeleteByID(ctx kratos.Context, id int64) error {
	return ctx.DB().Model(e).Delete(e, "id = ?", id).Error
}
