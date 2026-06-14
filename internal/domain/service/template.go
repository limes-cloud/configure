package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/kratosx/library/db/gormtranserror"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"gorm.io/gorm"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/configure/internal/pkg"
	"github.com/limes-cloud/configure/internal/types"
)

const defaultTemplateNamespace = "conf"

func normalizeTemplateNamespace(namespace string) string {
	namespace = strings.TrimSpace(namespace)
	if namespace == "" {
		return defaultTemplateNamespace
	}
	return namespace
}

type Template struct {
	repo      repository.Template
	namespace repository.TemplateNamespace
	business  repository.Business
	resource  repository.Resource
}

func NewTemplate(
	repo repository.Template,
	namespace repository.TemplateNamespace,
	business repository.Business,
	resource repository.Resource,
) *Template {
	return &Template{
		repo:      repo,
		namespace: namespace,
		business:  business,
		resource:  resource,
	}
}

// CurrentTemplate 获取当前版本的配置信息
func (t *Template) CurrentTemplate(ctx core.Context, app string, namespace string) (*entity.Template, error) {
	template, err := t.repo.CurrentTemplate(ctx, app, normalizeTemplateNamespace(namespace))
	if err != nil {
		if gormtranserror.Is(gorm.ErrRecordNotFound, err) {
			return nil, errors.ServerNotExistTemplateError()
		}
		return nil, errors.GetError(err.Error())
	}
	return template, nil
}

// ListCurrentTemplate 获取应用下当前版本的配置信息
func (t *Template) ListCurrentTemplate(ctx core.Context, app string) ([]*entity.Template, error) {
	templates, err := t.repo.ListCurrentTemplate(ctx, app)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return templates, nil
}

// GetTemplate 获取指定模板信息
func (t *Template) GetTemplate(ctx core.Context, id uint32) (*entity.Template, error) {
	template, err := t.repo.GetTemplate(ctx, id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return template, nil
}

// ListTemplate 获取分页模板信息
func (t *Template) ListTemplate(ctx core.Context, req *types.ListTemplateRequest) ([]*entity.Template, uint32, error) {
	if req.Namespace != nil {
		namespace := normalizeTemplateNamespace(*req.Namespace)
		req.Namespace = &namespace
	}
	list, total, err := t.repo.ListTemplate(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateTemplate 添加模板信息
func (t *Template) CreateTemplate(ctx core.Context, template *entity.Template) (uint32, error) {
	template.Namespace = normalizeTemplateNamespace(template.Namespace)
	if _, err := t.namespace.GetTemplateNamespaceByAppAndNamespace(ctx, template.App, template.Namespace); err != nil {
		return 0, errors.GetError(err.Error())
	}
	otc := map[string]any{}
	oe := encoding.GetCodec(template.Format)
	if err := oe.Unmarshal([]byte(template.Content), &otc); err != nil {
		return 0, errors.RenderTemplateError(err.Error())
	}

	template.Version = crypto.MD5ToUpper([]byte(template.Content))

	if _, err := t.repo.GetTemplateByVersion(ctx, template.Version); err == nil {
		return 0, errors.TemplateVersionExistError()
	}

	bsKeys, _ := t.business.AllBusinessField(ctx, template.App)
	rsKeys, _ := t.resource.AllResourceField(ctx, template.App)
	keys := append(bsKeys, rsKeys...)
	tp := pkg.NewTemplateParser()
	if err := tp.CheckTemplate(keys, template.Content); err != nil {
		return 0, errors.RenderTemplateError(err.Error())
	}

	current, err := t.repo.CurrentTemplate(ctx, template.App, template.Namespace)
	if err == nil {
		compare, err := t.CompareTemplate(ctx, &types.CompareTemplateRequest{
			Id:      current.Id,
			Format:  template.Format,
			Content: template.Content,
		})
		if err != nil {
			return 0, errors.RenderTemplateError(err.Error())
		}
		compareByte, _ := json.Marshal(compare)
		template.Compare = string(compareByte)
	}

	id, err := t.repo.CreateTemplate(ctx, template)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}

	_ = t.repo.UseTemplate(ctx, template.App, template.Namespace, id)
	return id, nil
}

// PreviewTemplate 预览配置
func (t *Template) PreviewTemplate(ctx core.Context, req *types.PreviewTemplateRequest) (*types.PreviewTemplateReply, error) {
	bvs, err := t.business.ListBusinessValue(ctx, &types.ListBusinessValueRequest{
		App:   &req.App,
		EnvId: &req.EnvId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get business value error", "err", err.Error())
		return nil, errors.DatabaseError()
	}
	rvs, err := t.resource.ListResourceValue(ctx, &types.ListResourceValueRequest{
		App:   &req.App,
		EnvId: &req.EnvId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get resource value error", "err", err.Error())
		return nil, errors.DatabaseError()
	}
	tp := pkg.NewTemplateParser()
	content, err := tp.RenderTemplate(bvs, rvs, req.Format, req.Content)
	if err != nil {
		return nil, err
	}

	return &types.PreviewTemplateReply{
		Content: content,
		Format:  req.Format,
	}, err
}

// PreviewCurrentTemplate 预览当前配置
func (t *Template) PreviewCurrentTemplate(ctx core.Context, req *types.PreviewCurrentTemplateRequest) (*types.PreviewTemplateReply, error) {
	namespace := normalizeTemplateNamespace(req.Namespace)
	tp, err := t.repo.CurrentTemplate(ctx, req.App, namespace)
	if err != nil {
		return nil, errors.DatabaseError(err.Error())
	}

	reply := &types.PreviewTemplateReply{
		Format: tp.Format,
	}

	bvs, err := t.business.ListBusinessValue(ctx, &types.ListBusinessValueRequest{
		App:   &req.App,
		EnvId: &req.EnvId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get business value error", "err", err.Error())
		return nil, errors.DatabaseError()
	}
	rvs, err := t.resource.ListResourceValue(ctx, &types.ListResourceValueRequest{
		App:   &req.App,
		EnvId: &req.EnvId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get resource value error", "err", err.Error())
		return nil, errors.DatabaseError()
	}
	tpr := pkg.NewTemplateParser()
	tp.Content, err = tpr.RenderTemplate(bvs, rvs, tp.Format, tp.Content)
	if err != nil {
		return nil, err
	}

	return reply, err
}

// SwitchTemplate 切换指定版本信息
func (t *Template) SwitchTemplate(ctx core.Context, app string, namespace string, tpId uint32) error {
	if err := t.repo.UseTemplate(ctx, app, normalizeTemplateNamespace(namespace), tpId); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// UpdateTemplate 更新模板信息
func (t *Template) UpdateTemplate(ctx core.Context, req *entity.Template) error {
	if err := t.repo.UpdateTemplate(ctx, req); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// DeleteUpdateTemplate 删除模板信息
func (t *Template) DeleteUpdateTemplate(ctx core.Context, id uint32) error {
	if err := t.repo.DeleteTemplate(ctx, id); err != nil {
		return errors.DatabaseError(err.Error())
	}
	return nil
}

// CompareTemplate 对比变更细节
func (t *Template) CompareTemplate(ctx core.Context, req *types.CompareTemplateRequest) ([]*types.CompareTemplateReply, error) {
	template, err := t.repo.GetTemplate(ctx, req.Id)
	if err != nil {
		return nil, errors.GetError(err.Error())
	}

	var oldKeys []string
	otc := map[string]any{}
	oe := encoding.GetCodec(template.Format)
	if err := oe.Unmarshal([]byte(template.Content), &otc); err != nil {
		return nil, errors.RenderTemplateError()
	}
	for key := range otc {
		oldKeys = append(oldKeys, key)
	}

	var curKeys []string
	ctc := map[string]any{}
	ce := encoding.GetCodec(req.Format)
	if err := ce.Unmarshal([]byte(req.Content), &ctc); err != nil {
		return nil, errors.RenderTemplateError()
	}
	for key := range ctc {
		curKeys = append(curKeys, key)
	}

	var reply []*types.CompareTemplateReply

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
		reply = append(reply, &types.CompareTemplateReply{Type: "add", Key: key, Cur: val})
	}

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
		reply = append(reply, &types.CompareTemplateReply{Type: "del", Key: key, Old: val})
	}

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

		reply = append(reply, &types.CompareTemplateReply{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return reply, nil
}
