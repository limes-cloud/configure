package business

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/initiator/environment"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context, config *config.Config) {
	db := ctx.DB().Begin()
	for _, bv := range businessValues {
		business := *bv.Business
		if err := db.Model(model.Business{}).Create(&business).Error; err != nil {
			db.Rollback()
			pt.Error(err.Error())
			return
		}

		for _, envId := range environment.Envs {
			businessValue := &model.BusinessValue{
				EnvironmentID: envId,
				BusinessID:    business.ID,
				Value:         bv.Value,
			}

			if err := db.Model(model.BusinessValue{}).Create(businessValue).Error; err != nil {
				db.Rollback()
				pt.Error("初始化业务变量失败" + err.Error())
				return
			}
		}
	}
	db.Commit()
}
