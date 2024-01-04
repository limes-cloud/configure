package model

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
	"gorm.io/gorm"
)

type ResourceValue struct {
	types.BaseModel
	EnvironmentID uint32      `json:"environment_id" gorm:"not null;comment:环境id"`
	ResourceID    uint32      `json:"resource_id" gorm:"not null;comment:资源id"`
	Values        string      `json:"values" gorm:"not null;type:text;comment:资源id"`
	Environment   Environment `json:"environment" gorm:"constraint:onDelete:cascade"`
	Resource      Resource    `json:"resource" gorm:"constraint:onDelete:cascade"`
}

// Create 新建资源
func (rv *ResourceValue) Create(ctx kratosx.Context) error {
	return ctx.DB().Model(rv).Create(rv).Error
}

// Creates 批量创建资源
func (rv *ResourceValue) Creates(ctx kratosx.Context, rid uint32, list []*ResourceValue) error {
	db := ctx.DB()
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(ResourceValue{}, "resource_id=?", rid).Error; err != nil {
			return err
		}
		return tx.Model(ResourceValue{}).Create(list).Error
	})
}

// All 查询全部资源
func (rv *ResourceValue) All(ctx kratosx.Context, scopes types.Scopes) ([]*ResourceValue, error) {
	var list []*ResourceValue

	db := ctx.DB().Model(rv)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源值
func (rv *ResourceValue) Update(ctx kratosx.Context) error {
	return ctx.DB().Model(rv).Updates(rv).Error
}

func (rv *ResourceValue) AllByEnvAndServer(ctx kratosx.Context, envId, srvId uint32) ([]*ResourceValue, error) {
	return rv.All(ctx, func(db *gorm.DB) *gorm.DB {
		db.Preload("Resource")
		db = db.Where("resource_id in (select resource_id from resource_server where server_id=? "+
			"union select id from resource where private=false)", srvId)
		return db.Where("environment_id = ?", envId)
	})
}
