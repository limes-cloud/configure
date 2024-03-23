package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/limes-cloud/kratosx"
	"gorm.io/gorm"

	biz "github.com/limes-cloud/configure/internal/biz/resource"
)

type repo struct {
}

func NewRepo() biz.Repo {
	return &repo{}
}

// AddResource 新建资源
func (s *repo) AddResource(ctx kratosx.Context, resource *biz.Resource) (uint32, error) {
	return resource.ID, ctx.DB().Create(resource).Error
}

// GetResource 通过ID查找指定资源
func (s *repo) GetResource(ctx kratosx.Context, id uint32) (*biz.Resource, error) {
	var resource biz.Resource
	return &resource, ctx.DB().First(&resource, "id=?", id).Error
}

// GetResourceByKeyword 通过关键词查找指定资源
func (s *repo) GetResourceByKeyword(ctx kratosx.Context, keyword string) (*biz.Resource, error) {
	var resource biz.Resource
	return &resource, ctx.DB().First(&resource, "keyword=?", keyword).Error
}

// PageResource 查询分页资源
func (s *repo) PageResource(ctx kratosx.Context, req *biz.PageResourceRequest) ([]*biz.Resource, uint32, error) {
	var list []*biz.Resource
	var total int64

	db := ctx.DB().Model(biz.Resource{}).Preload("ResourceServers")
	if req.Keyword != nil {
		db = db.Where("keyword like ?", "%"+*req.Keyword+"%")
	}
	if req.Tag != nil {
		db = db.Where("tag like ?", "%"+*req.Tag+"%")
	}
	if req.ServerId != nil {
		var ids []uint32
		if err := ctx.DB().Select("resource_id").Model(biz.ResourceServer{}).Where("server_id=?", *req.ServerId).Scan(&ids).Error; err != nil {
			return nil, 0, err
		}
		db = db.Where("id in ? or private=false", ids)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, uint32(total), err
	}

	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	return list, uint32(total), db.Find(&list).Error
}

// UpdateResource 更新指定id的资源
func (s *repo) UpdateResource(ctx kratosx.Context, resource *biz.Resource) error {
	return ctx.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(biz.ResourceServer{}, "resource_id=?", resource.ID).Error; err != nil {
			return err
		}
		return tx.Where("id=?", resource.ID).Updates(resource).Error
	})
}

// DeleteResource 删除指定id的资源
func (s *repo) DeleteResource(ctx kratosx.Context, id uint32) error {
	return ctx.DB().Delete(biz.Resource{}, "id = ?", id).Error
}

// GetResourceServerIds 获取指定资源的所有服务的id
func (s *repo) GetResourceServerIds(ctx kratosx.Context, id uint32) ([]uint32, error) {
	var ids []uint32
	return ids, ctx.DB().Model(biz.ResourceServer{}).
		Select("server_id").
		Where("resource_id=?", id).
		Scan(&ids).Error
}

// GetResourceValues 获取指定资源的所有值
func (s *repo) GetResourceValues(ctx kratosx.Context, rid uint32) ([]*biz.ResourceValue, error) {
	var list []*biz.ResourceValue
	return list, ctx.DB().Model(biz.ResourceValue{}).Find(&list, "resource_id = ?", rid).Error
}

// AllResourceValue 获取指定环境资源的所有值
func (s *repo) AllResourceValue(ctx kratosx.Context, envId, srvId uint32) ([]*biz.ResourceValue, error) {
	var ids []uint32
	if err := ctx.DB().Select("resource_id").Model(biz.ResourceServer{}).Where("server_id=?", srvId).Scan(&ids).Error; err != nil {
		return nil, err
	}
	var ids1 []uint32
	if err := ctx.DB().Select("id").Model(biz.Resource{}).Where("private=false").Scan(&ids1).Error; err != nil {
		return nil, err
	}
	ids = append(ids, ids1...)

	var list []*biz.ResourceValue
	return list, ctx.DB().Model(biz.ResourceValue{}).Preload("Resource").Find(&list, "env_id=? and resource_id in ?", envId, ids).Error
}

// UpdateResourceValue 更新指定资源的值
func (s *repo) UpdateResourceValue(ctx kratosx.Context, rv *biz.ResourceValue) error {
	if rv.ID == 0 {
		oldRv := &biz.ResourceValue{}
		if err := ctx.DB().First(oldRv, "resource_id=? and env_id=?", rv.ResourceId, rv.EnvId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ctx.DB().Create(rv).Error
			}
			return err
		}
		rv.ID = oldRv.ID
	}
	return ctx.DB().Updates(rv).Error
}

// AllResourceByServer 获取指定服务的所有资源
func (s *repo) AllResourceByServer(ctx kratosx.Context, id uint32) ([]*biz.Resource, error) {
	var ids []uint32
	if err := ctx.DB().Select("resource_id").
		Model(biz.ResourceServer{}).
		Where("server_id=?", id).
		Scan(&ids).Error; err != nil {
		return nil, err
	}

	var list []*biz.Resource
	return list, ctx.DB().Model(biz.Resource{}).Where("id in ? or private=false", ids).Find(&list).Error
}

// AllResourceField 获取指定服务的全部可用资源字段
func (s *repo) AllResourceField(ctx kratosx.Context, sid uint32) ([]string, error) {
	list, err := s.AllResourceByServer(ctx, sid)
	if err != nil {
		return nil, err
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

func (s *repo) CheckResourceValue(ctx kratosx.Context, rid uint32, value string) error {
	m := make(map[string]any)
	if err := json.Unmarshal([]byte(value), &m); err != nil {
		return err
	}
	if len(m) == 0 {
		return errors.New("类型必须为集合")
	}

	rs, err := s.GetResource(ctx, rid)
	if err != nil {
		return err
	}
	fields := strings.Split(rs.Fields, ",")

	// 判断值是否复合字段
	for _, key := range fields {
		if m[key] == nil {
			return fmt.Errorf("字段%s不存在", key)
		}
	}
	return nil
}
