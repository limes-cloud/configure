package service

import (
	"fmt"
	"strings"

	json "github.com/json-iterator/go"
	"github.com/limes-cloud/configure/internal/core"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/configure/internal/types"
)

type Resource struct {
	repo repository.Resource
}

func NewResource(repo repository.Resource) *Resource {
	return &Resource{repo: repo}
}

// GetResource 获取指定的资源配置信息
func (u *Resource) GetResource(ctx core.Context, id uint32) (*entity.Resource, error) {
	resource, err := u.repo.GetResource(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return resource, nil
}

// GetResourceByKeyword 获取指定的资源配置信息
func (u *Resource) GetResourceByKeyword(ctx core.Context, keyword string) (*entity.Resource, error) {
	resource, err := u.repo.GetResourceByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return resource, nil
}

// ListResource 获取资源配置信息列表
func (u *Resource) ListResource(ctx core.Context, req *types.ListResourceRequest) ([]*entity.Resource, uint32, error) {
	list, total, err := u.repo.ListResource(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateResource 创建资源配置信息
func (u *Resource) CreateResource(ctx core.Context, req *entity.Resource) (uint32, error) {
	id, err := u.repo.CreateResource(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateResource 更新资源配置信息
func (u *Resource) UpdateResource(ctx core.Context, req *entity.Resource) error {
	if err := u.repo.UpdateResource(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteResource 删除资源配置信息
func (u *Resource) DeleteResource(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteResource(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListResourceValue 获取业务配置值信息列表
func (u *Resource) ListResourceValue(ctx core.Context, rid uint32) ([]*entity.ResourceValue, error) {
	list, err := u.repo.ListResourceValue(ctx, &types.ListResourceValueRequest{
		ResourceId: &rid,
	})
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// UpdateResourceValue 更新业务配置值信息
func (u *Resource) UpdateResourceValue(ctx core.Context, list []*entity.ResourceValue) error {
	rid := list[0].ResourceId

	resource, err := u.repo.GetResource(ctx, list[0].ResourceId)
	if err != nil {
		return errors.GetError(err.Error())
	}
	fields := strings.Split(resource.Fields, ",")

	var result []*entity.ResourceValue
	for ind, item := range list {
		if item.ResourceId != rid {
			continue
		}

		valStr := item.Value
		m := make(map[string]any)
		if err := json.Unmarshal([]byte(valStr), &m); err != nil || len(m) == 0 {
			return errors.ResourceValueTypeError("字段类型必须是对象")
		}
		for _, key := range fields {
			if m[key] == nil {
				return fmt.Errorf("缺少字段%s", key)
			}
		}
		list[ind].Value, _ = json.MarshalToString(m)
		result = append(result, list[ind])
	}

	if err := u.repo.UpdateResourceValues(ctx, result); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
