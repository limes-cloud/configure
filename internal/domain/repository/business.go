package repository

import (
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type Business interface {
	// GetBusiness 获取指定的业务配置信息
	GetBusiness(ctx core.Context, id uint32) (*entity.Business, error)

	// ListBusiness 获取业务配置信息列表
	ListBusiness(ctx core.Context, req *types.ListBusinessRequest) ([]*entity.Business, uint32, error)

	// CreateBusiness 创建业务配置信息
	CreateBusiness(ctx core.Context, req *entity.Business) (uint32, error)

	// UpdateBusiness 更新业务配置信息
	UpdateBusiness(ctx core.Context, req *entity.Business) error

	// DeleteBusiness 删除业务配置信息
	DeleteBusiness(ctx core.Context, id uint32) error

	// ListBusinessValue 获取业务配置值信息列表
	ListBusinessValue(ctx core.Context, req *types.ListBusinessValueRequest) ([]*entity.BusinessValue, error)

	// UpdateBusinessValues 更新业务配置值信息
	UpdateBusinessValues(ctx core.Context, bs []*entity.BusinessValue) error

	// AllBusinessField 获取全部可用的字段
	AllBusinessField(ctx core.Context, app string) ([]string, error)
}
