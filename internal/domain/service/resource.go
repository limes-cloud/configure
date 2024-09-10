package service

import (
	"fmt"
	"strings"

	"github.com/limes-cloud/configure/internal/types"

	"github.com/limes-cloud/configure/internal/domain/entity"

	json "github.com/json-iterator/go"
	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
)

type Resource struct {
	conf       *conf.Config
	repo       repository.Resource
	permission repository.Permission
}

func NewResource(
	conf *conf.Config,
	repo repository.Resource,
	permission repository.Permission,
) *Resource {
	return &Resource{
		conf:       conf,
		repo:       repo,
		permission: permission,
	}
}

// GetResource 获取指定的资源配置信息
func (u *Resource) GetResource(ctx kratosx.Context, id uint32) (*entity.Resource, error) {
	resource, err := u.repo.GetResource(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return resource, nil
}

// GetResourceByKeyword 获取指定的资源配置信息
func (u *Resource) GetResourceByKeyword(ctx kratosx.Context, keyword string) (*entity.Resource, error) {
	resource, err := u.repo.GetResourceByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return resource, nil
}

// ListResource 获取资源配置信息列表
func (u *Resource) ListResource(ctx kratosx.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error) {
	list, total, err := u.repo.ListResource(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateResource 创建资源配置信息
func (u *Resource) CreateResource(ctx kratosx.Context, req *entity.Resource) (uint32, error) {
	id, err := u.repo.CreateResource(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateResource 更新资源配置信息
func (u *Resource) UpdateResource(ctx kratosx.Context, req *entity.Resource) error {
	if err := u.repo.UpdateResource(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteResource 删除资源配置信息
func (u *Resource) DeleteResource(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteResource(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListResourceValue 获取业务配置值信息列表
func (u *Resource) ListResourceValue(ctx kratosx.Context, rid uint32) ([]*entity.ResourceValue, error) {
	list, err := u.repo.ListResourceValue(ctx, &types.ListResourceValueRequest{
		ResourceId: &rid,
	})
	if err != nil {
		return nil, errors.ListError(err.Error())
	}

	all, scopes, err := u.permission.GetEnv(ctx)
	if err != nil {
		return nil, err
	}
	var (
		result []*entity.ResourceValue
		ir     = valx.New[uint32](scopes)
	)
	for _, item := range list {
		if all || ir.Has(item.EnvId) {
			result = append(result, item)
		}
	}

	return list, nil
}

// UpdateResourceValue 更新业务配置值信息
func (u *Resource) UpdateResourceValue(ctx kratosx.Context, list []*entity.ResourceValue) error {
	rid := list[0].ResourceId

	// 检验数据类型
	resource, err := u.repo.GetResource(ctx, list[0].ResourceId)
	if err != nil {
		return errors.GetError(err.Error())
	}
	fields := strings.Split(resource.Fields, ",")

	var result []*entity.ResourceValue
	all, scopes, err := u.permission.GetEnv(ctx)
	if err != nil {
		return err
	}
	ir := valx.New[uint32](scopes)

	for ind, item := range list {
		if item.ResourceId != rid {
			continue
		}

		value := item.Value
		m := make(map[string]any)
		if err := json.Unmarshal([]byte(value), &m); err != nil || len(m) == 0 {
			return errors.ResourceValueTypeError("字段类型必须是对象")
		}
		for _, key := range fields {
			if m[key] == nil {
				return fmt.Errorf("缺少字段%s", key)
			}
		}
		list[ind].Value, _ = json.MarshalToString(m)

		if all || ir.Has(item.EnvId) {
			result = append(result, list[ind])
		}
	}

	if err := u.repo.UpdateResourceValues(ctx, result); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
