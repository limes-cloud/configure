package repository

import (
	"github.com/limes-cloud/configure/internal/core"

	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type Configure interface {
	// GetConfigure 获取指定的配置
	GetConfigure(ctx core.Context, id uint32) (*entity.Configure, error)

	// GetConfigureByEnvAndApp 通过envId和app获取指定的配置
	GetConfigureByEnvAndApp(ctx core.Context, envId uint32, app string) (*entity.Configure, error)

	// ListConfigure 获取配置列表
	ListConfigure(ctx core.Context, req *types.ListConfigureRequest) ([]*entity.Configure, uint32, error)

	// CreateConfigure 创建配置
	CreateConfigure(ctx core.Context, c *entity.Configure) (uint32, error)

	// UpdateConfigure 更新配置
	UpdateConfigure(ctx core.Context, c *entity.Configure) error

	// BroadcastConfigure 广播配置变更
	BroadcastConfigure(ctx core.Context, envId uint32, app string) error

	// SubscribeConfigure 监听配置变更广播
	SubscribeConfigure(f func(ctx core.Context, envId uint32, app string) error)
}
