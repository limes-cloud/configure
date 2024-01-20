package data

import (
	"github.com/limes-cloud/configure/internal/biz/types"
	"gorm.io/gorm"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type serverRepo struct {
}

func NewServerRepo() biz.ServerRepo {
	return &serverRepo{}
}

// Create 新建服务
func (s *serverRepo) Create(ctx kratosx.Context, server *biz.Server) (uint32, error) {
	return server.ID, ctx.DB().Create(server).Error
}

// Get 通过ID查找指定服务
func (s *serverRepo) Get(ctx kratosx.Context, id uint32) (*biz.Server, error) {
	var server biz.Server
	return &server, ctx.DB().First(&server, "id=?", id).Error
}

// GetByKeyword 通过关键词查找指定服务
func (s *serverRepo) GetByKeyword(ctx kratosx.Context, keyword string) (*biz.Server, error) {
	var server biz.Server
	return &server, ctx.DB().First(&server, "keyword=?", keyword).Error
}

// Page 查询分页服务
func (s *serverRepo) Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*biz.Server, uint32, error) {
	var list []*biz.Server
	var total int64

	db := ctx.DB().Model(biz.Server{})
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}
func (s *serverRepo) PageServer(ctx kratosx.Context, req *types.PageServerRequest) ([]*biz.Server, uint32, error) {
	return s.Page(ctx, &ktypes.PageOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if req.Keyword != nil {
				db = db.Where("keyword like ?", *req.Keyword+"%")
			}
			return db
		},
	})
}

// All 查询全部服务
func (s *serverRepo) All(ctx kratosx.Context, scopes ktypes.Scopes) ([]*biz.Server, error) {
	var list []*biz.Server

	db := ctx.DB().Model(biz.Server{})
	if scopes != nil {
		db = db.Scopes(scopes)
	}

	return list, db.Find(&list).Error
}

// Update 更新指定id的服务
func (s *serverRepo) Update(ctx kratosx.Context, server *biz.Server) error {
	return ctx.DB().Where("id=?", server.ID).Updates(server).Error
}

// Delete 删除指定id的服务
func (s *serverRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Server{}, "id = ?", id).Error
}
