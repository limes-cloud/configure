package dbs

import (
	"sync"

	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/types"
)

type TemplateNamespace struct {
}

var (
	templateNamespaceIns  *TemplateNamespace
	templateNamespaceOnce sync.Once
)

func NewTemplateNamespace() *TemplateNamespace {
	templateNamespaceOnce.Do(func() {
		templateNamespaceIns = &TemplateNamespace{}
	})
	return templateNamespaceIns
}

// GetTemplateNamespace 获取指定的数据
func (r TemplateNamespace) GetTemplateNamespace(ctx core.Context, id uint32) (*entity.TemplateNamespace, error) {
	var (
		namespace = entity.TemplateNamespace{}
		fs        = []string{"*"}
	)
	return &namespace, ctx.DB().Select(fs).First(&namespace, id).Error
}

// GetTemplateNamespaceByAppAndNamespace 获取指定应用下的模板命名空间
func (r TemplateNamespace) GetTemplateNamespaceByAppAndNamespace(ctx core.Context, app string, namespace string) (*entity.TemplateNamespace, error) {
	var (
		item = entity.TemplateNamespace{}
		fs   = []string{"*"}
	)
	return &item, ctx.DB().Select(fs).Where("app=? and namespace=?", app, namespace).First(&item).Error
}

// ListTemplateNamespace 获取列表
func (r TemplateNamespace) ListTemplateNamespace(ctx core.Context, req *types.ListTemplateNamespaceRequest) ([]*entity.TemplateNamespace, uint32, error) {
	var (
		list  []*entity.TemplateNamespace
		total int64
	)

	db := ctx.DB().
		Model(entity.TemplateNamespace{}).
		Select("*").
		Where("app=?", req.App)
	if req.Namespace != nil && *req.Namespace != "" {
		db = db.Where("namespace=?", *req.Namespace)
	}
	if req.Name != nil && *req.Name != "" {
		db = db.Where("name LIKE ?", *req.Name+"%")
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	db = db.Order("sort asc,id asc")
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// CreateTemplateNamespace 创建数据
func (r TemplateNamespace) CreateTemplateNamespace(ctx core.Context, namespace *entity.TemplateNamespace) (uint32, error) {
	return namespace.Id, ctx.DB().Create(namespace).Error
}

// UpdateTemplateNamespace 更新数据
func (r TemplateNamespace) UpdateTemplateNamespace(ctx core.Context, namespace *entity.TemplateNamespace) error {
	return ctx.DB().Model(entity.TemplateNamespace{}).
		Where("id = ?", namespace.Id).
		Updates(map[string]any{
			"namespace":   namespace.Namespace,
			"name":        namespace.Name,
			"description": namespace.Description,
			"sort":        namespace.Sort,
		}).Error
}

// DeleteTemplateNamespace 删除数据
func (r TemplateNamespace) DeleteTemplateNamespace(ctx core.Context, id uint32) error {
	return ctx.DB().Where("id=?", id).Delete(&entity.TemplateNamespace{}).Error
}
