package model

import (
	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"
)

type BusinessValue struct {
	BaseModel
	EnvironmentID uint32      `json:"environment_id"`
	BusinessID    uint32      `json:"business_id"`
	Value         string      `json:"value"`
	Environment   Environment `json:"environment"`
	Business      Business    `json:"business"`
}

// Create 新建业务字段
func (bv *BusinessValue) Create(ctx kratosx.Context) error {
	return ctx.DB().Create(bv).Error
}

// Creates 批量创建资源
func (rv *BusinessValue) Creates(ctx kratosx.Context, rid uint32, list []*BusinessValue) error {
	db := ctx.DB()
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(BusinessValue{}, "business_id=?", rid).Error; err != nil {
			return err
		}
		return tx.Model(BusinessValue{}).Create(list).Error
	})
}

// OneByID 通过关键词查找指定业务字段
func (bv *BusinessValue) OneByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().First(bv, "id = ?", id).Error
}

// All 查询所有业务字段
func (bv *BusinessValue) All(ctx kratosx.Context, scopes Scopes) ([]*BusinessValue, error) {
	var list []*BusinessValue

	db := ctx.DB().Model(bv)
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的业务字段
func (bv *BusinessValue) Update(ctx kratosx.Context) error {
	return ctx.DB().Updates(bv).Error
}

// DeleteByID 删除指定id的业务字段
func (bv *BusinessValue) DeleteByID(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(bv, "id = ?", id).Error
}

func (bv *BusinessValue) AllByEnvAndServer(ctx kratosx.Context, envId, srvId uint32) ([]*BusinessValue, error) {
	return bv.All(ctx, func(db *gorm.DB) *gorm.DB {
		db.Preload("Business")
		db = db.Where("business_id in (select id from business where server_id = ?)", srvId)
		return db.Where("environment_id = ?", envId)
	})
}
