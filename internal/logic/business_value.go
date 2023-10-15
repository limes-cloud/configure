package logic

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/md"
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
func (l *BusinessValue) All(ctx kratos.Context, in *v1.AllBusinessValueRequest) (*v1.AllBusinessValueReply, error) {
	bv := model.BusinessValue{}
	list, err := bv.All(ctx, func(db *gorm.DB) *gorm.DB {
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
func (l *BusinessValue) Add(ctx kratos.Context, in *v1.AddBusinessValueRequest) (*emptypb.Empty, error) {
	bv := model.BusinessValue{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if util.Transform(in, &bv) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, bv.Create(ctx)
}

// Update 更新资源
func (l *BusinessValue) Update(ctx kratos.Context, in *v1.UpdateBusinessValueRequest) (*emptypb.Empty, error) {
	bv := model.BusinessValue{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if util.Transform(in, &bv) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, bv.Update(ctx)
}
