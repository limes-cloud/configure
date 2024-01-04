package template

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context, config *config.Config) {
	if err := ctx.DB().Model(model.Template{}).Create(&templates).Error; err != nil {
		pt.Error("init template error" + err.Error())
	}
}
