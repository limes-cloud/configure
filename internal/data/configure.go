package data

import (
	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/types"
)

type configureRepo struct {
}

func NewConfigureRepo() biz.ConfigureRepo {
	return &configureRepo{}
}

// Create 新建配置
func (s *configureRepo) Create(ctx kratosx.Context, configure *biz.Configure) (uint32, error) {
	return configure.ID, ctx.DB().Create(configure).Error
}

// Get 通过ID查找指定配置
func (s *configureRepo) Get(ctx kratosx.Context, id uint32) (*biz.Configure, error) {
	var configure biz.Configure
	return &configure, ctx.DB().First(&configure, "id=?", id).Error
}

// GetByEnvAndSrv 查找指定配置
func (s *configureRepo) GetByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*biz.Configure, error) {
	var configure biz.Configure
	return &configure, ctx.DB().First(&configure, "env_id=? and server_id=?", envId, srvId).Error
}

// Page 查询分页配置
func (s *configureRepo) Page(ctx kratosx.Context, options *types.PageOptions) ([]*biz.Configure, uint32, error) {
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

// Update 更新指定id的配置
func (s *configureRepo) Update(ctx kratosx.Context, configure *biz.Configure) error {
	return ctx.DB().Where("id=?", configure.ID).Updates(configure).Error
}

// Delete 删除指定id的配置
func (s *configureRepo) Delete(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Configure{}, "id = ?", id).Error
}
