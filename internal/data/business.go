package data

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm/clause"

	biz "github.com/limes-cloud/configure/internal/biz/business"
	"github.com/limes-cloud/configure/internal/data/model"
)

type businessRepo struct {
}

func NewBusinessRepo() biz.Repo {
	return &businessRepo{}
}

// ToBusinessEntity model转entity
func (r businessRepo) ToBusinessEntity(m *model.Business) *biz.Business {
	e := &biz.Business{}
	_ = valx.Transform(m, e)
	return e
}

// ToBusinessModel entity转model
func (r businessRepo) ToBusinessModel(e *biz.Business) *model.Business {
	m := &model.Business{}
	_ = valx.Transform(e, m)
	return m
}

// GetBusiness 获取指定的数据
func (r businessRepo) GetBusiness(ctx kratosx.Context, id uint32) (*biz.Business, error) {
	var (
		m  = model.Business{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToBusinessEntity(&m), nil
}

// ListBusiness 获取列表
func (r businessRepo) ListBusiness(ctx kratosx.Context, req *biz.ListBusinessRequest) ([]*biz.Business, uint32, error) {
	var (
		bs    []*biz.Business
		ms    []*model.Business
		total int64
		fs    = []string{"*"}
	)

	db := ctx.DB().Model(model.Business{}).Select(fs)

	if req.ServerId != nil {
		db = db.Where("server_id = ?", *req.ServerId)
	}
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

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToBusinessEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateBusiness 创建数据
func (r businessRepo) CreateBusiness(ctx kratosx.Context, req *biz.Business) (uint32, error) {
	m := r.ToBusinessModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateBusiness 更新数据
func (r businessRepo) UpdateBusiness(ctx kratosx.Context, req *biz.Business) error {
	return ctx.DB().Updates(r.ToBusinessModel(req)).Error
}

// DeleteBusiness 删除数据
func (r businessRepo) DeleteBusiness(ctx kratosx.Context, ids []uint32) (uint32, error) {
	db := ctx.DB().Where("id in ?", ids).Delete(&model.Business{})
	return uint32(db.RowsAffected), db.Error
}

// ToBusinessValueEntity model转entity
func (r businessRepo) ToBusinessValueEntity(m *model.BusinessValue) *biz.BusinessValue {
	e := &biz.BusinessValue{}
	_ = valx.Transform(m, e)
	return e
}

// ToBusinessValueModel entity转model
func (r businessRepo) ToBusinessValueModel(e *biz.BusinessValue) *model.BusinessValue {
	m := &model.BusinessValue{}
	_ = valx.Transform(e, m)
	return m
}

// ListBusinessValue 获取列表
func (r businessRepo) ListBusinessValue(ctx kratosx.Context, bid uint32) ([]*biz.BusinessValue, uint32, error) {
	var (
		bs []*biz.BusinessValue
		ms []*model.BusinessValue
		fs = []string{"*"}
	)

	db := ctx.DB().Model(model.BusinessValue{}).Select(fs).Where("business_id = ?", bid)

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToBusinessValueEntity(m))
	}
	return bs, uint32(len(bs)), nil
}

// UpdateBusinessValues 更新数据
func (r businessRepo) UpdateBusinessValues(ctx kratosx.Context, bs []*biz.BusinessValue) error {
	var (
		ms []*model.BusinessValue
	)
	for _, b := range bs {
		ms = append(ms, r.ToBusinessValueModel(b))
	}
	return ctx.DB().Clauses(clause.OnConflict{UpdateAll: true}).Create(&ms).Error
}

// AllBusinessValue 获取指定环境的变量的所有值
func (s *businessRepo) AllBusinessValue(ctx kratosx.Context, eid, sid uint32) ([]*biz.BusinessValue, error) {
	// 获取指定服务的全部业务字段
	var ids []uint32
	if err := ctx.DB().Select("id").Model(biz.Business{}).Where("server_id=?", sid).Scan(&ids).Error; err != nil {
		return nil, err
	}

	var (
		ms []*model.BusinessValue
		bs []*biz.BusinessValue
	)
	if err := ctx.DB().
		Model(model.BusinessValue{}).
		Preload("Business").
		Where("env_id=? and business_id in ?", eid, ids).
		Find(&ms).Error; err != nil {
		return nil, err
	}
	for _, item := range ms {
		bs = append(bs, s.ToBusinessValueEntity(item))
	}
	return bs, nil
}

// AllBusinessField 获取可用的业务字段key列表
func (s *businessRepo) AllBusinessField(ctx kratosx.Context, srvId uint32) ([]string, error) {
	var list []string
	return list, ctx.DB().Select("keyword").
		Model(model.Business{}).
		Where("server_id=?", srvId).
		Find(&list).Error
}
