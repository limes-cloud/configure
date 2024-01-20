package manager

import (
	"github.com/limes-cloud/configure/internal/initiator/env"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/internal/initiator/server"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*biz.Resource{
	{
		Keyword:     "ManagerServer",
		Description: "管理中心服务配置信息",
		Fields:      "Host,HttpPort,GrpcPort,Timeout",
		Private:     proto.Bool(false),
		Tag:         "server",
		ResourceValue: []*biz.ResourceValue{
			{
				EnvID: env.TEST,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 7001,
					"GrpcPort": 8001,
					"Timeout":  "10s",
				}),
			},
			{
				EnvID: env.PRE,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 7001,
					"GrpcPort": 8001,
					"Timeout":  "10s",
				}),
			},
			{
				EnvID: env.PROD,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 7001,
					"GrpcPort": 8001,
					"Timeout":  "10s",
				}),
			},
		},
	},
	{
		Keyword:     "ManagerDatabase",
		Description: "管理中心数据库配置",
		Fields:      "Username,Password,Type,Port,Database,Option,Host",
		Private:     proto.Bool(true),
		Tag:         "mysql",
		ResourceServer: []*biz.ResourceServer{
			{
				ServerID: server.Manager,
			},
		},
		ResourceValue: []*biz.ResourceValue{
			{
				EnvID: env.TEST,
				Value: util.MarshalString(map[string]any{
					"Username": "root",
					"Password": "root",
					"Host":     "127.0.0.1",
					"Port":     "3306",
					"Type":     "mysql",
					"Database": "manager",
					"Option":   "?charset=utf8mb4&parseTime=True&loc=Local",
				}),
			},
			{
				EnvID: env.PRE,
				Value: util.MarshalString(map[string]any{
					"Username": "root",
					"Password": "root",
					"Host":     "127.0.0.1",
					"Port":     "3306",
					"Type":     "mysql",
					"Database": "manager",
					"Option":   "?charset=utf8mb4&parseTime=True&loc=Local",
				}),
			},
			{
				EnvID: env.PROD,
				Value: util.MarshalString(map[string]any{
					"Username": "manager",
					"Password": "maGh8TzkjfyLbkYA",
					"Host":     "127.0.0.1",
					"Port":     "3306",
					"Type":     "mysql",
					"Database": "manager",
					"Option":   "?charset=utf8mb4&parseTime=True&loc=Local",
				}),
			},
		},
	},
}
