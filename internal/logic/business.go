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

type Business struct {
	conf *config.Config
}

func NewBusiness(conf *config.Config) *Business {
	return &Business{
		conf: conf,
	}
}

// Page 查询分页业务
func (l *Business) Page(ctx kratos.Context, in *v1.PageBusinessRequest) (*v1.PageBusinessReply, error) {
	business := model.Business{}
	list, total, err := business.Page(ctx, &model.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			if in.Keyword != nil {
				db = db.Where("keyword like ?", "%"+*in.Keyword+"%")
			}
			return db
		},
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.PageBusinessReply{
		Total: total,
	}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Add 添加资源
func (l *Business) Add(ctx kratos.Context, in *v1.AddBusinessRequest) (*emptypb.Empty, error) {
	business := model.Business{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if err := util.Transform(in, &business); err != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, business.Create(ctx)
}

// Update 更新资源
func (l *Business) Update(ctx kratos.Context, in *v1.UpdateBusinessRequest) (*emptypb.Empty, error) {
	business := model.Business{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if util.Transform(in, &business) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, business.Update(ctx)
}

// Delete 删除资源
func (l *Business) Delete(ctx kratos.Context, in *v1.DeleteBusinessRequest) (*emptypb.Empty, error) {
	business := model.Business{}
	return nil, business.DeleteByID(ctx, in.Id)
}
