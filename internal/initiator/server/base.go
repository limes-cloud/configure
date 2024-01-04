package server

import (
	"github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/configure/internal/model"
)

const (
	Gateway  = 1
	Manager  = 2
	Resource = 3
)

var Servers = []uint32{Gateway, Manager, Resource}

var servers = []*model.Server{
	{
		BaseModel:   types.BaseModel{ID: Gateway},
		Keyword:     "Gateway",
		Name:        "通用网关",
		Description: "主要负责前端到后端的转发",
	},
	{
		BaseModel:   types.BaseModel{ID: Manager},
		Keyword:     "Manager",
		Name:        "管理中心",
		Description: "主要负责系统的基础管理",
	},
	{
		BaseModel:   types.BaseModel{ID: Resource},
		Keyword:     "Resource",
		Name:        "资源中心",
		Description: "主要负责静态资源的管理",
	},
}
