package repository

import (
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type Resource interface {
	// GetResource 获取指定的资源配置信息
	GetResource(ctx core.Context, id uint32) (*entity.Resource, error)

	// ListResource 获取资源配置信息列表
	ListResource(ctx core.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error)

	// CreateResource 创建资源配置信息
	CreateResource(ctx core.Context, resource *entity.Resource) (uint32, error)

	// UpdateResource 更新资源配置信息
	UpdateResource(ctx core.Context, resource *entity.Resource) error

	// DeleteResource 删除资源配置信息
	DeleteResource(ctx core.Context, id uint32) error

	// GetResourceByKeyword 获取指定的资源配置信息
	GetResourceByKeyword(ctx core.Context, keyword string) (*entity.Resource, error)

	// ListResourceValue 获取业务配置值信息列表
	ListResourceValue(ctx core.Context, req *types.ListResourceValueRequest) ([]*entity.ResourceValue, error)

	// UpdateResourceValues 更新业务配置值信息
	UpdateResourceValues(ctx core.Context, rvs []*entity.ResourceValue) error

	// AllResourceField 获取全部可用的字段
	AllResourceField(ctx core.Context, app string) ([]string, error)
}
