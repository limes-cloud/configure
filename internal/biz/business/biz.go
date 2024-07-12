package business

import (
	"strconv"

	json "github.com/json-iterator/go"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/pkg/permission"
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
	if !permission.HasServer(ctx, req.ServerId) {
		return nil, 0, errors.NotPermissionError()
	}

	list, total, err := u.repo.ListBusiness(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateBusiness 创建业务配置信息
func (u *UseCase) CreateBusiness(ctx kratosx.Context, req *Business) (uint32, error) {
	if !permission.HasServer(ctx, req.ServerId) {
		return 0, errors.NotPermissionError()
	}

	id, err := u.repo.CreateBusiness(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateBusiness 更新业务配置信息
func (u *UseCase) UpdateBusiness(ctx kratosx.Context, req *Business) error {
	if !permission.HasServer(ctx, req.ServerId) {
		return errors.NotPermissionError()
	}

	if err := u.repo.UpdateBusiness(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteBusiness 删除业务配置信息
func (u *UseCase) DeleteBusiness(ctx kratosx.Context, id uint32) error {
	business, err := u.repo.GetBusiness(ctx, id)
	if err != nil {
		return errors.DeleteError(err.Error())
	}

	if !permission.HasServer(ctx, business.ServerId) {
		return errors.NotPermissionError()
	}

	if err := u.repo.DeleteBusiness(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}

// ListBusinessValue 获取业务配置值信息列表
func (u *UseCase) ListBusinessValue(ctx kratosx.Context, bid uint32) ([]*BusinessValue, uint32, error) {
	business, err := u.repo.GetBusiness(ctx, bid)
	if err != nil {
		return nil, 0, errors.DeleteError(err.Error())
	}

	if !permission.HasServer(ctx, business.ServerId) {
		return nil, 0, errors.NotPermissionError()
	}

	list, total, err := u.repo.ListBusinessValue(ctx, bid)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	all, scopes, err := permission.GetEnv(ctx)
	if err != nil {
		return nil, 0, err
	}
	if !all {
		var result []*BusinessValue
		ir := valx.New[uint32](scopes)
		for _, item := range list {
			if ir.Has(item.EnvId) {
				result = append(result, item)
			}
		}
		list = result
	}

	return list, total, nil
}

// UpdateBusinessValue 更新业务配置值信息
func (u *UseCase) UpdateBusinessValue(ctx kratosx.Context, list []*BusinessValue) error {
	if len(list) == 0 {
		return nil
	}

	bid := list[0].BusinessId
	business, err := u.repo.GetBusiness(ctx, bid)
	if err != nil {
		return errors.UpdateError(err.Error())
	}

	if !permission.HasServer(ctx, business.ServerId) {
		return errors.NotPermissionError()
	}

	var result []*BusinessValue
	all, scopes, err := permission.GetEnv(ctx)
	if err != nil {
		return err
	}
	ir := valx.New[uint32](scopes)

	for ind, item := range list {
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
			list[ind].Value, _ = json.MarshalToString(m)
		default:
			isAllow = false
		}
		if !isAllow {
			return errors.BusinessValueTypeError()
		}
		if all || ir.Has(item.EnvId) {
			result = append(result, list[ind])
		}
	}

	if err := u.repo.UpdateBusinessValues(ctx, result); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}
