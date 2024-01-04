package migrate

import (
	"github.com/limes-cloud/kratosx"
	gte "github.com/limes-cloud/kratosx/library/db/gormtranserror"

	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
)

func IsInit(ctx kratosx.Context) bool {
	db := ctx.DB().Migrator()
	return db.HasTable(model.Environment{}) &&
		db.HasTable(model.Server{}) &&
		db.HasTable(model.Resource{}) &&
		db.HasTable(model.ResourceValue{}) &&
		db.HasTable(model.ResourceServer{}) &&
		db.HasTable(model.Business{}) &&
		db.HasTable(model.BusinessValue{}) &&
		db.HasTable(model.Template{}) &&
		db.HasTable(model.Configure{}) &&
		db.HasTable(model.OperateLog{})
}

func Init(ctx kratosx.Context, config *config.Config) {
	db := ctx.DB()
	_ = db.Set("gorm:table_options", "COMMENT='环境信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Environment{})
	_ = db.Set("gorm:table_options", "COMMENT='服务信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Server{})
	_ = db.Set("gorm:table_options", "COMMENT='资源变量' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Resource{})
	_ = db.Set("gorm:table_options", "COMMENT='资源变量值' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.ResourceValue{})
	_ = db.Set("gorm:table_options", "COMMENT='资源服务信息' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.ResourceServer{})
	_ = db.Set("gorm:table_options", "COMMENT='业务变量' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Business{})
	_ = db.Set("gorm:table_options", "COMMENT='业务变量值' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.BusinessValue{})
	_ = db.Set("gorm:table_options", "COMMENT='配置模板' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Template{})
	_ = db.Set("gorm:table_options", "COMMENT='配置内容' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.Configure{})
	_ = db.Set("gorm:table_options", "COMMENT='操作记录' ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(model.OperateLog{})

	// 重新载入gorm错误插件
	_ = gte.NewGlobalGormErrorPlugin().Initialize(ctx.DB())
}
