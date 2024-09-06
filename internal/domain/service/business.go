package service

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/configure/internal/types"
)

type BusinessService struct {
	conf       *conf.Config
	repo       repository.BusinessRepository
	permission repository.PermissionRepository
}

func NewBusinessService(
	conf *conf.Config,
	repo repository.BusinessRepository,
	permission repository.PermissionRepository,
) *BusinessService {
	return &BusinessService{
		conf:       conf,
		repo:       repo,
		permission: permission,
	}
}

// ListBusiness 获取业务配置信息列表
func (u *BusinessService) ListBusiness(ctx kratosx.Context, req *types.ListBusinessRequest) ([]*entity.Business, uint32, error) {
	if !u.permission.HasServer(ctx, req.ServerId) {
		return nil, 0, errors.NotPermissionError()
	}

	list, total, err := u.repo.ListBusiness(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateBusiness 创建业务配置信息
func (u *BusinessService) CreateBusiness(ctx kratosx.Context, req *entity.Business) (uint32, error) {
	if !u.permission.HasServer(ctx, req.ServerId) {
		return 0, errors.NotPermissionError()
	}

	id, err := u.repo.CreateBusiness(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateBusiness 更新业务配置信息
func (u *BusinessService) UpdateBusiness(ctx kratosx.Context, req *entity.Business) error {
	if !u.permission.HasServer(ctx, req.ServerId) {
		return errors.NotPermissionError()
	}

	if err := u.repo.UpdateBusiness(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteBusiness 删除业务配置信息
func (u *BusinessService) DeleteBusiness(ctx kratosx.Context, id uint32) error {
	business, err := u.repo.GetBusiness(ctx, id)
	if err != nil {
		return errors.DeleteError(err.Error())
	}

	if !u.permission.HasServer(ctx, business.ServerId) {
		return errors.NotPermissionError()
	}

	if err := u.repo.DeleteBusiness(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListBusinessValue 获取业务配置值信息列表
func (u *BusinessService) ListBusinessValue(ctx kratosx.Context, bid uint32) ([]*entity.BusinessValue, error) {
	business, err := u.repo.GetBusiness(ctx, bid)
	if err != nil {
		return nil, errors.DeleteError(err.Error())
	}

	if !u.permission.HasServer(ctx, business.ServerId) {
		return nil, errors.NotPermissionError()
	}

	list, err := u.repo.ListBusinessValue(ctx, &types.ListBusinessValueRequest{
		BusinessId: &bid,
	})
	if err != nil {
		return nil, errors.ListError(err.Error())
	}
	all, scopes, err := u.permission.GetEnv(ctx)
	if err != nil {
		return nil, err
	}
	if !all {
		var result []*entity.BusinessValue
		ir := valx.New[uint32](scopes)
		for _, item := range list {
			if ir.Has(item.EnvId) {
				result = append(result, item)
			}
		}
		list = result
	}

	return list, nil
}

// UpdateBusinessValue 更新业务配置值信息
func (u *BusinessService) UpdateBusinessValue(ctx kratosx.Context, list []*entity.BusinessValue) error {
	if len(list) == 0 {
		return nil
	}

	bid := list[0].BusinessId
	business, err := u.repo.GetBusiness(ctx, bid)
	if err != nil {
		return errors.UpdateError(err.Error())
	}

	if !u.permission.HasServer(ctx, business.ServerId) {
		return errors.NotPermissionError()
	}

	var result []*entity.BusinessValue
	all, scopes, err := u.permission.GetEnv(ctx)
	if err != nil {
		return err
	}

	ir := valx.New(scopes)
	for ind, item := range list {
		// 去除掉不一致的数据
		if item.BusinessId != bid {
			continue
		}

		// 检验类型&序列化值
		value, err := item.MarshalValue(business.Type)
		if err != nil {
			return err
		}
		list[ind].Value = value

		// 检验是否拥有权限
		if all || ir.Has(item.EnvId) {
			result = append(result, list[ind])
		}
	}

	// 更新数据
	if err := u.repo.UpdateBusinessValues(ctx, result); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
