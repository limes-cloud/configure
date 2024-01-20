package business

import (
	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/internal/initiator/env"
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context) {
	db := ctx.DB().Begin()
	for _, bv := range businessValues {
		business := *bv.Business
		if err := db.Model(biz.Business{}).Create(&business).Error; err != nil {
			db.Rollback()
			pt.Error(err.Error())
			return
		}

		for _, envId := range env.Envs {
			businessValue := &biz.BusinessValue{
				EnvID:      envId,
				BusinessID: business.ID,
				Value:      bv.Value,
			}

			if err := db.Model(biz.BusinessValue{}).Create(businessValue).Error; err != nil {
				db.Rollback()
				pt.Error("初始化业务变量失败" + err.Error())
				return
			}
		}
	}
	db.Commit()
}
