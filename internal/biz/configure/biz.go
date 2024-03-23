package configure

import (
	"fmt"
	"sync"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/util"
	ktypes "github.com/limes-cloud/kratosx/types"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/config"
	"github.com/limes-cloud/configure/internal/pkg"
)

type watcher struct {
	reply WatcherConfigReplyFunc
	close chan string
}

type UseCase struct {
	config *config.Config
	repo   Repo
	mutex  sync.Mutex
	rs     map[string][]watcher
}

func NewUseCase(config *config.Config, repo Repo) *UseCase {
	return &UseCase{
		config: config,
		repo:   repo,
		rs:     map[string][]watcher{},
	}
}

// GetConfigureByEnvAndSrv 获取指定标识的配置信息
func (u *UseCase) GetConfigureByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*Configure, error) {
	configure, err := u.repo.GetConfigureByEnvAndSrv(ctx, envId, srvId)
	if err != nil {
		return nil, errors.NotRecordError()
	}
	return configure, nil
}

// PageConfigure 获取分页配置信息
func (u *UseCase) PageConfigure(ctx kratosx.Context, page, pageSize uint32) ([]*Configure, uint32, error) {
	list, total, err := u.repo.PageConfigure(ctx, &ktypes.PageOptions{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, 0, errors.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// UpdateConfigure 更新模配置
func (u *UseCase) UpdateConfigure(ctx kratosx.Context, req *Configure) error {
	// 生成版本号
	req.Version = util.MD5([]byte(req.Content))

	// 查询往期配置
	configure, err := u.repo.GetConfigureByEnvAndSrv(ctx, req.EnvId, req.ServerId)
	if err == nil {
		if req.Version == configure.Version {
			return errors.VersionExistError()
		}
		req.ID = configure.ID

		if err := u.repo.UpdateConfigure(ctx, req); err != nil {
			return errors.DatabaseErrorFormat(err.Error())
		}
	} else {
		if _, err := u.repo.AddConfigure(ctx, req); err != nil {
			return errors.DatabaseErrorFormat(err.Error())
		}
	}

	u.SendWatcher(req)
	return nil
}

// CompareConfigure 对比配置
func (u *UseCase) CompareConfigure(ctx kratosx.Context, req *CompareConfigureRequest) ([]*CompareConfigureReply, error) {
	// 获取当前配置
	configure, err := u.repo.GetConfigureByEnvAndSrv(ctx, req.EnvId, req.ServerId)

	// 解析生产中的配置
	var oldKeys []string
	otc := map[string]any{}
	if err == nil {
		oe := encoding.GetCodec(configure.Format)
		if err := oe.Unmarshal([]byte(configure.Content), &otc); err != nil {
			return nil, errors.ParseTemplateError()
		}
		for key := range otc {
			oldKeys = append(oldKeys, key)
		}
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

	var reply []*CompareConfigureReply

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
		reply = append(reply, &CompareConfigureReply{Type: "add", Key: key, Cur: val})
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
		reply = append(reply, &CompareConfigureReply{Type: "del", Key: key, Old: val})
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

		reply = append(reply, &CompareConfigureReply{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return reply, nil
}

func (u *UseCase) channelKey(envId, srvId uint32) string {
	return fmt.Sprintf("%v:%v", envId, srvId)
}

func (u *UseCase) Watch(ctx kratosx.Context, in *WatcherConfigRequest, reply WatcherConfigReplyFunc) error {
	// 获取配置
	configure, err := u.repo.GetConfigureByEnvAndSrv(ctx, in.EnvId, in.ServerId)
	if err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}

	// 第一次链接先返回结果
	if err := reply(&WatcherConfigureReply{
		Content: configure.Content,
		Format:  configure.Format,
	}); err != nil {
		return errors.WatchConfigureErrorFormat(err.Error())
	}

	// 注册回调监听
	closer := <-u.registerWatch(in.EnvId, in.ServerId, reply)
	return errors.WatchConfigureErrorFormat(closer)
}

func (u *UseCase) registerWatch(envId, srvId uint32, reply WatcherConfigReplyFunc) <-chan string {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	key := u.channelKey(envId, srvId)
	closer := make(chan string, 1)
	u.rs[key] = append(u.rs[key], watcher{
		reply: reply,
		close: closer,
	})
	return closer
}

func (u *UseCase) SendWatcher(in *Configure) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	key := u.channelKey(in.EnvId, in.ServerId)
	for ind := 0; ind < len(u.rs[key]); ind++ {
		item := u.rs[key][ind]
		if err := item.reply(&WatcherConfigureReply{
			Content: in.Content,
			Format:  in.Format,
		}); err != nil {
			item.close <- err.Error()
			u.rs[key] = append(u.rs[key][:ind], u.rs[key][ind+1:]...)
			ind--
		}
	}
}
