package menu

import (
	"github.com/limes-cloud/kratosx/pkg/value"
	managerMenu "github.com/limes-cloud/manager/api/menu"
)

const (
	appKeyword = "configure"

	menuTypeApi   = "A"
	menuTypePage  = "M"
	menuTypeGroup = "G"
)

var menus = []*managerMenu.RegisterMenuNode{
	{
		Key:      value.Pointer("ConfigureBaseApi"),
		Title:    "配置中心基础接口",
		Type:     menuTypeGroup,
		Icon:     value.Pointer("apps"),
		IsHidden: value.Pointer(true),
		Weight:   value.Pointer[int32](50),
		Children: []*managerMenu.RegisterMenuNode{
			api("获取配置", "/configure/api/configure", "GET", ""),
			api("同步配置", "/configure/api/configure", "PUT", "configure:configure:sync"),
			api("配置对比", "/configure/api/configure/compare", "POST", "configure:configure:compare"),
			api("监听配置变更", "/configure.api.configure.Configure/WatchConfigure", "POST", ""),
		},
	},
	{
		Key:       value.Pointer("ConfigureEnv"),
		Title:     "环境管理",
		Type:      menuTypePage,
		Icon:      value.Pointer("iconsys-yunduan"),
		Path:      value.Pointer("/configure/env"),
		Component: value.Pointer("/configure/env/index"),
		Weight:    value.Pointer[int32](40),
		Children: []*managerMenu.RegisterMenuNode{
			api("查询环境", "/configure/api/envs", "GET", "configure:env:query"),
			api("创建环境", "/configure/api/env", "POST", "configure:env:add"),
			api("更新环境", "/configure/api/env", "PUT", "configure:env:update"),
			api("删除环境", "/configure/api/env", "DELETE", "configure:env:delete"),
			api("获取环境Token", "/configure/api/env/token", "GET", "configure:env:token:query"),
			api("重置环境Token", "/configure/api/env/token", "PUT", "configure:env:token:reset"),
		},
	},
	{
		Key:       value.Pointer("ConfigureBusiness"),
		Title:     "业务变量",
		Type:      menuTypePage,
		Icon:      value.Pointer("iconsys-liebiaoxingshi"),
		Path:      value.Pointer("/configure/business"),
		Component: value.Pointer("/configure/business/index"),
		Weight:    value.Pointer[int32](20),
		Children: []*managerMenu.RegisterMenuNode{
			api("查询业务配置", "/configure/api/businesses", "GET", "configure:business:query"),
			api("创建业务配置", "/configure/api/business", "POST", "configure:business:add"),
			api("更新业务配置", "/configure/api/business", "PUT", "configure:business:update"),
			api("删除业务配置", "/configure/api/business", "DELETE", "configure:business:delete"),
			api("查询业务配置值", "/configure/api/business/values", "GET", "configure:business:value:query"),
			api("更新业务配置值", "/configure/api/business/values", "PUT", "configure:business:value:update"),
		},
	},
	{
		Key:       value.Pointer("ConfigureResource"),
		Title:     "资源变量",
		Type:      menuTypePage,
		Icon:      value.Pointer("iconsys-ziyuan1"),
		Path:      value.Pointer("/configure/resource"),
		Component: value.Pointer("/configure/resource/index"),
		Weight:    value.Pointer[int32](30),
		Children: []*managerMenu.RegisterMenuNode{
			api("获取资源配置", "/configure/api/resource", "GET", "configure:resource:query"),
			api("查询资源配置", "/configure/api/resources", "GET", "configure:resource:query"),
			api("创建资源配置", "/configure/api/resource", "POST", "configure:resource:add"),
			api("更新资源配置", "/configure/api/resource", "PUT", "configure:resource:update"),
			api("删除资源配置", "/configure/api/resource", "DELETE", "configure:resource:delete"),
			api("查询资源配置值", "/configure/api/resource/values", "GET", "configure:resource:value:query"),
			api("更新资源配置值", "/configure/api/resource/values", "PUT", "configure:resource:value:update"),
		},
	},
	{
		Key:       value.Pointer("ConfigureTemplate"),
		Title:     "资源模板",
		Type:      menuTypePage,
		Icon:      value.Pointer("iconsys-code-block"),
		Path:      value.Pointer("/configure/template"),
		Component: value.Pointer("/configure/template/index"),
		Weight:    value.Pointer[int32](10),
		Children: []*managerMenu.RegisterMenuNode{
			api("查询模板", "/configure/api/templates", "GET", "configure:template:query"),
			api("获取模板", "/configure/api/template", "GET", "configure:template:query"),
			api("获取当前模板", "/configure/api/template/current", "GET", "configure:template:query"),
			api("创建模板", "/configure/api/template", "POST", "configure:template:add"),
			api("切换模板", "/configure/api/template/switch", "POST", "configure:template:switch"),
			api("对比模板", "/configure/api/template/compare", "POST", "configure:template:compare"),
			api("预览模板", "/configure/api/template/preview", "POST", "configure:template:preview"),
			api("查询模板命名空间", "/configure/api/template/namespaces", "GET", "configure:template:namespace:query"),
			api("创建模板命名空间", "/configure/api/template/namespace", "POST", "configure:template:namespace:add"),
			api("更新模板命名空间", "/configure/api/template/namespace", "PUT", "configure:template:namespace:update"),
			api("删除模板命名空间", "/configure/api/template/namespace", "DELETE", "configure:template:namespace:delete"),
		},
	},
	{
		Key:       value.Pointer("ConfigureSettings"),
		Title:     "系统配置",
		Type:      menuTypePage,
		Icon:      value.Pointer("iconsys-shezhi3"),
		Path:      value.Pointer("/configure/settings"),
		Component: value.Pointer("/configure/settings/index"),
		Weight:    value.Pointer[int32](5),
		Children: []*managerMenu.RegisterMenuNode{
			api("获取系统配置", "/configure/api/settings", "GET", "configure:settings:query"),
			api("更新系统配置", "/configure/api/settings", "PUT", "configure:settings:update"),
		},
	},
}

func api(title string, path string, method string, permission string) *managerMenu.RegisterMenuNode {
	node := &managerMenu.RegisterMenuNode{
		Title:  title,
		Type:   menuTypeApi,
		Api:    value.Pointer(path),
		Method: value.Pointer(method),
		Weight: value.Pointer[int32](0),
	}
	if permission != "" {
		node.Permission = value.Pointer(permission)
	}
	return node
}
