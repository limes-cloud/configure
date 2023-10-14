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

type Template struct {
	conf *config.Config
}

func NewTemplate(conf *config.Config) *Template {
	return &Template{
		conf: conf,
	}
}

// Page 查询分页业务
func (l *Template) Page(ctx kratos.Context, in *v1.PageTemplateRequest) (*v1.PageTemplateReply, error) {
	template := model.Template{}
	list, total, err := template.Page(ctx, &model.PageOptions{
		Page:     in.Page,
		PageSize: in.PageSize,
		Scopes: func(db *gorm.DB) *gorm.DB {
			return db.Where("server_id = ?", in.ServerId)
		},
	})

	if err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.PageTemplateReply{
		Total: total,
	}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Add 添加资源
func (l *Template) Add(ctx kratos.Context, in *v1.AddTemplateRequest) (*emptypb.Empty, error) {
	template := model.Template{}
	if err := util.Transform(in, &template); err != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, template.Create(ctx)
}

// UseVersion 使用指定版本
func (l *Template) UseVersion(ctx kratos.Context, in *v1.UseTemplateVersionRequest) (*emptypb.Empty, error) {
	template := model.Template{}
	if util.Transform(in, &template) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, template.UseVersionByID(ctx, in.ServerId, in.Id)
}

// Parse 使用指定版本
// todo 版本解析
func (l *Template) Parse(ctx kratos.Context, in *v1.ParseTemplateRequest) (*v1.ParseTemplateReply, error) {
	template := model.Template{}
	if util.Transform(in, &template) != nil {
		return nil, v1.ErrorTransform()
	}

	return nil, nil // template.UseVersionByID(ctx, in.ServerId, in.Id)
}
