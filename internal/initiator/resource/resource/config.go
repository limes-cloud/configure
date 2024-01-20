package resource

import (
	"github.com/limes-cloud/configure/internal/initiator/env"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/internal/initiator/server"
	"github.com/limes-cloud/configure/pkg/util"
)

var Resources = []*biz.Resource{
	{
		Keyword:     "ResourceServer",
		Description: "资源中心服务配置信息",
		Fields:      "Host,HttpPort,GrpcPort,Timeout",
		Private:     proto.Bool(false),
		Tag:         "server",
		ResourceValue: []*biz.ResourceValue{
			{
				EnvID: env.TEST,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 7003,
					"GrpcPort": 8003,
					"Timeout":  "10s",
				}),
			},
			{
				EnvID: env.PRE,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 7003,
					"GrpcPort": 8003,
					"Timeout":  "10s",
				}),
			},
			{
				EnvID: env.PROD,
				Value: util.MarshalString(map[string]any{
					"Host":     "127.0.0.1",
					"HttpPort": 7003,
					"GrpcPort": 8003,
					"Timeout":  "10s",
				}),
			},
		},
	},
	{
		Keyword:     "ResourceDatabase",
		Description: "资源中心数据库配置信息",
		Fields:      "Username,Password,Type,Port,Database,Option,Host",
		Private:     proto.Bool(true),
		Tag:         "mysql",
		ResourceServer: []*biz.ResourceServer{
			{
				ServerID: server.Resource,
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
					"Database": "resource",
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
					"Database": "resource",
					"Option":   "?charset=utf8mb4&parseTime=True&loc=Local",
				}),
			},
			{
				EnvID: env.PROD,
				Value: util.MarshalString(map[string]any{
					"Username": "resource",
					"Password": "Ti7MaKJJznywNBJb",
					"Host":     "127.0.0.1",
					"Port":     "3306",
					"Type":     "mysql",
					"Database": "resource",
					"Option":   "?charset=utf8mb4&parseTime=True&loc=Local",
				}),
			},
		},
	},
}
