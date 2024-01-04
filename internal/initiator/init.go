package initiator

import (
	"context"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/initiator/business"
	configure "github.com/limes-cloud/configure/internal/initiator/configrue"
	"github.com/limes-cloud/configure/internal/initiator/environment"
	"github.com/limes-cloud/configure/internal/initiator/migrate"
	"github.com/limes-cloud/configure/internal/initiator/resource"
	"github.com/limes-cloud/configure/internal/initiator/server"
	"github.com/limes-cloud/configure/internal/initiator/template"
	"github.com/limes-cloud/configure/pkg/pt"
)

type Initiator struct {
	conf *config.Config
}

func New(conf *config.Config) *Initiator {
	return &Initiator{
		conf: conf,
	}
}

// Run 执行系统初始化
func (a *Initiator) Run() error {
	ctx := kratosx.MustContext(context.Background())

	if migrate.IsInit(ctx) {
		pt.Cyan("already init server")
		return nil
	}

	// 自动迁移
	migrate.Init(ctx, a.conf)

	// 环境初始化
	environment.Init(ctx, a.conf)

	// 服务初始化
	server.Init(ctx, a.conf)

	// 资源初始化
	resource.Init(ctx, a.conf)

	// 业务服务初始化
	business.Init(ctx, a.conf)

	// 配置模板初始化
	template.Init(ctx, a.conf)

	// 配置初始化
	configure.Init(ctx, a.conf)
	return nil
}
