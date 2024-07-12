package business

import (
	"github.com/limes-cloud/kratosx"
)

type Repo interface {
	// GetBusiness 获取指定的业务配置信息
	GetBusiness(ctx kratosx.Context, id uint32) (*Business, error)

	// ListBusiness 获取业务配置信息列表
	ListBusiness(ctx kratosx.Context, req *ListBusinessRequest) ([]*Business, uint32, error)

	// CreateBusiness 创建业务配置信息
	CreateBusiness(ctx kratosx.Context, req *Business) (uint32, error)

	// UpdateBusiness 更新业务配置信息
	UpdateBusiness(ctx kratosx.Context, req *Business) error

	// DeleteBusiness 删除业务配置信息
	DeleteBusiness(ctx kratosx.Context, id uint32) error

	// ListBusinessValue 获取业务配置值信息列表
	ListBusinessValue(ctx kratosx.Context, bid uint32) ([]*BusinessValue, uint32, error)

	// UpdateBusinessValues 更新业务配置值信息
	UpdateBusinessValues(ctx kratosx.Context, bs []*BusinessValue) error

	// AllBusinessField 获取全部可用的字段
	AllBusinessField(ctx kratosx.Context, sid uint32) ([]string, error)

	// AllBusinessValue 获取全部可以用的值
	AllBusinessValue(ctx kratosx.Context, eid, sid uint32) ([]*BusinessValue, error)
}
