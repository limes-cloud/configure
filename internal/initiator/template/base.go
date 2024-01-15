package template

import (
	"github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/configure/internal/initiator/server"
	"github.com/limes-cloud/configure/internal/initiator/template/yaml"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

var templates = []*model.Template{
	// Getaway 配置模板
	{
		BaseModel:   types.BaseModel{ID: server.Gateway},
		ServerID:    server.Gateway,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.Gateway,
		Version:     util.MD5ToUpper([]byte(yaml.Gateway))[:12],
		Description: "初始化模板",
	},
	// Manager业务配置
	{
		BaseModel:   types.BaseModel{ID: server.Manager},
		ServerID:    server.Manager,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.Manager,
		Version:     util.MD5ToUpper([]byte(yaml.Manager))[:12],
		Description: "初始化模板",
	},
	{
		BaseModel:   types.BaseModel{ID: server.Resource},
		ServerID:    server.Resource,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.Resource,
		Version:     util.MD5ToUpper([]byte(yaml.Resource))[:12],
		Description: "初始化模板",
	},
	{
		BaseModel:   types.BaseModel{ID: server.UserCenter},
		ServerID:    server.UserCenter,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.UserCenter,
		Version:     util.MD5ToUpper([]byte(yaml.UserCenter))[:12],
		Description: "初始化模板",
	},
	{
		BaseModel:   types.BaseModel{ID: server.PartyAffairs},
		ServerID:    server.PartyAffairs,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.PartyAffairs,
		Version:     util.MD5ToUpper([]byte(yaml.PartyAffairs))[:12],
		Description: "初始化模板",
	},
}
