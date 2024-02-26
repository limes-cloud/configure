package business

import (
	"fmt"

	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/configure/internal/initiator/server"
	"github.com/limes-cloud/configure/pkg/util"
)

var businessValues = []*biz.BusinessValue{
	// Manager业务配置
	{
		Business: &biz.Business{
			ServerID:    server.Manager,
			Keyword:     "AuthSkipRoles",
			Type:        "object",
			Description: "跳过权限检测的角色列表",
		},
		Value: util.MarshalString([]string{"superAdmin"}),
	},
	{
		Business: &biz.Business{
			ServerID:    server.Manager,
			Keyword:     "DefaultUserPassword",
			Type:        "string",
			Description: "默认账号密码",
		},
		Value: `12345678`,
	},
	{
		Business: &biz.Business{
			ServerID:    server.Manager,
			Keyword:     "JwtWhitelist",
			Type:        "object",
			Description: "jwt校验白名单",
		},
		Value: util.MarshalString(map[string]any{
			"GET:/manager/v1/setting":        true,
			"POST:/manager/v1/login/captcha": true,
			"POST:/manager/v1/login":         true,
			"POST:/manager/v1/token/refresh": true,
		}),
	},
	{
		Business: &biz.Business{
			ServerID:    server.Manager,
			Keyword:     "LoginPrivatePath",
			Type:        "string",
			Description: "rsa解密私钥路径",
		},
		Value: `static/cert/login.pem`,
	},
	{
		Business: &biz.Business{
			ServerID:    server.Manager,
			Keyword:     "Setting",
			Type:        "object",
			Description: "系统设置",
		},
		Value: util.MarshalString(map[string]any{
			"name":      "go-platform",
			"title":     "go-platform 快速开发框架",
			"desc":      "开放协作，拥抱未来，插件化编程实现",
			"copyright": "Copyright © 2023 qlime.cn. All rights reserved.",
			"logo":      "http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image",
		}),
	},

	// Resource业务配置
	{
		Business: &biz.Business{
			ServerID:    server.Resource,
			Keyword:     "AcceptTypes",
			Type:        "object",
			Description: "允许上传的文件后缀",
		},
		Value: util.MarshalString([]string{"jpg", "png", "txt", "ppt", "pptx", "mp4"}),
	},
	{
		Business: &biz.Business{
			ServerID:    server.Resource,
			Keyword:     "MaxChunkCount",
			Type:        "int",
			Description: "最大切片数量",
		},
		Value: fmt.Sprint(200),
	},
	{
		Business: &biz.Business{
			ServerID:    server.Resource,
			Keyword:     "MaxChunkSize",
			Type:        "int",
			Description: "单个切片最大大小（M）",
		},
		Value: fmt.Sprint(200),
	},
	{
		Business: &biz.Business{
			ServerID:    server.Resource,
			Keyword:     "MaxSingularSize",
			Type:        "int",
			Description: "单个文件最大大小（M），超过后被切片",
		},
		Value: fmt.Sprint(400),
	},
}
