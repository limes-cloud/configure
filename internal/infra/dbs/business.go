package dbs

import (
	"fmt"
	"sync"

	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"

	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type BusinessInfra struct {
}

var (
	businessIns  *BusinessInfra
	businessOnce sync.Once
)

func NewBusinessInfra() *BusinessInfra {
	businessOnce.Do(func() {
		businessIns = &BusinessInfra{}
	})
	return businessIns
}

// GetBusiness 获取指定的数据
func (r BusinessInfra) GetBusiness(ctx kratosx.Context, id uint32) (*entity.Business, error) {
	var (
		business = entity.Business{}
		fs       = []string{"*"}
	)
	return &business, ctx.DB().Select(fs).First(&business, id).Error
}

// ListBusiness 获取列表
func (r BusinessInfra) ListBusiness(ctx kratosx.Context, req *types.ListBusinessRequest) ([]*entity.Business, uint32, error) {
	var (
		bs    []*entity.Business
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(entity.Business{}).Select(fs).Where("server_id = ?", req.ServerId)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if req.OrderBy == nil || *req.OrderBy == "" {
		req.OrderBy = proto.String("id")
	}
	if req.Order == nil || *req.Order == "" {
		req.Order = proto.String("asc")
	}
	db = db.Order(fmt.Sprintf("%s %s", *req.OrderBy, *req.Order))
	if *req.OrderBy != "id" {
		db = db.Order("id asc")
	}

	return bs, uint32(total), db.Find(&bs).Error
}

// CreateBusiness 创建数据
func (r BusinessInfra) CreateBusiness(ctx kratosx.Context, business *entity.Business) (uint32, error) {
	return business.Id, ctx.DB().Create(business).Error
}

// UpdateBusiness 更新数据
func (r BusinessInfra) UpdateBusiness(ctx kratosx.Context, business *entity.Business) error {
	return ctx.DB().Where("id = ?", business.Id).Updates(business).Error
}

// DeleteBusiness 删除数据
func (r BusinessInfra) DeleteBusiness(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id = ?", id).Delete(&entity.Business{}).Error
}

// ListBusinessValue 获取指定业务字段的值列表
func (r BusinessInfra) ListBusinessValue(ctx kratosx.Context, req *types.ListBusinessValueRequest) ([]*entity.BusinessValue, error) {
	var (
		bvs []*entity.BusinessValue
		fs  = []string{"*"}
	)

	db := ctx.DB().Model(entity.BusinessValue{}).Select(fs).Preload("Business")
	if req.BusinessId != nil {
		db = db.Where("business_id=?", *req.BusinessId)
	}
	if req.ServerId != nil {
		// 获取指定服务的全部业务字段
		var ids []uint32
		if err := ctx.DB().Select("id").Model(entity.Business{}).
			Where("server_id=?", *req.ServerId).Scan(&ids).Error; err != nil {
			return nil, err
		}
		db = db.Where("business_id in ?", ids)
	}
	if req.EnvId != nil {
		db = db.Where("env_id = ?", *req.EnvId)
	}
	return bvs, db.Find(&bvs).Error
}

// UpdateBusinessValues 更新数据
func (r BusinessInfra) UpdateBusinessValues(ctx kratosx.Context, bvs []*entity.BusinessValue) error {
	return ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(&bvs).Error
}

// AllBusinessField 获取指定服务的可用的业务字段key列表
func (r BusinessInfra) AllBusinessField(ctx kratosx.Context, srvId uint32) ([]string, error) {
	var list []string
	return list, ctx.DB().Select("keyword").
		Model(entity.Business{}).
		Where("server_id=?", srvId).
		Find(&list).Error
}
