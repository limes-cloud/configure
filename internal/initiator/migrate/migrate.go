package migrate

import (
	"github.com/limes-cloud/kratosx"
	gte "github.com/limes-cloud/kratosx/library/db/gormtranserror"

	"github.com/limes-cloud/configure/internal/biz"
)

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

func Init(ctx kratosx.Context) {
	db := ctx.DB()
	_ = db.Set("gorm:table_options", "COMMENT='环境信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Env{})
	_ = db.Set("gorm:table_options", "COMMENT='服务信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Server{})
	_ = db.Set("gorm:table_options", "COMMENT='资源变量' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Resource{})
	_ = db.Set("gorm:table_options", "COMMENT='资源变量值' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.ResourceValue{})
	_ = db.Set("gorm:table_options", "COMMENT='资源服务信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.ResourceServer{})
	_ = db.Set("gorm:table_options", "COMMENT='业务变量' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Business{})
	_ = db.Set("gorm:table_options", "COMMENT='业务变量值' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.BusinessValue{})
	_ = db.Set("gorm:table_options", "COMMENT='配置模板' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Template{})
	_ = db.Set("gorm:table_options", "COMMENT='配置内容' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(biz.Configure{})

	// 重新载入gorm错误插件
	_ = gte.NewGlobalGormErrorPlugin().Initialize(ctx.DB())
}
