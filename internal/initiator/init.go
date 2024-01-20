package initiator

import (
	"context"

	"github.com/limes-cloud/configure/internal/service"

	"github.com/limes-cloud/configure/internal/initiator/env"

	"github.com/limes-cloud/configure/internal/biz"

	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/initiator/business"
	configure "github.com/limes-cloud/configure/internal/initiator/configrue"
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

func IsInit(ctx kratosx.Context) bool {
	db := ctx.DB().Migrator()
	return db.HasTable(biz.Env{}) &&
		db.HasTable(biz.Server{}) &&
		db.HasTable(biz.Resource{}) &&
		db.HasTable(biz.ResourceValue{}) &&
		db.HasTable(biz.ResourceServer{}) &&
		db.HasTable(biz.Business{}) &&
		db.HasTable(biz.BusinessValue{}) &&
		db.HasTable(biz.Template{}) &&
		db.HasTable(biz.Configure{})
}

// Run 执行系统初始化
func (a *Initiator) Run(srv *service.Service) error {
	ctx := kratosx.MustContext(context.Background())

	if migrate.IsInit(ctx) {
		pt.Cyan("already init server")
		return nil
	}

	// 自动迁移
	migrate.Init(ctx)

	// 环境初始化
	env.Init(ctx)

	// 服务初始化
	server.Init(ctx)

	// 资源初始化
	resource.Init(ctx)

	// 业务服务初始化
	business.Init(ctx)

	// 配置模板初始化
	template.Init(ctx)

	// 配置初始化
	configure.Init(ctx, srv)
	return nil
}
