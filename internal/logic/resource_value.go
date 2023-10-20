package logic

import (
	"errors"
	"fmt"

	json "github.com/json-iterator/go"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
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
func (rv *ResourceValue) All(ctx kratos.Context, in *v1.AllResourceValueRequest) (*v1.AllResourceValueReply, error) {
	resource := model.ResourceValue{}
	list, err := resource.All(ctx, func(db *gorm.DB) *gorm.DB {
		db.Preload("Environment")
		return db.Where("resource_id =?", in.ResourceId)
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.AllResourceValueReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Add 添加资源值
func (rv *ResourceValue) Add(ctx kratos.Context, in *v1.AddResourceValueRequest) (*emptypb.Empty, error) {
	resource := model.ResourceValue{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	// 判断值是否复合字段
	if err := rv.checkValue(ctx, in.ResourceId, in.Values); err != nil {
		return nil, v1.ErrorResourceFormatValueFormat(err.Error())
	}

	return nil, resource.Create(ctx)
}

func (rv *ResourceValue) checkValue(ctx kratos.Context, rid int64, values string) error {
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
	_ = json.Unmarshal([]byte(resource.Fields), &fields)

	// 判断值是否复合字段
	for _, key := range fields {
		if m[key] == nil {
			return fmt.Errorf("资源字段%s不存在", key)
		}
	}

	return nil
}

// Update 更新资源值
func (rv *ResourceValue) Update(ctx kratos.Context, in *v1.UpdateResourceValueRequest) (*emptypb.Empty, error) {
	resource := model.ResourceValue{}
	if util.Transform(in, &resource) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, resource.Update(ctx)
}
