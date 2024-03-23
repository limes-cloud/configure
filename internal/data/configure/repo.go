package configure

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"

	biz "github.com/limes-cloud/configure/internal/biz/configure"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

// AddConfigure 新建配置
func (s *repo) AddConfigure(ctx kratosx.Context, configure *biz.Configure) (uint32, error) {
	return configure.ID, ctx.DB().Create(configure).Error
}

// GetConfigure 通过ID查找指定配置
func (s *repo) GetConfigure(ctx kratosx.Context, id uint32) (*biz.Configure, error) {
	var configure biz.Configure
	return &configure, ctx.DB().First(&configure, "id=?", id).Error
}

// GetConfigureByEnvAndSrv 查找指定配置
func (s *repo) GetConfigureByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*biz.Configure, error) {
	var configure biz.Configure
	return &configure, ctx.DB().First(&configure, "env_id=? and server_id=?", envId, srvId).Error
}

// PageConfigure 查询分页配置
func (s *repo) PageConfigure(ctx kratosx.Context, options *types.PageOptions) ([]*biz.Configure, uint32, error) {
	var list []*biz.Configure
	var total int64

	db := ctx.DB().Model(biz.Configure{})
	if options.Scopes != nil {
		db = db.Scopes(options.Scopes)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((options.Page - 1) * options.PageSize)).Limit(int(options.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// UpdateConfigure 更新指定id的配置
func (s *repo) UpdateConfigure(ctx kratosx.Context, configure *biz.Configure) error {
	return ctx.DB().Where("id=?", configure.ID).Updates(configure).Error
}

// DeleteConfigure 删除指定id的配置
func (s *repo) DeleteConfigure(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Configure{}, "id = ?", id).Error
}
