package data

import (
	"errors"

	"github.com/limes-cloud/configure/internal/biz/types"

	"gorm.io/gorm"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type resourceRepo struct {
}

func NewResourceRepo() biz.ResourceRepo {
	return &resourceRepo{}
}

// Create 新建资源
func (s *resourceRepo) Create(ctx kratosx.Context, resource *biz.Resource) (uint32, error) {
	return resource.ID, ctx.DB().Create(resource).Error
}

// Get 通过ID查找指定资源
func (s *resourceRepo) Get(ctx kratosx.Context, id uint32) (*biz.Resource, error) {
	var resource biz.Resource
	return &resource, ctx.DB().First(&resource, "id=?", id).Error
}

// GetByKeyword 通过关键词查找指定资源
func (s *resourceRepo) GetByKeyword(ctx kratosx.Context, keyword string) (*biz.Resource, error) {
	var resource biz.Resource
	return &resource, ctx.DB().First(&resource, "keyword=?", keyword).Error
}

// Page 查询分页资源
func (s *resourceRepo) Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*biz.Resource, uint32, error) {
	var list []*biz.Resource
	var total int64

	db := ctx.DB().Model(biz.Resource{})
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// All 查询全部资源
func (s *resourceRepo) All(ctx kratosx.Context, scopes ktypes.Scopes) ([]*biz.Resource, error) {
	var list []*biz.Resource

	db := ctx.DB().Model(biz.Resource{})
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的资源
func (s *resourceRepo) Update(ctx kratosx.Context, resource *biz.Resource) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(biz.ResourceServer{}, "resource_id=?", resource.ID).Error; err != nil {
			return err
		}
		return tx.Where("id=?", resource.ID).Updates(resource).Error
	})
}

// Delete 删除指定id的资源
func (s *resourceRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Resource{}, "id = ?", id).Error
}

// PageResource 获取分页资源
func (s *resourceRepo) PageResource(ctx kratosx.Context, req *types.PageResourceRequest) ([]*biz.Resource, uint32, error) {
	return s.Page(ctx, &ktypes.PageOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
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

// PageServerResource 获取指定服务的分页资源
func (s *resourceRepo) PageServerResource(ctx kratosx.Context, id uint32, options *ktypes.PageOptions) ([]*biz.Resource, uint32, error) {
	var ids []uint32
	if err := ctx.DB().Select("resource_id").Model(biz.ResourceServer{}).Where("server_id=?", id).Scan(id).Error; err != nil {
		return nil, 0, err
	}

	options.Scopes = func(db *gorm.DB) *gorm.DB {
		return db.Where("id in ? or private=false", ids)
	}
	return s.Page(ctx, options)
}

// AllServerResource 获取指定服务的所有资源
func (s *resourceRepo) AllServerResource(ctx kratosx.Context, id uint32) ([]*biz.Resource, error) {
	var ids []uint32
	if err := ctx.DB().Select("resource_id").Model(biz.ResourceServer{}).Where("server_id=?", id).Scan(&ids).Error; err != nil {
		return nil, err
	}

	return s.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("id in ? or private=false", ids)
	})
}

// AllResourceServerId 获取指定资源的所有服务的id
func (s *resourceRepo) AllResourceServerId(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	if err := ctx.DB().Select("server_id").Model(biz.ResourceServer{}).Where("resource_id=?", id).Scan(&ids).Error; err != nil {
		return nil, err
	}

	return ids, nil
}

// GetValues 获取指定资源的所有值
func (s *resourceRepo) GetValues(ctx kratosx.Context, rid uint32) ([]*biz.ResourceValue, error) {
	var list []*biz.ResourceValue
	return list, ctx.DB().Model(biz.ResourceValue{}).Find(&list, "resource_id = ?", rid).Error
}

// GetEnvValues 获取指定环境资源的所有值
func (s *resourceRepo) GetEnvValues(ctx kratosx.Context, eid, sid uint32) ([]*biz.ResourceValue, error) {
	var ids []uint32
	if err := ctx.DB().Select("resource_id").Model(biz.ResourceServer{}).Where("server_id=?", sid).Scan(&ids).Error; err != nil {
		return nil, err
	}
	var ids1 []uint32
	if err := ctx.DB().Select("id").Model(biz.Resource{}).Where("private=false").Scan(&ids1).Error; err != nil {
		return nil, err
	}
	ids = append(ids, ids1...)

	var list []*biz.ResourceValue
	return list, ctx.DB().Model(biz.ResourceValue{}).Preload("Resource").Find(&list, "env_id=? and resource_id in ?", eid, ids).Error
}

// UpdateValue 更新指定资源的值
func (s *resourceRepo) UpdateValue(ctx kratosx.Context, rv *biz.ResourceValue) error {
	if rv.ID == 0 {
		oldRv := &biz.ResourceValue{}
		if err := ctx.DB().First(oldRv, "resource_id=? and env_id=?", rv.ResourceID, rv.EnvID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.DB().Create(rv).Error
			}
			return err
		}
		rv.ID = oldRv.ID
	}
	return ctx.DB().Updates(rv).Error
}
