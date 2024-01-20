package template

import (
	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context) {
	if err := ctx.DB().Model(biz.Template{}).Create(&templates).Error; err != nil {
		pt.Error("init template error" + err.Error())
	}
}
