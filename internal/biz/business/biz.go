package business

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

// GetBusiness 获取指定业务变量信息
func (u *UseCase) GetBusiness(ctx kratosx.Context, id uint32) (*Business, error) {
	business, err := u.repo.GetBusiness(ctx, id)
	if err != nil {
		return nil, errors.NotRecordError()
	}
	return business, nil
}

// GetBusinessByKeyword 获取指定标识的业务变量信息
func (u *UseCase) GetBusinessByKeyword(ctx kratosx.Context, keyword string) (*Business, error) {
	business, err := u.repo.GetBusinessByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.NotRecordError()
	}
	return business, nil
}

// AddBusiness 添加业务变量信息
func (u *UseCase) AddBusiness(ctx kratosx.Context, business *Business) (uint32, error) {
	id, err := u.repo.AddBusiness(ctx, business)
	if err != nil {
		return 0, errors.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// UpdateBusiness 更新业务变量信息
func (u *UseCase) UpdateBusiness(ctx kratosx.Context, business *Business) error {
	if err := u.repo.UpdateBusiness(ctx, business); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// DeleteBusiness 删除业务变量信息
func (u *UseCase) DeleteBusiness(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteBusiness(ctx, id); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// PageBusiness 获取指定服务的业务变量
func (u *UseCase) PageBusiness(ctx kratosx.Context, req *PageBusinessRequest) ([]*Business, uint32, error) {
	list, total, err := u.repo.PageBusiness(ctx, req)
	if err != nil {
		return nil, 0, errors.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// GetBusinessValues 获取指定业务变量的所有环境值
func (u *UseCase) GetBusinessValues(ctx kratosx.Context, bid uint32) ([]*BusinessValue, error) {
	list, err := u.repo.GetBusinessValues(ctx, bid)
	if err != nil {
		return nil, errors.DatabaseErrorFormat(err.Error())
	}
	return list, nil
}

// UpdateBusinessValue 更新指定业务变量的值
func (u *UseCase) UpdateBusinessValue(ctx kratosx.Context, rv *BusinessValue) error {
	if u.repo.CheckBusinessValue(ctx, rv.BusinessId, rv.Value) {
		return errors.BusinessValueError()
	}
	if err := u.repo.UpdateBusinessValue(ctx, rv); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}
