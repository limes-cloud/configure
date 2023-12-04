package logic

import (
	"errors"
	"fmt"
	"strings"

	json "github.com/json-iterator/go"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type ResourceValue struct {
	conf *config.Config
}

func NewResourceValue(conf *config.Config) *ResourceValue {
	return &ResourceValue{
		conf: conf,
	}
}

// All 分页字段值
func (rv *ResourceValue) All(ctx kratosx.Context, in *v1.AllResourceValueRequest) (*v1.AllResourceValueReply, error) {
	resource := model.ResourceValue{}
	list, err := resource.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("resource_id =?", in.ResourceId)
	})

	if err != nil {
		return nil, v1.DatabaseError()
	}
	reply := v1.AllResourceValueReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (rv *ResourceValue) checkValue(ctx kratosx.Context, rid uint32, values string) error {
	m := make(map[string]any)
	if err := json.Unmarshal([]byte(values), &m); err != nil {
		return err
	}
	if len(values) == 0 {
		return errors.New("值类型必须是一个map")
	}

	var fields []string
	resource := model.Resource{}
	if err := resource.OneByID(ctx, rid); err != nil {
		return err
	}
	fields = strings.Split(resource.Fields, ",")

	// 判断值是否复合字段
	for _, key := range fields {
		if m[key] == nil {
			return fmt.Errorf("资源字段%s不存在", key)
		}
	}

	return nil
}

// Update 更新资源值
func (rv *ResourceValue) Update(ctx kratosx.Context, in *v1.UpdateResourceValueRequest) (*emptypb.Empty, error) {
	var rvs []*model.ResourceValue
	// 遍历
	for _, item := range in.List {
		temp := model.ResourceValue{
			ResourceID:    in.ResourceId,
			EnvironmentID: item.EnvironmentId,
			Values:        item.Values,
		}
		rvs = append(rvs, &temp)

		if err := rv.checkValue(ctx, in.ResourceId, item.Values); err != nil {
			return nil, v1.ResourceFormatValueErrorFormat("[%s]%s", item.EnvKeyword, err.Error())
		}
	}

	value := model.ResourceValue{}
	return nil, value.Creates(ctx, in.ResourceId, rvs)
}
