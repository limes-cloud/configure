package business

import (
	"errors"

	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/configure/internal/biz/business"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

// AddBusiness 新建业务变量
func (s *repo) AddBusiness(ctx kratosx.Context, business *biz.Business) (uint32, error) {
	return business.ID, ctx.DB().Create(business).Error
}

// GetBusiness 通过ID查找指定业务变量
func (s *repo) GetBusiness(ctx kratosx.Context, id uint32) (*biz.Business, error) {
	var business biz.Business
	return &business, ctx.DB().First(&business, "id=?", id).Error
}

// GetBusinessByKeyword 通过关键词查找指定业务变量
func (s *repo) GetBusinessByKeyword(ctx kratosx.Context, keyword string) (*biz.Business, error) {
	var business biz.Business
	return &business, ctx.DB().First(&business, "keyword=?", keyword).Error
}

// PageBusiness 查询分页业务变量
func (s *repo) PageBusiness(ctx kratosx.Context, req *biz.PageBusinessRequest) ([]*biz.Business, uint32, error) {
	var list []*biz.Business
	var total int64

	db := ctx.DB().Model(biz.Business{}).Where("server_id=?", req.ServerId)
	if req.Keyword != nil {
		db = db.Where("keyword like ?", "%"+*req.Keyword+"%")
	}
	if req.Tag != nil {
		db = db.Where("tag like ?", "%"+*req.Tag+"%")
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// UpdateBusiness 更新指定id的业务变量
func (s *repo) UpdateBusiness(ctx kratosx.Context, business *biz.Business) error {
	return ctx.DB().Where("id=?", business.ID).Updates(business).Error
}

// DeleteBusiness 删除指定id的业务变量
func (s *repo) DeleteBusiness(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Business{}, "id = ?", id).Error
}

// GetBusinessValues 获取指定业务变量的所有值
func (s *repo) GetBusinessValues(ctx kratosx.Context, bid uint32) ([]*biz.BusinessValue, error) {
	var list []*biz.BusinessValue
	return list, ctx.DB().Model(biz.BusinessValue{}).Find(&list, "business_id = ?", bid).Error
}

// UpdateBusinessValue 更新指定业务变量的值
func (s *repo) UpdateBusinessValue(ctx kratosx.Context, rv *biz.BusinessValue) error {
	if rv.ID == 0 {
		oldRv := &biz.BusinessValue{}
		if err := ctx.DB().First(oldRv, "business_id=? and env_id=?", rv.BusinessId, rv.EnvId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.DB().Create(rv).Error
			}
			return err
		}
		rv.ID = oldRv.ID
	}
	return ctx.DB().Updates(rv).Error
}

// AllBusinessValue 获取指定环境的变量的所有值
func (s *repo) AllBusinessValue(ctx kratosx.Context, eid, sid uint32) ([]*biz.BusinessValue, error) {
	// 获取指定服务的全部业务字段
	var ids []uint32
	if err := ctx.DB().Select("id").Model(biz.Business{}).Where("server_id=?", sid).Scan(&ids).Error; err != nil {
		return nil, err
	}

	var list []*biz.BusinessValue
	return list, ctx.DB().
		Model(biz.BusinessValue{}).
		Preload("Business").
		Where("env_id=? and business_id in ?", eid, ids).
		Find(&list).Error
}

// AllBusinessField 查询指定服务的全部业务变量
func (s *repo) AllBusinessField(ctx kratosx.Context, srvId uint32) ([]string, error) {
	var list []string
	return list, ctx.DB().Select("keyword").
		Model(biz.Business{}).
		Where("server_id=?", srvId).
		Find(&list).Error
}

func (s *repo) CheckBusinessValue(ctx kratosx.Context, id uint32, value string) bool {
	// business, err := s.GetBusiness(ctx, id)
	// if err != nil {
	//	return false
	// }
	//
	return true
}
