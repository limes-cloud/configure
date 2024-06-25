package business

import (
	"encoding/json"
	"strconv"

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

// GetBusiness 获取指定的业务配置信息
func (u *UseCase) GetBusiness(ctx kratosx.Context, req *GetBusinessRequest) (*Business, error) {
	var (
		res *Business
		err error
	)

	if req.Id != nil {
		res, err = u.repo.GetBusiness(ctx, *req.Id)
	} else {
		return nil, errors.ParamsError()
	}

	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return res, nil
}

// ListBusiness 获取业务配置信息列表
func (u *UseCase) ListBusiness(ctx kratosx.Context, req *ListBusinessRequest) ([]*Business, uint32, error) {
	list, total, err := u.repo.ListBusiness(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateBusiness 创建业务配置信息
func (u *UseCase) CreateBusiness(ctx kratosx.Context, req *Business) (uint32, error) {
	id, err := u.repo.CreateBusiness(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateBusiness 更新业务配置信息
func (u *UseCase) UpdateBusiness(ctx kratosx.Context, req *Business) error {
	if err := u.repo.UpdateBusiness(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteBusiness 删除业务配置信息
func (u *UseCase) DeleteBusiness(ctx kratosx.Context, ids []uint32) (uint32, error) {
	total, err := u.repo.DeleteBusiness(ctx, ids)
	if err != nil {
		return 0, errors.DeleteError(err.Error())
	}
	return total, nil
}

// ListBusinessValue 获取业务配置值信息列表
func (u *UseCase) ListBusinessValue(ctx kratosx.Context, bid uint32) ([]*BusinessValue, uint32, error) {
	list, total, err := u.repo.ListBusinessValue(ctx, bid)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// UpdateBusinessValue 更新业务配置值信息
func (u *UseCase) UpdateBusinessValue(ctx kratosx.Context, list []*BusinessValue) error {
	// 检验数据类型
	business, err := u.repo.GetBusiness(ctx, list[0].BusinessId)
	if err != nil {
		return errors.GetError(err.Error())
	}

	for _, item := range list {
		var (
			isAllow = true
			value   = item.Value
		)
		switch business.Type {
		case "int":
			_, err := strconv.Atoi(value)
			isAllow = err == nil
		case "float":
			_, err := strconv.ParseFloat(value, 64)
			isAllow = err == nil
		case "string":
			isAllow = true
		case "bool":
			isAllow = value == "true" || value == "false"
		case "object":
			var m any
			isAllow = json.Unmarshal([]byte(value), &m) == nil
		default:
			isAllow = false
		}
		if !isAllow {
			return errors.BusinessValueTypeError()
		}
	}

	if err := u.repo.UpdateBusinessValues(ctx, list); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
