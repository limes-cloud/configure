package redis

import (
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/initiator/environment"
	"github.com/limes-cloud/configure/internal/model"
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

var Resources = []*model.Resource{
	{
		Keyword:     "Redis",
		Description: "redis中间件配置信息",
		Fields:      "Host,Username,Password,Port",
		Private:     proto.Bool(false),
		Tag:         "redis",
		ResourceValue: []*model.ResourceValue{
			{
				EnvironmentID: environment.TEST,
				Values: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"Username": "",
					"Password": "",
					"Port":     6379,
				}),
			},
			{
				EnvironmentID: environment.PRE,
				Values: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"Username": "",
					"Password": "",
					"Port":     6379,
				}),
			},
			{
				EnvironmentID: environment.PROD,
				Values: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"Username": "",
					"Password": "",
					"Port":     6379,
				}),
			},
		},
	},
}
