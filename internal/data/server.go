package data

import (
	"fmt"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"google.golang.org/protobuf/proto"

	biz "github.com/limes-cloud/configure/internal/biz/server"
	"github.com/limes-cloud/configure/internal/data/model"
)

type serverRepo struct {
}

func NewServerRepo() biz.Repo {
	return &serverRepo{}
}

// ToServerEntity model转entity
func (r serverRepo) ToServerEntity(m *model.Server) *biz.Server {
	e := &biz.Server{}
	_ = valx.Transform(m, e)
	return e
}

// ToServerModel entity转model
func (r serverRepo) ToServerModel(e *biz.Server) *model.Server {
	m := &model.Server{}
	_ = valx.Transform(e, m)
	return m
}

// GetServerByKeyword 获取指定数据
func (r serverRepo) GetServerByKeyword(ctx kratosx.Context, keyword string) (*biz.Server, error) {
	var (
		m  = model.Server{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("keyword = ?", keyword).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToServerEntity(&m), nil
}

// GetServer 获取指定的数据
func (r serverRepo) GetServer(ctx kratosx.Context, id uint32) (*biz.Server, error) {
	var (
		m  = model.Server{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return r.ToServerEntity(&m), nil
}

// ListServer 获取列表
func (r serverRepo) ListServer(ctx kratosx.Context, req *biz.ListServerRequest) ([]*biz.Server, uint32, error) {
	var (
		bs    []*biz.Server
		ms    []*model.Server
		total int64
		fs    = []string{"id", "keyword", "name", "description", "status", "created_at", "updated_at"}
	)

	db := ctx.DB().Model(model.Server{}).Select(fs)

	if req.Keyword != nil {
		db = db.Where("keyword = ?", *req.Keyword)
	}
	if req.Name != nil {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Ids != nil {
		db = db.Where("id in ?", req.Ids)
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
		bs = append(bs, r.ToServerEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateServer 创建数据
func (r serverRepo) CreateServer(ctx kratosx.Context, req *biz.Server) (uint32, error) {
	m := r.ToServerModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateServer 更新数据
func (r serverRepo) UpdateServer(ctx kratosx.Context, req *biz.Server) error {
	return ctx.DB().Updates(r.ToServerModel(req)).Error
}

// DeleteServer 删除数据
func (r serverRepo) DeleteServer(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Where("id=?", id).Delete(&model.Server{}).Error
}

// UpdateServerStatus 更新数据状态
func (r serverRepo) UpdateServerStatus(ctx kratosx.Context, id uint32, status bool) error {
	return ctx.DB().Model(model.Server{}).Where("id=?", id).Update("status", status).Error
}
