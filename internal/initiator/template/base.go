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
		BaseModel:   types.BaseModel{ID: 1},
		ServerID:    server.Gateway,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.Gateway,
		Version:     util.MD5ToUpper([]byte(yaml.Gateway))[:12],
		Description: "初始化模板",
	},
	// Manager业务配置
	{
		BaseModel:   types.BaseModel{ID: 2},
		ServerID:    server.Manager,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.Manager,
		Version:     util.MD5ToUpper([]byte(yaml.Gateway))[:12],
		Description: "初始化模板",
	},

	{
		BaseModel:   types.BaseModel{ID: 3},
		ServerID:    server.Resource,
		IsUse:       true,
		Format:      "yaml",
		Content:     yaml.Resource,
		Version:     util.MD5ToUpper([]byte(yaml.Gateway))[:12],
		Description: "初始化模板",
	},
}
