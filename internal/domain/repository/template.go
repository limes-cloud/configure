package repository

import (
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type Template interface {
	// GetTemplateByVersion 获取指定版本的模板
	GetTemplateByVersion(ctx core.Context, version string) (*entity.Template, error)

	// GetTemplate 获取指定ID的模板
	GetTemplate(ctx core.Context, id uint32) (*entity.Template, error)

	// CurrentTemplate 获取当前正在使用的模板
	CurrentTemplate(ctx core.Context, app string, namespace string) (*entity.Template, error)

	// ListCurrentTemplate 获取应用下当前正在使用的模板
	ListCurrentTemplate(ctx core.Context, app string) ([]*entity.Template, error)

	// ListTemplate 获取模板列表
	ListTemplate(ctx core.Context, options *types.ListTemplateRequest) ([]*entity.Template, uint32, error)

	// CreateTemplate 创建模板
	CreateTemplate(ctx core.Context, c *entity.Template) (uint32, error)

	// UpdateTemplate 更新模板
	UpdateTemplate(ctx core.Context, c *entity.Template) error

	// UseTemplate 使用指定模板
	UseTemplate(ctx core.Context, app string, namespace string, tpId uint32) error

	// DeleteTemplate 删除指定模板
	DeleteTemplate(ctx core.Context, id uint32) error
}
