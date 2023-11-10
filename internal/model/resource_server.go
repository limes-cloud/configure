package model

import (
	"github.com/limes-cloud/kratos"
	"gorm.io/gorm"
)

type ResourceServer struct {
	ID         uint32   `json:"id"`
	ServerID   uint32   `json:"server_id"`
	ResourceID uint32   `json:"resource_id"`
	Resource   Resource `json:"resource"`
	Server     Server   `json:"server"`
}

// CreateBySrvIds 新建资源服务
func (rs *ResourceServer) CreateBySrvIds(ctx kratos.Context, rid uint32, srvIds []uint32) error {
	var list []*ResourceServer
	for _, srvId := range srvIds {
		list = append(list, &ResourceServer{
			ResourceID: rid,
			ServerID:   srvId,
		})
	}

	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(ResourceServer{}, "resource_id=?", rid).Error; err != nil {
			return err
		}
		return tx.Model(ResourceServer{}).Create(list).Error
	})
}

// All 查询全部资源
func (rs *ResourceServer) All(ctx kratos.Context, scopes Scopes) ([]*ResourceServer, error) {
	var list []*ResourceServer

	db := ctx.DB().Model(rs)
	if scopes != nil {
		db = db.Scopes(scopes)
	}
	return list, db.Find(&list).Error
}

// DeleteByID 删除指定id的资源
func (rs *ResourceServer) DeleteByID(ctx kratos.Context, id uint32) error {
	return ctx.DB().Model(rs).Delete(rs, "id = ?", id).Error
}
