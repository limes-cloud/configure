package service

import (
	"github.com/limes-cloud/configure/internal/core"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/configure/internal/types"
)

type Business struct {
	repo repository.Business
}

func NewBusiness(repo repository.Business) *Business {
	return &Business{repo: repo}
}

// GetBusiness 获取指定的业务配置信息
func (u *Business) GetBusiness(ctx core.Context, id uint32) (*entity.Business, error) {
	business, err := u.repo.GetBusiness(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return business, nil
}

// ListBusiness 获取业务配置信息列表
func (u *Business) ListBusiness(ctx core.Context, req *types.ListBusinessRequest) ([]*entity.Business, uint32, error) {
	list, total, err := u.repo.ListBusiness(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateBusiness 创建业务配置信息
func (u *Business) CreateBusiness(ctx core.Context, req *entity.Business) (uint32, error) {
	id, err := u.repo.CreateBusiness(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateBusiness 更新业务配置信息
func (u *Business) UpdateBusiness(ctx core.Context, req *entity.Business) error {
	if err := u.repo.UpdateBusiness(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteBusiness 删除业务配置信息
func (u *Business) DeleteBusiness(ctx core.Context, id uint32) error {
	if err := u.repo.DeleteBusiness(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListBusinessValue 获取业务配置值信息列表
func (u *Business) ListBusinessValue(ctx core.Context, bid uint32) ([]*entity.BusinessValue, error) {
	list, err := u.repo.ListBusinessValue(ctx, &types.ListBusinessValueRequest{
		BusinessId: &bid,
	})
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	return list, nil
}

// UpdateBusinessValue 更新业务配置值信息
func (u *Business) UpdateBusinessValue(ctx core.Context, list []*entity.BusinessValue) error {
	if len(list) == 0 {
		return nil
	}

	bid := list[0].BusinessId
	business, err := u.repo.GetBusiness(ctx, bid)
	if err != nil {
		return errors.UpdateError(err.Error())
	}

	var result []*entity.BusinessValue
	for ind, item := range list {
		if item.BusinessId != bid {
			continue
		}
		v, err := item.MarshalValue(business.Type, business.Extra)
		if err != nil {
			return err
		}
		list[ind].Value = v
		result = append(result, list[ind])
	}

	if err := u.repo.UpdateBusinessValues(ctx, result); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
