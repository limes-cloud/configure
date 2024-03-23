package resource

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/config"
)

type UseCase struct {
	config *config.Config
	repo   Repo
}

func NewUseCase(config *config.Config, repo Repo) *UseCase {
	return &UseCase{config: config, repo: repo}
}

// GetResource 获取指定资源信息
func (u *UseCase) GetResource(ctx kratosx.Context, id uint32) (*Resource, error) {
	resource, err := u.repo.GetResource(ctx, id)
	if err != nil {
		return nil, errors.NotRecordError()
	}
	return resource, nil
}

// GetResourceByKeyword 获取指定标识的资源信息
func (u *UseCase) GetResourceByKeyword(ctx kratosx.Context, keyword string) (*Resource, error) {
	resource, err := u.repo.GetResourceByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.NotRecordError()
	}
	return resource, nil
}

// PageResource 获取分页资源信息
func (u *UseCase) PageResource(ctx kratosx.Context, req *PageResourceRequest) ([]*Resource, uint32, error) {
	list, total, err := u.repo.PageResource(ctx, req)
	if err != nil {
		return nil, 0, errors.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// AddResource 添加资源信息
func (u *UseCase) AddResource(ctx kratosx.Context, resource *Resource) (uint32, error) {
	id, err := u.repo.AddResource(ctx, resource)
	if err != nil {
		return 0, errors.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// UpdateResource 更新资源信息
func (u *UseCase) UpdateResource(ctx kratosx.Context, resource *Resource) error {
	if err := u.repo.UpdateResource(ctx, resource); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// DeleteResource 删除资源信息
func (u *UseCase) DeleteResource(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteResource(ctx, id); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// AllResourceField 获取指定服务资源的字段
func (u *UseCase) AllResourceField(ctx kratosx.Context, sid uint32) ([]string, error) {
	list, err := u.repo.AllResourceField(ctx, sid)
	if err != nil {
		return nil, errors.DatabaseErrorFormat(err.Error())
	}
	return list, nil
}

// GetResourceServerIds 获取指定资源的关联服务id
func (u *UseCase) GetResourceServerIds(ctx kratosx.Context, rid uint32) ([]uint32, error) {
	return u.repo.GetResourceServerIds(ctx, rid)
}

// GetResourceValues 获取指定资源的所有环境值
func (u *UseCase) GetResourceValues(ctx kratosx.Context, rid uint32) ([]*ResourceValue, error) {
	list, err := u.repo.GetResourceValues(ctx, rid)
	if err != nil {
		return nil, errors.DatabaseErrorFormat(err.Error())
	}
	return list, nil
}

// UpdateResourceValue 更新指定资源的值
func (u *UseCase) UpdateResourceValue(ctx kratosx.Context, rv *ResourceValue) error {
	if err := u.repo.CheckResourceValue(ctx, rv.ResourceId, rv.Value); err != nil {
		return err
	}
	if err := u.repo.UpdateResourceValue(ctx, rv); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}
