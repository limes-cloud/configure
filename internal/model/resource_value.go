package model

import (
	"github.com/limes-cloud/kratos"
	"gorm.io/gorm"
)

type ResourceValue struct {
	BaseModel
	EnvironmentID int64       `json:"environment_id"`
	ResourceID    int64       `json:"resource_id"`
	Values        string      `json:"values"`
	Environment   Environment `json:"environment"`
	Resource      Resource    `json:"resource"`
}

// Create 新建资源
func (rv *ResourceValue) Create(ctx kratos.Context) error {
	return ctx.DB().Model(rv).Create(rv).Error
}

// Creates 批量创建资源
func (rv *ResourceValue) Creates(ctx kratos.Context, list []*ResourceValue) error {
	return ctx.DB().Model(rv).Create(list).Error
}

// All 查询全部资源
func (rv *ResourceValue) All(ctx kratos.Context, scopes Scopes) ([]*ResourceValue, error) {
	var list []*ResourceValue

	db := ctx.DB().Model(rv)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源值
func (rv *ResourceValue) Update(ctx kratos.Context) error {
	return ctx.DB().Model(rv).Updates(rv).Error
}

func (rv *ResourceValue) AllByEnvAndServer(ctx kratos.Context, envId, srvId int64) ([]*ResourceValue, error) {
	return rv.All(ctx, func(db *gorm.DB) *gorm.DB {
		db.Preload("Resource")
		db = db.Where("resource_id in (select resource_id from resource_server where server_id=?)", srvId)
		return db.Where("environment_id = ?", envId)
	})
}
