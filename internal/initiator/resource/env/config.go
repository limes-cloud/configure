package env

import (
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/initiator/environment"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*model.Resource{
	{
		Keyword:     "Env",
		Description: "环境标识信息",
		Fields:      "Keyword",
		Private:     proto.Bool(false),
		Tag:         "env",
		ResourceValue: []*model.ResourceValue{
			{
				EnvironmentID: environment.TEST,
				Values: util.MarshalString(map[string]any{
					"Keyword": "TEST",
				}),
			},
			{
				EnvironmentID: environment.PRE,
				Values: util.MarshalString(map[string]any{
					"Keyword": "PRE",
				}),
			},
			{
				EnvironmentID: environment.PROD,
				Values: util.MarshalString(map[string]any{
					"Keyword": "PROD",
				}),
			},
		},
	},
}
