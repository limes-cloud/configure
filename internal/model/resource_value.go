package model

import (
	"github.com/limes-cloud/kratos"
)

type ResourceValue struct {
	BaseModel
	EnvKeyword string `json:"env_keyword"`
	ResourceID int64  `json:"resource_id"`
	Values     string `json:"values"`
	Operator   string `json:"operator,omitempty"`
	OperatorID int64  `json:"operator_id,omitempty"`
}

// Create 新建资源
func (e *ResourceValue) Create(ctx kratos.Context) error {
	return ctx.DB().Model(e).Create(e).Error
}

// Creates 批量创建资源
func (e *ResourceValue) Creates(ctx kratos.Context, list []*ResourceValue) error {
	return ctx.DB().Model(e).Create(list).Error
}

// All 查询全部资源
func (e *ResourceValue) All(ctx kratos.Context, scopes Scopes) ([]*ResourceValue, error) {
	var list []*ResourceValue

	db := ctx.DB().Model(e)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, ctx.DB().Model(e).Find(&list).Error
}
