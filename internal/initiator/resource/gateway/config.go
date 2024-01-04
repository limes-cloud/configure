package gateway

import (
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/initiator/environment"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*model.Resource{
	{
		Keyword:     "GatewayServer",
		Description: "通用网关服务配置信息",
		Fields:      "HttpPort,Host,Timeout",
		Private:     proto.Bool(false),
		Tag:         "server",
		ResourceValue: []*model.ResourceValue{
			{
				EnvironmentID: environment.TEST,
				Values: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 7080,
					"Timeout":  "10s",
				}),
			},
			{
				EnvironmentID: environment.PRE,
				Values: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 7080,
					"Timeout":  "10s",
				}),
			},
			{
				EnvironmentID: environment.PROD,
				Values: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 7080,
					"Timeout":  "10s",
				}),
			},
		},
	},
}
