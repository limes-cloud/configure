package repository

import (
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type TemplateNamespace interface {
	// GetTemplateNamespace 获取指定的模板命名空间
	GetTemplateNamespace(ctx core.Context, id uint32) (*entity.TemplateNamespace, error)

	// GetTemplateNamespaceByAppAndNamespace 获取指定应用下的模板命名空间
	GetTemplateNamespaceByAppAndNamespace(ctx core.Context, app string, namespace string) (*entity.TemplateNamespace, error)

	// ListTemplateNamespace 获取应用下的模板命名空间列表
	ListTemplateNamespace(ctx core.Context, req *types.ListTemplateNamespaceRequest) ([]*entity.TemplateNamespace, uint32, error)

	// CreateTemplateNamespace 创建模板命名空间
	CreateTemplateNamespace(ctx core.Context, req *entity.TemplateNamespace) (uint32, error)

	// UpdateTemplateNamespace 更新模板命名空间
	UpdateTemplateNamespace(ctx core.Context, req *entity.TemplateNamespace) error

	// DeleteTemplateNamespace 删除模板命名空间
	DeleteTemplateNamespace(ctx core.Context, id uint32) error
}
