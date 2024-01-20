package biz

import (
	"encoding/json"
	"fmt"
	"strings"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/biz/types"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type Resource struct {
	ktypes.BaseModel
	Keyword        string            `json:"keyword" gorm:"not null;type:char(32) binary;comment:变量标识"`
	Description    string            `json:"description" gorm:"not null;size:128;comment:变量描述"`
	Fields         string            `json:"fields" gorm:"not null;size:256;comment:变量字段集合"`
	Tag            string            `json:"tag" gorm:"not null;size:32;comment:变量标签"`
	Private        *bool             `json:"private" gorm:"default:false;comment:是否私有"`
	ResourceServer []*ResourceServer `json:"resource_server" gorm:"constraint:onDelete:cascade"`
	ResourceValue  []*ResourceValue  `json:"resource_value" gorm:"constraint:onDelete:cascade"`
}

type ResourceValue struct {
	ktypes.BaseModel
	EnvID      uint32    `json:"env_id" gorm:"uniqueIndex:er;not null;comment:环境id"`
	ResourceID uint32    `json:"resource_id" gorm:"uniqueIndex:er;not null;comment:资源id"`
	Value      string    `json:"value" gorm:"not null;type:text;comment:资源id"`
	Env        *Env      `json:"env" gorm:"constraint:onDelete:cascade"`
	Resource   *Resource `json:"resource" gorm:"constraint:onDelete:cascade"`
}

type ResourceServer struct {
	ServerID   uint32    `json:"server_id" gorm:"uniqueIndex:sr;not null;comment:服务id"`
	ResourceID uint32    `json:"resource_id" gorm:"uniqueIndex:sr;not null;comment:资源id"`
	Server     *Server   `json:"server" gorm:"constraint:onDelete:cascade"`
	Resource   *Resource `json:"resource" gorm:"constraint:onDelete:cascade"`
}

type ResourceRepo interface {
	Get(ctx kratosx.Context, id uint32) (*Resource, error)
	GetByKeyword(ctx kratosx.Context, key string) (*Resource, error)
	PageResource(ctx kratosx.Context, req *types.PageResourceRequest) ([]*Resource, uint32, error)
	Create(ctx kratosx.Context, c *Resource) (uint32, error)
	Update(ctx kratosx.Context, c *Resource) error
	Delete(ctx kratosx.Context, uint322 uint32) error
	GetValues(ctx kratosx.Context, rid uint32) ([]*ResourceValue, error)
	GetEnvValues(ctx kratosx.Context, eid, sid uint32) ([]*ResourceValue, error)
	UpdateValue(ctx kratosx.Context, rv *ResourceValue) error
	PageServerResource(ctx kratosx.Context, id uint32, options *ktypes.PageOptions) ([]*Resource, uint32, error)
	AllServerResource(ctx kratosx.Context, id uint32) ([]*Resource, error)
	AllResourceServerId(ctx kratosx.Context, id uint32) ([]uint32, error)
}

type ResourceUseCase struct {
	config  *config.Config
	rsRepo  ResourceRepo
	srvRepo ServerRepo
}

func NewResourceUseCase(config *config.Config, rsRepo ResourceRepo) *ResourceUseCase {
	return &ResourceUseCase{config: config, rsRepo: rsRepo}
}

// Get 获取指定资源信息
func (u *ResourceUseCase) Get(ctx kratosx.Context, id uint32) (*Resource, error) {
	resource, err := u.rsRepo.Get(ctx, id)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return resource, nil
}

// GetByKeyword 获取指定标识的资源信息
func (u *ResourceUseCase) GetByKeyword(ctx kratosx.Context, keyword string) (*Resource, error) {
	resource, err := u.rsRepo.GetByKeyword(ctx, keyword)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return resource, nil
}

// Page 获取分页资源信息
func (u *ResourceUseCase) Page(ctx kratosx.Context, req *types.PageResourceRequest) ([]*Resource, uint32, error) {
	list, total, err := u.rsRepo.PageResource(ctx, req)
	if err != nil {
		return nil, 0, v1.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// Add 添加资源信息
func (u *ResourceUseCase) Add(ctx kratosx.Context, resource *Resource) (uint32, error) {
	if resource.ResourceValue != nil {
		for _, rv := range resource.ResourceValue {
			if err := u.CheckResourceValue(ctx, rv.ResourceID, rv.Value); err != nil {
				return 0, err
			}
		}
	}
	id, err := u.rsRepo.Create(ctx, resource)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新资源信息
func (u *ResourceUseCase) Update(ctx kratosx.Context, resource *Resource) error {
	if err := u.rsRepo.Update(ctx, resource); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除资源信息
func (u *ResourceUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.rsRepo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// PageServerResource 获取指定服务的分页资源
func (u *ResourceUseCase) PageServerResource(ctx kratosx.Context, req *types.PageServerResourceRequest) ([]*Resource, uint32, error) {
	list, total, err := u.rsRepo.PageServerResource(ctx, req.ServerID, &ktypes.PageOptions{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, 0, v1.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// AllServerResourceField 获取指定服务资源的字段
func (u *ResourceUseCase) AllServerResourceField(ctx kratosx.Context, sid uint32) ([]string, error) {
	list, err := u.rsRepo.AllServerResource(ctx, sid)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	var arr []string
	for _, item := range list {
		fields := strings.Split(item.Fields, ",")
		for _, val := range fields {
			arr = append(arr, item.Keyword+"."+val)
		}
	}
	return arr, nil
}

// AllResourceServerIds 获取指定资源的关联服务id
func (u *ResourceUseCase) AllResourceServerIds(ctx kratosx.Context, rid uint32) ([]uint32, error) {
	return u.rsRepo.AllResourceServerId(ctx, rid)
}

// AllResourceValue 获取指定资源的所有环境值
func (u *ResourceUseCase) AllResourceValue(ctx kratosx.Context, rid uint32) ([]*ResourceValue, error) {
	list, err := u.rsRepo.GetValues(ctx, rid)
	if err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	return list, nil
}

// UpdateResourceValue 更新指定资源的值
func (u *ResourceUseCase) UpdateResourceValue(ctx kratosx.Context, rv *ResourceValue) error {
	if err := u.CheckResourceValue(ctx, rv.ResourceID, rv.Value); err != nil {
		return err
	}
	if err := u.rsRepo.UpdateValue(ctx, rv); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// CheckResourceValue 检测资源值是否合法
func (u *ResourceUseCase) CheckResourceValue(ctx kratosx.Context, rid uint32, values string) error {
	m := make(map[string]any)
	if err := json.Unmarshal([]byte(values), &m); err != nil {
		return err
	}
	if len(values) == 0 {
		return v1.ResourceValueErrorFormat("类型必须为集合")
	}

	rs, err := u.rsRepo.Get(ctx, rid)
	if err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	fields := strings.Split(rs.Fields, ",")

	// 判断值是否复合字段
	for _, key := range fields {
		if m[key] == nil {
			return v1.ResourceValueErrorFormat(fmt.Sprintf("字段%s不存在", key))
		}
	}
	return nil
}
