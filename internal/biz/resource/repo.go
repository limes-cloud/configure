package resource

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// GetResource 获取指定的资源配置信息
	GetResource(ctx kratosx.Context, id uint32) (*Resource, error)

	// ListResource 获取资源配置信息列表
	ListResource(ctx kratosx.Context, req *ListResourceRequest) ([]*Resource, uint32, error)

	// CreateResource 创建资源配置信息
	CreateResource(ctx kratosx.Context, req *Resource) (uint32, error)

	// UpdateResource 更新资源配置信息
	UpdateResource(ctx kratosx.Context, req *Resource) error

	// DeleteResource 删除资源配置信息
	DeleteResource(ctx kratosx.Context, ids []uint32) (uint32, error)

	// GetResourceByKeyword 获取指定的资源配置信息
	GetResourceByKeyword(ctx kratosx.Context, keyword string) (*Resource, error)

	// ListResourceValue 获取业务配置值信息列表
	ListResourceValue(ctx kratosx.Context, bid uint32) ([]*ResourceValue, uint32, error)

	// UpdateResourceValues 更新业务配置值信息
	UpdateResourceValues(ctx kratosx.Context, bs []*ResourceValue) error

	// AllResourceValue 获取全部可用的值
	AllResourceValue(ctx kratosx.Context, eid, sid uint32) ([]*ResourceValue, error)

	// AllResourceField 获取全部可用的字段
	AllResourceField(ctx kratosx.Context, sid uint32) ([]string, error)
}
