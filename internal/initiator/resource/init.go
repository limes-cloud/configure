package resource

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/initiator/resource/configure"
	"github.com/limes-cloud/configure/internal/initiator/resource/email"
	"github.com/limes-cloud/configure/internal/initiator/resource/env"
	"github.com/limes-cloud/configure/internal/initiator/resource/gateway"
	"github.com/limes-cloud/configure/internal/initiator/resource/jwt"
	"github.com/limes-cloud/configure/internal/initiator/resource/manager"
	"github.com/limes-cloud/configure/internal/initiator/resource/partyaffairs"
	"github.com/limes-cloud/configure/internal/initiator/resource/redis"
	"github.com/limes-cloud/configure/internal/initiator/resource/resource"
	"github.com/limes-cloud/configure/internal/initiator/resource/usercenter"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context, config *config.Config) {
	var resourceList []*model.Resource

	// env
	resourceList = append(resourceList, env.Resources...)
	// jwt
	resourceList = append(resourceList, jwt.Resources...)
	// redis
	resourceList = append(resourceList, redis.Resources...)
	// email
	resourceList = append(resourceList, email.Resources...)
	// gateway
	resourceList = append(resourceList, gateway.Resources...)
	// manager
	resourceList = append(resourceList, manager.Resources...)
	// resource
	resourceList = append(resourceList, resource.Resources...)
	// configure
	resourceList = append(resourceList, configure.Resources...)
	// user-center
	resourceList = append(resourceList, usercenter.UserCenter...)
	// party-affairs
	resourceList = append(resourceList, partyaffairs.PartyAffairs...)
	db := ctx.DB().Begin()
	// 进行便利创建
	for _, item := range resourceList {
		// 创建资源变量
		rs := *item
		if err := db.Create(&rs).Error; err != nil {
			pt.Error("创建资源变量失败：" + err.Error())
			return
		}

		// 创建资源变量值
		rsv := item.ResourceValue
		for ind := range rsv {
			rsv[ind].ResourceID = rs.ID
		}
		if err := db.Create(&rsv).Error; err != nil {
			pt.Error("创建资源变量值失败：" + err.Error())
			return
		}

		// 创建私有服务
		if len(item.ResourceServer) != 0 {
			srs := item.ResourceServer
			for ind := range srs {
				srs[ind].ResourceID = rs.ID
			}
			if err := db.Create(&srs).Error; err != nil {
				pt.Error("创建资源服务值失败：" + err.Error())
				return
			}
		}
	}

	db.Commit()
}
