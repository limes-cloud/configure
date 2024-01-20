package email

import (
	"github.com/limes-cloud/configure/internal/initiator/env"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*biz.Resource{
	{
		Keyword:     "Email",
		Description: "邮箱服务配置信息",
		Fields:      "Username,Company,Password,Host,Port",
		Private:     proto.Bool(false),
		Tag:         "email",
		ResourceValue: []*biz.ResourceValue{
			{
				EnvID: env.TEST,
				Value: util.MarshalString(map[string]any{
					"Username": "860808187@qq.com",
					"Company":  "青岑云",
					"Password": "xxx",
					"Host":     "smtp.qq.com",
					"Port":     25,
				}),
			},
			{
				EnvID: env.PRE,
				Value: util.MarshalString(map[string]any{
					"Username": "860808187@qq.com",
					"Company":  "青岑云",
					"Password": "xxx",
					"Host":     "smtp.qq.com",
					"Port":     25,
				}),
			},
			{
				EnvID: env.PROD,
				Value: util.MarshalString(map[string]any{
					"Username": "860808187@qq.com",
					"Company":  "青岑云",
					"Password": "xxx",
					"Host":     "smtp.qq.com",
					"Port":     25,
				}),
			},
		},
	},
}
