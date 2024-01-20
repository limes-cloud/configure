package env

import (
	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context) {
	if err := ctx.DB().Model(biz.Env{}).Create(&envs).Error; err != nil {
		pt.Error("init env error :" + err.Error())
	}
}
