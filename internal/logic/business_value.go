package logic

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type BusinessValue struct {
	conf *config.Config
}

func NewBusinessValue(conf *config.Config) *BusinessValue {
	return &BusinessValue{
		conf: conf,
	}
}

// All 分页资源
func (bv *BusinessValue) All(ctx kratos.Context, in *v1.AllBusinessValueRequest) (*v1.AllBusinessValueReply, error) {
	mbv := model.BusinessValue{}
	list, err := mbv.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("business_id = ?", in.BusinessId)
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.AllBusinessValueReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Add 添加资源
func (bv *BusinessValue) Add(ctx kratos.Context, in *v1.AddBusinessValueRequest) (*emptypb.Empty, error) {
	mbv := model.BusinessValue{}
	if util.Transform(in, &mbv) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, mbv.Create(ctx)
}

// Update 更新资源
func (bv *BusinessValue) Update(ctx kratos.Context, in *v1.UpdateBusinessValueRequest) (*emptypb.Empty, error) {
	mbv := model.BusinessValue{}
	if util.Transform(in, &mbv) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, mbv.Update(ctx)
}