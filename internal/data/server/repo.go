package server

import (
	"github.com/limes-cloud/kratosx"

	biz "github.com/limes-cloud/configure/internal/biz/server"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

// AddServer 新建服务
func (s *repo) AddServer(ctx kratosx.Context, server *biz.Server) (uint32, error) {
	return server.ID, ctx.DB().Create(server).Error
}

// GetServer 通过ID查找指定服务
func (s *repo) GetServer(ctx kratosx.Context, id uint32) (*biz.Server, error) {
	var server biz.Server
	return &server, ctx.DB().First(&server, "id=?", id).Error
}

// GetServerByKeyword 通过关键词查找指定服务
func (s *repo) GetServerByKeyword(ctx kratosx.Context, keyword string) (*biz.Server, error) {
	var server biz.Server
	return &server, ctx.DB().First(&server, "keyword=?", keyword).Error
}

// PageServer 查询分页服务
func (s *repo) PageServer(ctx kratosx.Context, in *biz.PageServerRequest) ([]*biz.Server, uint32, error) {
	var list []*biz.Server
	var total int64

	db := ctx.DB().Model(biz.Server{})
	if in.Keyword != nil {
		db = db.Where("keyword like ?", *in.Keyword+"%")
	}
	if in.IsBusiness != nil {
		db = db.Where("is_business=?", *in.IsBusiness)
	}
	if in.ServerScope != nil {
		db = db.Where("id in ?", in.ServerScope)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((in.Page - 1) * in.PageSize)).Limit(int(in.PageSize))
	return list, uint32(total), db.Find(&list).Error
}

// // All 查询全部服务
// func (s *repo) All(ctx kratosx.Context, scopes ktypes.Scopes) ([]*biz.Server, error) {
//	var list []*biz.Server
//
//	db := ctx.DB().Model(biz.Server{})
//	if scopes != nil {
//		db = db.Scopes(scopes)
//	}
//
//	return list, db.Find(&list).Error
// }

// UpdateServer 更新指定id的服务
func (s *repo) UpdateServer(ctx kratosx.Context, server *biz.Server) error {
	return ctx.DB().Where("id=?", server.ID).Updates(server).Error
}

// DeleteServer 删除指定id的服务
func (s *repo) DeleteServer(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Server{}, "id = ?", id).Error
}

// func (u *repo) GetServerByIds(ctx kratosx.Context, ids []uint32) ([]*biz.Server, error) {
//	var list []*biz.Server
//	return list, ctx.DB().Model(biz.Server{}).Where("id in ?", ids).Find(&list).Error
// }
