package jwt

import (
	"github.com/limes-cloud/configure/internal/initiator/env"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*biz.Resource{
	{
		Keyword:     "AdminJwt",
		Description: "后台管理服务jwt配置信息",
		Fields:      "Secret,Expire,Renewal,Whitelist",
		Private:     proto.Bool(false),
		Tag:         "jwt",
		ResourceValue: []*biz.ResourceValue{
			{
				EnvID: env.TEST,
				Value: util.MarshalString(map[string]any{
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
				EnvID: env.PRE,
				Value: util.MarshalString(map[string]any{
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
				EnvID: env.PROD,
				Value: util.MarshalString(map[string]any{
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
