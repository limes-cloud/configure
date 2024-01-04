package configure

import (
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/initiator/environment"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*model.Resource{
	{
		Keyword:     "ConfigureServer",
		Description: "配置中心服务配置信息",
		Fields:      "Host,HttpPort,GrpcPort,Timeout",
		Private:     proto.Bool(false),
		Tag:         "server",
		ResourceValue: []*model.ResourceValue{
			{
				EnvironmentID: environment.TEST,
				Values: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 6081,
					"GrpcPort": 6082,
					"Timeout":  "10s",
				}),
			},
			{
				EnvironmentID: environment.PRE,
				Values: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 6081,
					"GrpcPort": 6082,
					"Timeout":  "10s",
				}),
			},
			{
				EnvironmentID: environment.PROD,
				Values: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 6081,
					"GrpcPort": 6082,
					"Timeout":  "10s",
				}),
			},
		},
	},
}
