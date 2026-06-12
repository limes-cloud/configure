package service

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/kratosx/pkg/crypto"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/domain/entity"
	"github.com/limes-cloud/configure/internal/domain/repository"
	"github.com/limes-cloud/configure/internal/pkg"
	"github.com/limes-cloud/configure/internal/types"
)

type watcher struct {
	reply types.WatcherConfigReplyFunc
	close chan string
}

type Configure struct {
	repo     repository.Configure
	env      repository.Env
	business repository.Business
	resource repository.Resource
	template repository.Template
	mutex    sync.Mutex
	rs       map[string][]*watcher
}

func NewConfigure(
	repo repository.Configure,
	env repository.Env,
	business repository.Business,
	resource repository.Resource,
	template repository.Template,
) *Configure {
	srv := &Configure{
		repo:     repo,
		env:      env,
		business: business,
		resource: resource,
		template: template,
		rs:       map[string][]*watcher{},
	}

	go repo.SubscribeConfigure(func(ctx core.Context, envId uint32, app string) error {
		return srv.SendWatcher(ctx, envId, app)
	})

	return srv
}

// GetConfigureByEnvAndApp 获取指定标识的配置信息
func (u *Configure) GetConfigureByEnvAndApp(ctx core.Context, envId uint32, app string) (*entity.Configure, error) {
	configure, err := u.repo.GetConfigureByEnvAndApp(ctx, envId, app)
	if err != nil {
		return nil, errors.GetError()
	}
	return configure, nil
}

// ListConfigure 获取分页配置信息
func (u *Configure) ListConfigure(ctx core.Context, req *types.ListConfigureRequest) ([]*entity.Configure, uint32, error) {
	list, total, err := u.repo.ListConfigure(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// UpdateConfigure 更新配置
func (u *Configure) UpdateConfigure(ctx core.Context, req *entity.Configure) error {
	format, content, err := u.RenderCurrentTemplate(ctx, req.App, req.EnvId)
	if err != nil {
		return errors.RenderTemplateError(err.Error())
	}

	req.Version = crypto.MD5([]byte(content))
	req.Format = format
	req.Content = content

	configure, err := u.repo.GetConfigureByEnvAndApp(ctx, req.EnvId, req.App)
	if err == nil {
		if req.Version == configure.Version {
			return errors.ConfigureVersionExistError()
		}
		req.Id = configure.Id
		if err := u.repo.UpdateConfigure(ctx, req); err != nil {
			return errors.DatabaseError(err.Error())
		}
	} else {
		if _, err := u.repo.CreateConfigure(ctx, req); err != nil {
			return errors.DatabaseError(err.Error())
		}
	}

	if err := u.repo.BroadcastConfigure(ctx, req.EnvId, req.App); err != nil {
		return errors.BroadcastConfigureError()
	}

	return nil
}

func (u *Configure) RenderCurrentTemplate(ctx core.Context, app string, envId uint32) (string, string, error) {
	templates, err := u.template.ListCurrentTemplate(ctx, app)
	if err != nil {
		return "", "", err
	}
	if len(templates) == 0 {
		return "", "", errors.ServerNotExistTemplateError()
	}

	bvs, err := u.business.ListBusinessValue(ctx, &types.ListBusinessValueRequest{
		App:   &app,
		EnvId: &envId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get business value error", "err", err.Error())
		return "", "", errors.DatabaseError()
	}
	rvs, err := u.resource.ListResourceValue(ctx, &types.ListResourceValueRequest{
		App:   &app,
		EnvId: &envId,
	})
	if err != nil {
		ctx.Logger().Errorw("msg", "get resource value error", "err", err.Error())
		return "", "", errors.DatabaseError()
	}

	tpr := pkg.NewTemplateParser()
	format := templates[0].Format
	merged := map[string]any{}
	for _, template := range templates {
		content, err := tpr.RenderTemplate(bvs, rvs, template.Format, template.Content)
		if err != nil {
			return "", "", errors.RenderTemplateError(err.Error())
		}
		codec := encoding.GetCodec(template.Format)
		if codec == nil {
			return "", "", errors.RenderTemplateError("unsupported template format")
		}
		current := map[string]any{}
		if err := codec.Unmarshal([]byte(content), &current); err != nil {
			return "", "", errors.RenderTemplateError(err.Error())
		}
		mergeConfigureMap(merged, current)
		if strings.EqualFold(template.Namespace, "conf") {
			format = template.Format
		}
	}

	codec := encoding.GetCodec(format)
	if codec == nil {
		return "", "", errors.RenderTemplateError("unsupported template format")
	}
	content, err := codec.Marshal(merged)
	if err != nil {
		return "", "", errors.RenderTemplateError(err.Error())
	}

	return format, string(content), nil
}

func mergeConfigureMap(dst map[string]any, src map[string]any) {
	for key, val := range src {
		srcMap, srcIsMap := val.(map[string]any)
		dstMap, dstIsMap := dst[key].(map[string]any)
		if srcIsMap && dstIsMap {
			mergeConfigureMap(dstMap, srcMap)
			continue
		}
		dst[key] = val
	}
}

// CompareConfigure 对比配置
func (u *Configure) CompareConfigure(ctx core.Context, req *types.CompareConfigureRequest) ([]*types.CompareConfigureReply, error) {
	configure, err := u.repo.GetConfigureByEnvAndApp(ctx, req.EnvId, req.App)

	var oldKeys []string
	otc := map[string]any{}
	if err == nil {
		oe := encoding.GetCodec(configure.Format)
		if err := oe.Unmarshal([]byte(configure.Content), &otc); err != nil {
			return nil, errors.RenderTemplateError()
		}
		for key := range otc {
			oldKeys = append(oldKeys, key)
		}
	}

	var curKeys []string
	format, content, err := u.RenderCurrentTemplate(ctx, req.App, req.EnvId)
	if err != nil {
		return nil, errors.RenderTemplateError(err.Error())
	}
	ctc := map[string]any{}
	ce := encoding.GetCodec(format)
	if err := ce.Unmarshal([]byte(content), &ctc); err != nil {
		return nil, errors.RenderTemplateError()
	}
	for key := range ctc {
		curKeys = append(curKeys, key)
	}

	var reply []*types.CompareConfigureReply

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
		reply = append(reply, &types.CompareConfigureReply{Type: "add", Key: key, Cur: val})
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
		reply = append(reply, &types.CompareConfigureReply{Type: "del", Key: key, Old: val})
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

		reply = append(reply, &types.CompareConfigureReply{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return reply, nil
}

func (u *Configure) channelKey(envId uint32, app string) string {
	return fmt.Sprintf("%v:%v", envId, app)
}

func (u *Configure) Watch(ctx core.Context, in *types.WatcherConfigRequest, reply types.WatcherConfigReplyFunc) error {
	env, err := u.env.GetEnvByToken(ctx, in.Token)
	if err != nil {
		return errors.TokenAuthError()
	}

	configure, err := u.repo.GetConfigureByEnvAndApp(ctx, env.Id, in.App)
	if err != nil {
		return errors.DatabaseError()
	}

	if err := reply(&types.WatcherConfigureReply{
		Content: configure.Content,
		Format:  configure.Format,
	}); err != nil {
		return errors.WatchConfigureError(err.Error())
	}

	closer := <-u.registerWatch(ctx, env.Id, in.App, reply)
	return errors.WatchConfigureError(closer)
}

func (u *Configure) removeNilWatcher(arr []*watcher) []*watcher {
	var result []*watcher
	for _, ptr := range arr {
		if ptr != nil {
			result = append(result, ptr)
		}
	}
	return result
}

func (u *Configure) registerWatch(ctx context.Context, envId uint32, app string, reply types.WatcherConfigReplyFunc) <-chan string {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	key := u.channelKey(envId, app)
	closer := make(chan string, 1)
	u.rs[key] = append(u.rs[key], &watcher{
		reply: reply,
		close: closer,
	})
	index := len(u.rs[key]) - 1
	go func() {
		<-ctx.Done()
		closer <- ctx.Err().Error()

		u.mutex.Lock()
		defer u.mutex.Unlock()

		u.rs[key][index] = nil
		u.rs[key] = u.removeNilWatcher(u.rs[key])
	}()

	return closer
}

func (u *Configure) SendWatcher(ctx core.Context, envId uint32, app string) error {
	cfg, err := u.repo.GetConfigureByEnvAndApp(ctx, envId, app)
	if err != nil {
		return errors.DatabaseError(err.Error())
	}

	u.mutex.Lock()
	defer u.mutex.Unlock()

	key := u.channelKey(envId, app)
	for ind := 0; ind < len(u.rs[key]); ind++ {
		item := u.rs[key][ind]
		if err := item.reply(&types.WatcherConfigureReply{
			Content: cfg.Content,
			Format:  cfg.Format,
		}); err != nil {
			item.close <- err.Error()
			u.rs[key] = append(u.rs[key][:ind], u.rs[key][ind+1:]...)
			ind--
		}
	}
	return nil
}
