package server

import (
	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context) {
	if err := ctx.DB().Model(biz.Server{}).Create(&servers).Error; err != nil {
		pt.Error("init server error :" + err.Error())
	}
}
