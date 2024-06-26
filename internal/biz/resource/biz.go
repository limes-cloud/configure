package resource

import (
	"fmt"
	"strings"

	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
)

type UseCase struct {
	conf *conf.Config
	repo Repo
}

func NewUseCase(config *conf.Config, repo Repo) *UseCase {
	return &UseCase{conf: config, repo: repo}
}

// GetResource 获取指定的资源配置信息
func (u *UseCase) GetResource(ctx kratosx.Context, req *GetResourceRequest) (*Resource, error) {
	var (
		res *Resource
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetResource(ctx, *req.Id)
	} else if req.Keyword != nil {
		res, err = u.repo.GetResourceByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListResource 获取资源配置信息列表
func (u *UseCase) ListResource(ctx kratosx.Context, req *ListResourceRequest) ([]*Resource, uint32, error) {
	list, total, err := u.repo.ListResource(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateResource 创建资源配置信息
func (u *UseCase) CreateResource(ctx kratosx.Context, req *Resource) (uint32, error) {
	id, err := u.repo.CreateResource(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateResource 更新资源配置信息
func (u *UseCase) UpdateResource(ctx kratosx.Context, req *Resource) error {
	if err := u.repo.UpdateResource(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteResource 删除资源配置信息
func (u *UseCase) DeleteResource(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteResource(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// ListResourceValue 获取业务配置值信息列表
func (u *UseCase) ListResourceValue(ctx kratosx.Context, bid uint32) ([]*ResourceValue, uint32, error) {
	list, total, err := u.repo.ListResourceValue(ctx, bid)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// UpdateResourceValue 更新业务配置值信息
func (u *UseCase) UpdateResourceValue(ctx kratosx.Context, list []*ResourceValue) error {
	// 检验数据类型
	resource, err := u.repo.GetResource(ctx, list[0].ResourceId)
	if err != nil {
		return errors.GetError(err.Error())
	}
	fields := strings.Split(resource.Fields, ",")

	for ind, item := range list {
		var (
			value = item.Value
		)
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
	}

	if err := u.repo.UpdateResourceValues(ctx, list); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
