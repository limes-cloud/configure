package jwt

import (
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/initiator/environment"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*model.Resource{
	{
		Keyword:     "AdminJwt",
		Description: "后台管理服务jwt配置信息",
		Fields:      "Secret,Expire,Renewal,Whitelist",
		Private:     proto.Bool(false),
		Tag:         "jwt",
		ResourceValue: []*model.ResourceValue{
			{
				EnvironmentID: environment.TEST,
				Values: util.MarshalString(map[string]any{
					"Secret":  "limes-cloud-test",
					"Expire":  "2h",
					"Renewal": "120s",
					"Whitelist": map[string]bool{
						"GET:/manager/v1/setting":        true,
						"POST:/manager/v1/login/captcha": true,
						"POST:/manager/v1/login":         true,
						"POST:/manager/v1/token/refresh": true,
					},
				}),
			},
			{
				EnvironmentID: environment.PRE,
				Values: util.MarshalString(map[string]any{
					"Secret":  "limes-cloud-pre",
					"Expire":  "2h",
					"Renewal": "120s",
					"Whitelist": map[string]bool{
						"GET:/manager/v1/setting":        true,
						"POST:/manager/v1/login/captcha": true,
						"POST:/manager/v1/login":         true,
						"POST:/manager/v1/token/refresh": true,
					},
				}),
			},
			{
				EnvironmentID: environment.PROD,
				Values: util.MarshalString(map[string]any{
					"Secret":  "limes-cloud-prod",
					"Expire":  "2h",
					"Renewal": "120s",
					"Whitelist": map[string]bool{
						"GET:/manager/v1/setting":        true,
						"POST:/manager/v1/login/captcha": true,
						"POST:/manager/v1/login":         true,
						"POST:/manager/v1/token/refresh": true,
					},
				}),
			},
		},
	},
}
