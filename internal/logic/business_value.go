package logic

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratosx"
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
func (bv *BusinessValue) All(ctx kratosx.Context, in *v1.AllBusinessValueRequest) (*v1.AllBusinessValueReply, error) {
	mbv := model.BusinessValue{}
	list, err := mbv.All(ctx, func(db *gorm.DB) *gorm.DB {
		return db.Where("business_id = ?", in.BusinessId)
	})

	if err != nil {
		return nil, v1.DatabaseError()
	}
	reply := v1.AllBusinessValueReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

// Update 更新资源
func (bv *BusinessValue) Update(ctx kratosx.Context, in *v1.UpdateBusinessValueRequest) (*emptypb.Empty, error) {
	var bvs []*model.BusinessValue
	for _, item := range in.List {
		temp := model.BusinessValue{
			BusinessID:    in.BusinessId,
			EnvironmentID: item.EnvironmentId,
			Value:         item.Value,
		}
		bvs = append(bvs, &temp)
	}

	value := model.BusinessValue{}
	return nil, value.Creates(ctx, in.BusinessId, bvs)
}
