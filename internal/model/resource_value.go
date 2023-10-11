package model

import (
	"github.com/limes-cloud/kratos"
)

type ResourceValue struct {
	BaseModel
	EnvironmentID int64       `json:"environment_id"`
	ResourceID    int64       `json:"resource_id"`
	Values        string      `json:"values"`
	Operator      string      `json:"operator"`
	OperatorID    int64       `json:"operator_id"`
	Environment   Environment `json:"environment"`
	Resource      Resource    `json:"resource"`
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

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源值
func (e *ResourceValue) Update(ctx kratos.Context) error {
	return ctx.DB().Model(e).Updates(e).Error
}
