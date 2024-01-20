package env

import (
	"github.com/limes-cloud/configure/internal/initiator/env"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*biz.Resource{
	{
		Keyword:     "Env",
		Description: "环境标识信息",
		Fields:      "Keyword",
		Private:     proto.Bool(false),
		Tag:         "env",
		ResourceValue: []*biz.ResourceValue{
			{
				EnvID: env.TEST,
				Value: util.MarshalString(map[string]any{
					"Keyword": "TEST",
				}),
			},
			{
				EnvID: env.PRE,
				Value: util.MarshalString(map[string]any{
					"Keyword": "PRE",
				}),
			},
			{
				EnvID: env.PROD,
				Value: util.MarshalString(map[string]any{
					"Keyword": "PROD",
				}),
			},
		},
	},
}
