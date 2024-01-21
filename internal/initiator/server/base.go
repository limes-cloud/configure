package server

import (
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/internal/biz"
)

const (
	Gateway      = 1
	Manager      = 2
	Resource     = 3
	UserCenter   = 4
	PartyAffairs = 5
)

var Servers = []uint32{Gateway, Manager, Resource, UserCenter, PartyAffairs}

var servers = []*biz.Server{
	{
		BaseModel:   types.BaseModel{ID: Gateway},
		Keyword:     "Gateway",
		Name:        "通用网关",
		IsBusiness:  proto.Bool(false),
		Description: "主要负责前端到后端的转发",
	},
	{
		BaseModel:   types.BaseModel{ID: Manager},
		Keyword:     "Manager",
		Name:        "管理中心",
		IsBusiness:  proto.Bool(false),
		Description: "主要负责系统的基础管理",
	},
	{
		BaseModel:   types.BaseModel{ID: Resource},
		Keyword:     "Resource",
		Name:        "资源中心",
		IsBusiness:  proto.Bool(false),
		Description: "主要负责静态资源的管理",
	},
	{
		BaseModel:   types.BaseModel{ID: UserCenter},
		Keyword:     "UserCenter",
		Name:        "用户中心",
		IsBusiness:  proto.Bool(false),
		Description: "主要负责业务用户的管理",
	},
	{
		BaseModel:   types.BaseModel{ID: PartyAffairs},
		Keyword:     "PartyAffairs",
		Name:        "信号灯系统",
		IsBusiness:  proto.Bool(true),
		Description: "指尖上的党务系统",
	},
}
