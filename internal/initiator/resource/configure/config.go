package configure

import (
	"github.com/limes-cloud/configure/internal/initiator/env"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*biz.Resource{
	{
		Keyword:     "ConfigureServer",
		Description: "配置中心服务配置信息",
		Fields:      "Host,HttpPort,GrpcPort,Timeout",
		Private:     proto.Bool(false),
		Tag:         "server",
		ResourceValue: []*biz.ResourceValue{
			{
				EnvID: env.TEST,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 6081,
					"GrpcPort": 6082,
					"Timeout":  "10s",
				}),
			},
			{
				EnvID: env.PRE,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 6081,
					"GrpcPort": 6082,
					"Timeout":  "10s",
				}),
			},
			{
				EnvID: env.PROD,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 6081,
					"GrpcPort": 6082,
					"Timeout":  "10s",
				}),
			},
		},
	},
}
