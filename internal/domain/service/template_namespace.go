package service

import (
	"strings"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/configure/internal/types"
)

type TemplateNamespace struct {
	repo     repository.TemplateNamespace
	template repository.Template
}

func NewTemplateNamespace(repo repository.TemplateNamespace, template repository.Template) *TemplateNamespace {
	return &TemplateNamespace{
		repo:     repo,
		template: template,
	}
}

func NormalizeTemplateNamespace(namespace string) string {
	return normalizeTemplateNamespace(namespace)
}

// GetTemplateNamespace 获取指定的模板命名空间
func (u *TemplateNamespace) GetTemplateNamespace(ctx core.Context, id uint32) (*entity.TemplateNamespace, error) {
	namespace, err := u.repo.GetTemplateNamespace(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return namespace, nil
}

// ListTemplateNamespace 获取模板命名空间列表
func (u *TemplateNamespace) ListTemplateNamespace(ctx core.Context, req *types.ListTemplateNamespaceRequest) ([]*entity.TemplateNamespace, uint32, error) {
	list, total, err := u.repo.ListTemplateNamespace(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateTemplateNamespace 创建模板命名空间
func (u *TemplateNamespace) CreateTemplateNamespace(ctx core.Context, req *entity.TemplateNamespace) (uint32, error) {
	req.Namespace = normalizeTemplateNamespace(req.Namespace)
	req.Name = strings.TrimSpace(req.Name)
	id, err := u.repo.CreateTemplateNamespace(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateTemplateNamespace 更新模板命名空间
func (u *TemplateNamespace) UpdateTemplateNamespace(ctx core.Context, req *entity.TemplateNamespace) error {
	req.Namespace = normalizeTemplateNamespace(req.Namespace)
	req.Name = strings.TrimSpace(req.Name)
	old, err := u.repo.GetTemplateNamespace(ctx, req.Id)
	if err != nil {
		return errors.GetError(err.Error())
	}
	if old.Namespace != req.Namespace {
		if err := u.ensureEmptyTemplateNamespace(ctx, old.App, old.Namespace); err != nil {
			return err
		}
	}
	if err := u.repo.UpdateTemplateNamespace(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteTemplateNamespace 删除模板命名空间
func (u *TemplateNamespace) DeleteTemplateNamespace(ctx core.Context, id uint32) error {
	namespace, err := u.repo.GetTemplateNamespace(ctx, id)
	if err != nil {
		return errors.GetError(err.Error())
	}
	if err := u.ensureEmptyTemplateNamespace(ctx, namespace.App, namespace.Namespace); err != nil {
		return err
	}
	if err := u.repo.DeleteTemplateNamespace(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

func (u *TemplateNamespace) ensureEmptyTemplateNamespace(ctx core.Context, app string, namespace string) error {
	_, total, err := u.template.ListTemplate(ctx, &types.ListTemplateRequest{
		Page:      1,
		PageSize:  1,
		App:       app,
		Namespace: &namespace,
	})
	if err != nil {
		return errors.ListError(err.Error())
	}
	if total != 0 {
		return errors.DeleteError("当前命名空间下存在模板")
	}
	return nil
}

// ExistTemplateNamespace 判断模板命名空间是否存在
func (u *TemplateNamespace) ExistTemplateNamespace(ctx core.Context, app string, namespace string) error {
	if _, err := u.repo.GetTemplateNamespaceByAppAndNamespace(ctx, app, normalizeTemplateNamespace(namespace)); err != nil {
		return errors.GetError(err.Error())
	}
	return nil
}
