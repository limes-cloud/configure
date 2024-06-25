package template

import (
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"gorm.io/gorm"

	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/factory"
	"github.com/limes-cloud/configure/internal/pkg"
)

type UseCase struct {
	config  *conf.Config
	repo    Repo
	factory factory.Factory
}

func NewUseCase(config *conf.Config, repo Repo, factory factory.Factory) *UseCase {
	return &UseCase{
		config:  config,
		repo:    repo,
		factory: factory,
	}
}

// CurrentTemplate 获取当前版本的配置信息
func (t *UseCase) CurrentTemplate(ctx kratosx.Context, srvId uint32) (*Template, error) {
	template, err := t.repo.CurrentTemplate(ctx, srvId)
	if err != nil {
		if gormtranserror.Is(gorm.ErrRecordNotFound, err) {
			return nil, errors.ServerNotExistTemplateError()
		}
		return nil, errors.GetError(err.Error())
	}
	return template, nil
}

// GetTemplate 获取指定模板信息
func (t *UseCase) GetTemplate(ctx kratosx.Context, id uint32) (*Template, error) {
	template, err := t.repo.GetTemplate(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return template, nil
}

// ListTemplate 获取分页模板信息
func (t *UseCase) ListTemplate(ctx kratosx.Context, req *ListTemplateRequest) ([]*Template, uint32, error) {
	list, total, err := t.repo.ListTemplate(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateTemplate 添加模板信息
func (t *UseCase) CreateTemplate(ctx kratosx.Context, template *Template) (uint32, error) {
	// 序列化数据，判断格式是否正确
	otc := map[string]any{}
	oe := encoding.GetCodec(template.Format)
	if err := oe.Unmarshal([]byte(template.Content), &otc); err != nil {
		return 0, errors.ParseTemplateError(err.Error())
	}

	// 当前的版本
	template.Version = crypto.MD5ToUpper([]byte(template.Content))

	// 查找是否存在相同的版本
	if _, err := t.repo.GetTemplateByVersion(ctx, template.Version); err == nil {
		return 0, errors.TemplateVersionExistError()
	}

	// 进行模板检查是否正确
	if err := t.factory.CheckTemplate(ctx, template.ServerId, template.Content); err != nil {
		return 0, errors.ParseTemplateError(err.Error())
	}

	// 获取当前的模板
	current, err := t.repo.CurrentTemplate(ctx, template.ServerId)
	// 获取变更对比信息
	if err == nil {
		compare, err := t.CompareTemplate(ctx, &CompareTemplateRequest{
			Id:      current.Id,
			Format:  template.Format,
			Content: template.Content,
		})
		if err != nil {
			return 0, errors.ParseTemplateError(err.Error())
		}
		compareByte, _ := json.Marshal(compare)
		template.Compare = string(compareByte)
	}

	// 数据入库
	id, err := t.repo.CreateTemplate(ctx, template)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}

	// 使用当前配置
	_ = t.repo.UseTemplate(ctx, template.ServerId, id)
	return id, nil
}

// PreviewTemplate 使用指定版本配置
func (t *UseCase) PreviewTemplate(ctx kratosx.Context, req *PreviewTemplateRequest) (*PreviewTemplateReply, error) {
	content, err := t.factory.ParseByContent(ctx, &factory.ParseByContentRequest{
		EnvId:    req.EnvId,
		ServerId: req.ServerId,
		Format:   req.Format,
		Content:  req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &PreviewTemplateReply{
		Content: content,
		Format:  req.Format,
	}, err
}

// ParseTemplate 使用指定版本配置
func (t *UseCase) ParseTemplate(ctx kratosx.Context, req *ParseTemplateRequest) (*ParseTemplateReply, error) {
	tp, err := t.repo.CurrentTemplate(ctx, req.ServerId)
	if err != nil {
		return nil, errors.DatabaseError(err.Error())
	}

	reply := &ParseTemplateReply{
		Format: tp.Format,
	}

	reply.Content, err = t.factory.ParseByContent(ctx, &factory.ParseByContentRequest{
		EnvId:    req.EnvId,
		ServerId: req.ServerId,
		Format:   tp.Format,
		Content:  tp.Content,
	})

	return reply, err
}

// SwitchTemplate 切换指定版本信息
func (t *UseCase) SwitchTemplate(ctx kratosx.Context, srvId, tpId uint32) error {
	if err := t.repo.UseTemplate(ctx, srvId, tpId); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// UpdateTemplate 更新模板信息
func (t *UseCase) UpdateTemplate(ctx kratosx.Context, template *Template) error {
	if err := t.repo.UpdateTemplate(ctx, template); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteUpdateTemplate 删除模板信息
func (t *UseCase) DeleteUpdateTemplate(ctx kratosx.Context, id uint32) error {
	if err := t.repo.DeleteTemplate(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// CompareTemplate 对比变更细节
func (t *UseCase) CompareTemplate(ctx kratosx.Context, req *CompareTemplateRequest) ([]*CompareTemplateReply, error) {
	// 查询版本
	template, err := t.repo.GetTemplate(ctx, req.Id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}

	// 解析当前版本数据
	var oldKeys []string
	otc := map[string]any{}

	oe := encoding.GetCodec(template.Format)
	if err := oe.Unmarshal([]byte(template.Content), &otc); err != nil {
		return nil, errors.ParseTemplateError()
	}
	for key := range otc {
		oldKeys = append(oldKeys, key)
	}

	// 解析上传的模板
	var curKeys []string
	ctc := map[string]any{}

	ce := encoding.GetCodec(req.Format)
	if err := ce.Unmarshal([]byte(req.Content), &ctc); err != nil {
		return nil, errors.ParseTemplateError()
	}
	for key := range ctc {
		curKeys = append(curKeys, key)
	}

	var reply []*CompareTemplateReply

	// 新增字段
	addKeys := pkg.Diff(curKeys, oldKeys)
	for _, key := range addKeys {
		val := ""
		switch ctc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(ctc[key])
			val = string(b)
		default:
			val = fmt.Sprint(ctc[key])
		}
		reply = append(reply, &CompareTemplateReply{Type: "add", Key: key, Cur: val})
	}

	// 删除的字段
	delKeys := pkg.Diff(oldKeys, curKeys)
	for _, key := range delKeys {
		val := ""
		switch otc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(otc[key])
			val = string(b)
		default:
			val = fmt.Sprint(otc[key])
		}
		reply = append(reply, &CompareTemplateReply{Type: "del", Key: key, Old: val})
	}

	// 变更的字段
	for key, val := range ctc {
		if otc[key] == nil {
			continue
		}
		if fmt.Sprint(val) == fmt.Sprint(otc[key]) {
			continue
		}

		oldVal := ""
		switch otc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(otc[key])
			oldVal = string(b)
		default:
			oldVal = fmt.Sprint(otc[key])
		}

		curVal := ""
		switch ctc[key].(type) {
		case map[string]any, []any:
			b, _ := ce.Marshal(ctc[key])
			curVal = string(b)
		default:
			curVal = fmt.Sprint(ctc[key])
		}

		reply = append(reply, &CompareTemplateReply{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return reply, nil
}
