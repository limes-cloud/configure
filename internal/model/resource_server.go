package model

import "github.com/limes-cloud/kratos"

type ResourceServer struct {
	CreateModel
	ServerID   int64  `json:"server_id"`
	ResourceID int64  `json:"resource_id"`
	Operator   string `json:"operator"`
	OperatorID int64  `json:"operator_id"`
}

// Create 新建资源
func (e *ResourceServer) Create(ctx kratos.Context) error {
	return ctx.DB().Model(e).Create(e).Error
}

// All 查询全部资源
func (e *ResourceServer) All(ctx kratos.Context, scopes Scopes) ([]*ResourceServer, error) {
	var list []*ResourceServer

	db := ctx.DB().Model(e)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// DeleteByID 删除指定id的资源
func (e *ResourceServer) DeleteByID(ctx kratos.Context, id int64) error {
	return ctx.DB().Model(e).Delete(e, "id = ?", id).Error
}
