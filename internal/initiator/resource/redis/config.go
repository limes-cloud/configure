package redis

import (
	"github.com/limes-cloud/configure/internal/initiator/env"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/pkg/util"
)

var TEST = map[string]any{
	"Host":     "127.0.0.1",
	"Username": "",
	"Password": "",
	"Port":     6379,
}

var PRE = map[string]any{
	"Host":     "127.0.0.1",
	"Username": "",
	"Password": "",
	"Port":     6379,
}

var PROD = map[string]any{
	"Host":     "127.0.0.1",
	"Username": "",
	"Password": "",
	"Port":     6379,
}

var Resources = []*biz.Resource{
	{
		Keyword:     "Redis",
		Description: "redis中间件配置信息",
		Fields:      "Host,Username,Password,Port",
		Private:     proto.Bool(false),
		Tag:         "redis",
		ResourceValue: []*biz.ResourceValue{
			{
				EnvID: env.TEST,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"Username": "",
					"Password": "",
					"Port":     6379,
				}),
			},
			{
				EnvID: env.PRE,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"Username": "",
					"Password": "",
					"Port":     6379,
				}),
			},
			{
				EnvID: env.PROD,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"Username": "",
					"Password": "",
					"Port":     6379,
				}),
			},
		},
	},
}
