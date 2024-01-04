package email

import (
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/initiator/environment"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*model.Resource{
	{
		Keyword:     "Email",
		Description: "邮箱服务配置信息",
		Fields:      "Username,Company,Password,Host,Port",
		Private:     proto.Bool(false),
		Tag:         "email",
		ResourceValue: []*model.ResourceValue{
			{
				EnvironmentID: environment.TEST,
				Values: util.MarshalString(map[string]any{
					"Username": "860808187@qq.com",
					"Company":  "青岑云",
					"Password": "xxx",
					"Host":     "smtp.qq.com",
					"Port":     25,
				}),
			},
			{
				EnvironmentID: environment.PRE,
				Values: util.MarshalString(map[string]any{
					"Username": "860808187@qq.com",
					"Company":  "青岑云",
					"Password": "xxx",
					"Host":     "smtp.qq.com",
					"Port":     25,
				}),
			},
			{
				EnvironmentID: environment.PROD,
				Values: util.MarshalString(map[string]any{
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
