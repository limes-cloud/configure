package server

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context, config *config.Config) {
	if err := ctx.DB().Model(model.Server{}).Create(&servers).Error; err != nil {
		pt.Error("init server error :" + err.Error())
	}
}
