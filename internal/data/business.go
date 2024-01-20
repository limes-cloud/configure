package data

import (
	"errors"

	"github.com/limes-cloud/configure/internal/biz/types"

	"gorm.io/gorm"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type businessRepo struct {
}

func NewBusinessRepo() biz.BusinessRepo {

	return &businessRepo{}
}

// Create 新建业务变量
func (s *businessRepo) Create(ctx kratosx.Context, business *biz.Business) (uint32, error) {
	return business.ID, ctx.DB().Create(business).Error
}

// Get 通过ID查找指定业务变量
func (s *businessRepo) Get(ctx kratosx.Context, id uint32) (*biz.Business, error) {
	var business biz.Business
	return &business, ctx.DB().First(&business, "id=?", id).Error
}

// GetByKeyword 通过关键词查找指定业务变量
func (s *businessRepo) GetByKeyword(ctx kratosx.Context, keyword string) (*biz.Business, error) {
	var business biz.Business
	return &business, ctx.DB().First(&business, "keyword=?", keyword).Error
}

// Page 查询分页业务变量
func (s *businessRepo) Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*biz.Business, uint32, error) {
	var list []*biz.Business
	var total int64

	db := ctx.DB().Model(biz.Business{})
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// PageServerBusiness 查询指定服务的全部业务变量
func (s *businessRepo) PageServerBusiness(ctx kratosx.Context, req *types.PageServerBusinessRequest) ([]*biz.Business, uint32, error) {
	return s.Page(ctx, &ktypes.PageOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			db = db.Where("server_id=?", req.ServerID)
			if req.Keyword != nil {
				db = db.Where("keyword like ?", "%"+*req.Keyword+"%")
			}
			if req.Tag != nil {
				db = db.Where("tag like ?", "%"+*req.Tag+"%")
			}
			return db
		},
	})
}

// AllServerBusiness 查询指定服务的全部业务变量
func (s *businessRepo) AllServerBusiness(ctx kratosx.Context, srvId uint32) ([]*biz.Business, error) {
	var list []*biz.Business
	return list, ctx.DB().Model(s).Where("server_id=?", srvId).Find(&list).Error
}

// Update 更新指定id的业务变量
func (s *businessRepo) Update(ctx kratosx.Context, business *biz.Business) error {
	return ctx.DB().Where("id=?", business.ID).Updates(business).Error
}

// Delete 删除指定id的业务变量
func (s *businessRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Business{}, "id = ?", id).Error
}

// GetValues 获取指定业务变量的所有值
func (s *businessRepo) GetValues(ctx kratosx.Context, bid uint32) ([]*biz.BusinessValue, error) {
	var list []*biz.BusinessValue
	return list, ctx.DB().Model(biz.BusinessValue{}).Find(&list, "business_id = ?", bid).Error
}

// GetEnvValues 获取指定环境的变量的所有值
func (s *businessRepo) GetEnvValues(ctx kratosx.Context, eid, sid uint32) ([]*biz.BusinessValue, error) {
	// 获取指定服务的全部业务字段
	var ids []uint32
	if err := ctx.DB().Select("id").Model(biz.Business{}).Where("server_id=?", sid).Scan(&ids).Error; err != nil {
		return nil, err
	}

	var list []*biz.BusinessValue
	return list, ctx.DB().Model(biz.BusinessValue{}).Preload("Business").Find(&list, "env_id=? and business_id in ?", eid, ids).Error
}

// UpdateValue 更新指定业务变量的值
func (s *businessRepo) UpdateValue(ctx kratosx.Context, rv *biz.BusinessValue) error {
	if rv.ID == 0 {
		oldRv := &biz.BusinessValue{}
		if err := ctx.DB().First(oldRv, "business_id=? and env_id=?", rv.BusinessID, rv.EnvID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.DB().Create(rv).Error
			}
			return err
		}
		rv.ID = oldRv.ID
	}
	return ctx.DB().Updates(rv).Error
}
