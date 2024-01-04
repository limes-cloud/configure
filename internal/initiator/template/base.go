package template

import (
	"github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/configure/internal/initiator/server"
	"github.com/limes-cloud/configure/internal/initiator/template/yaml"
	"github.com/limes-cloud/configure/internal/model"
)

var templates = []*model.Template{
	// Getaway 配置模板
	{
		BaseModel:   types.BaseModel{ID: 1},
		ServerID:    server.Gateway,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.Gateway,
		Description: "初始化模板",
	},
	// Manager业务配置
	{
		BaseModel:   types.BaseModel{ID: 2},
		ServerID:    server.Manager,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.Manager,
		Description: "初始化模板",
	},

	{
		BaseModel:   types.BaseModel{ID: 3},
		ServerID:    server.Resource,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.Resource,
		Description: "初始化模板",
	},
}
