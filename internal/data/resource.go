package data

import (
	"fmt"
	"strings"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	biz "github.com/limes-cloud/configure/internal/biz/resource"
	"github.com/limes-cloud/configure/internal/data/model"
)

type resourceRepo struct {
}

func NewResourceRepo() biz.Repo {
	return &resourceRepo{}
}

// ToResourceEntity model转entity
func (r resourceRepo) ToResourceEntity(m *model.Resource) *biz.Resource {
	e := &biz.Resource{}
	_ = valx.Transform(m, e)
	return e
}

// ToResourceModel entity转model
func (r resourceRepo) ToResourceModel(e *biz.Resource) *model.Resource {
	m := &model.Resource{}
	_ = valx.Transform(e, m)
	return m
}

// GetResource 获取指定的数据
func (r resourceRepo) GetResource(ctx kratosx.Context, id uint32) (*biz.Resource, error) {
	var (
		m  = model.Resource{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	db = db.Preload("Servers")
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToResourceEntity(&m), nil
}

// ListResource 获取列表
func (r resourceRepo) ListResource(ctx kratosx.Context, req *biz.ListResourceRequest) ([]*biz.Resource, uint32, error) {
	var (
		bs    []*biz.Resource
		ms    []*model.Resource
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Resource{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Tag != nil {
		db = db.Where("tag = ?", *req.Tag)
	}
	if req.Private != nil {
		db = db.Where("private = ?", *req.Private)
	}
	if req.ServerId != nil {
		var ids []uint32
		if err := ctx.DB().Select("resource_id").Model(model.ResourceServer{}).Where("server_id=?", *req.ServerId).Scan(&ids).Error; err != nil {
			return nil, 0, err
		}
		db = db.Where("id in ? or private=false", ids)
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

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToResourceEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateResource 创建数据
func (r resourceRepo) CreateResource(ctx kratosx.Context, req *biz.Resource) (uint32, error) {
	m := r.ToResourceModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateResource 更新数据
func (r resourceRepo) UpdateResource(ctx kratosx.Context, req *biz.Resource) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("resource_id=?", req.Id).Delete(model.ResourceServer{}).Error; err != nil {
			return err
		}
		return tx.Updates(r.ToResourceModel(req)).Error
	})
}

// DeleteResource 删除数据
func (r resourceRepo) DeleteResource(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Resource{})
	return uint32(db.RowsAffected), db.Error
}

// GetResourceByKeyword 获取指定数据
func (r resourceRepo) GetResourceByKeyword(ctx kratosx.Context, keyword string) (*biz.Resource, error) {
	var (
		m  = model.Resource{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	db = db.Preload("Servers")
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToResourceEntity(&m), nil
}

// ToResourceValueEntity model转entity
func (r resourceRepo) ToResourceValueEntity(m *model.ResourceValue) *biz.ResourceValue {
	e := &biz.ResourceValue{}
	_ = valx.Transform(m, e)
	return e
}

// ToResourceValueModel entity转model
func (r resourceRepo) ToResourceValueModel(e *biz.ResourceValue) *model.ResourceValue {
	m := &model.ResourceValue{}
	_ = valx.Transform(e, m)
	return m
}

// ListResourceValue 获取列表
func (r resourceRepo) ListResourceValue(ctx kratosx.Context, bid uint32) ([]*biz.ResourceValue, uint32, error) {
	var (
		bs []*biz.ResourceValue
		ms []*model.ResourceValue
		fs = []string{"*"}
	)

	db := ctx.DB().Model(model.ResourceValue{}).Select(fs).Where("resource_id = ?", bid)

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToResourceValueEntity(m))
	}
	return bs, uint32(len(bs)), nil
}

// UpdateResourceValues 更新数据
func (r resourceRepo) UpdateResourceValues(ctx kratosx.Context, bs []*biz.ResourceValue) error {
	var (
		ms []*model.ResourceValue
	)

	for _, b := range bs {
		ms = append(ms, r.ToResourceValueModel(b))
	}
	return ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(&ms).Error
}

// AllResourceField 获取指定服务的全部可用资源字段
func (r resourceRepo) AllResourceField(ctx kratosx.Context, sid uint32) ([]string, error) {
	var (
		ids  []uint32
		list []*model.Resource
	)
	if err := ctx.DB().Select("resource_id").
		Model(model.ResourceServer{}).
		Where("server_id=?", sid).
		Scan(&ids).Error; err != nil {
		return nil, err
	}

	if err := ctx.DB().Model(model.Resource{}).
		Select("fields", "keyword").
		Where("id in ? or private=false", ids).
		Find(&list).Error; err != nil {
		return nil, err
	}

	var arr []string
	for _, item := range list {
		fields := strings.Split(item.Fields, ",")
		for _, val := range fields {
			arr = append(arr, item.Keyword+"."+val)
		}
	}
	return arr, nil
}

// AllResourceValue 获取指定环境资源的所有值
func (r resourceRepo) AllResourceValue(ctx kratosx.Context, envId, srvId uint32) ([]*biz.ResourceValue, error) {
	var ids []uint32
	if err := ctx.DB().Select("resource_id").Model(model.ResourceServer{}).Where("server_id=?", srvId).Scan(&ids).Error; err != nil {
		return nil, err
	}
	var ids1 []uint32
	if err := ctx.DB().Select("id").Model(model.Resource{}).Where("private=false").Scan(&ids1).Error; err != nil {
		return nil, err
	}
	ids = append(ids, ids1...)

	var (
		ms []*model.ResourceValue
		bs []*biz.ResourceValue
	)
	if err := ctx.DB().Model(model.ResourceValue{}).
		Preload("Resource").
		Find(&ms, "env_id=? and resource_id in ?", envId, ids).
		Error; err != nil {
		return nil, err
	}
	for _, item := range ms {
		bs = append(bs, r.ToResourceValueEntity(item))
	}
	return bs, nil
}
