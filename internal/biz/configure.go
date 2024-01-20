package biz

import (
	"fmt"
	"sync"

	"github.com/limes-cloud/configure/internal/biz/types"

	"github.com/go-kratos/kratos/v2/encoding"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type Configure struct {
	ktypes.BaseModel
	ServerID    uint32  `json:"server_id" gorm:"not null;comment:服务id"`
	EnvID       uint32  `json:"env_id" gorm:"not null;comment:环境id"`
	Content     string  `json:"content" gorm:"not null;type:text;comment:配置内容"`
	Version     string  `json:"version" gorm:"not null;size:32;comment:配置版本"`
	Format      string  `json:"format" gorm:"not null;size:32;comment:配置格式"`
	Description *string `json:"description" gorm:"size:128;comment:配置描述"`
	Server      *Server `json:"server" gorm:"constraint:onDelete:cascade"`
	Env         *Env    `json:"env" gorm:"constraint:onDelete:cascade"`
}

type ConfigureRepo interface {
	Get(ctx kratosx.Context, id uint32) (*Configure, error)
	GetByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*Configure, error)
	Page(ctx kratosx.Context, options *ktypes.PageOptions) ([]*Configure, uint32, error)
	Create(ctx kratosx.Context, c *Configure) (uint32, error)
	Update(ctx kratosx.Context, c *Configure) error
}

type watcher struct {
	reply types.WatcherConfigReplyFunc
	close chan string
}
type ConfigureUseCase struct {
	config *config.Config
	repo   ConfigureRepo
	mutex  sync.Mutex
	rs     map[string][]watcher
}

func NewConfigureUseCase(config *config.Config, repo ConfigureRepo) *ConfigureUseCase {
	return &ConfigureUseCase{config: config, repo: repo, rs: map[string][]watcher{}}
}

// GetByEnvAndSrc 获取指定标识的配置信息
func (u *ConfigureUseCase) GetByEnvAndSrc(ctx kratosx.Context, envId, srvId uint32) (*Configure, error) {
	configure, err := u.repo.GetByEnvAndSrv(ctx, envId, srvId)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return configure, nil
}

// Page 获取分页配置信息
func (u *ConfigureUseCase) Page(ctx kratosx.Context, page, pageSize uint32) ([]*Configure, uint32, error) {
	list, total, err := u.repo.Page(ctx, &ktypes.PageOptions{
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return nil, 0, v1.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// Add 添加配置信息
func (u *ConfigureUseCase) Add(ctx kratosx.Context, configure *Configure) (uint32, error) {
	id, err := u.repo.Create(ctx, configure)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新模配置
func (u *ConfigureUseCase) Update(ctx kratosx.Context, req *Configure) error {
	// 生成版本号
	req.Version = util.MD5([]byte(req.Content))

	// 查询往期配置
	configure, err := u.repo.GetByEnvAndSrv(ctx, req.EnvID, req.ServerID)
	if err == nil {
		if req.Version == configure.Version {
			return v1.VersionExistError()
		}
		req.ID = configure.ID

		if err := u.repo.Update(ctx, req); err != nil {
			return v1.DatabaseErrorFormat(err.Error())
		}
	} else {
		if _, err := u.repo.Create(ctx, req); err != nil {
			return v1.DatabaseErrorFormat(err.Error())
		}
	}

	u.SendWatcher(req)
	return nil
}

// Compare 对比配置
func (u *ConfigureUseCase) Compare(ctx kratosx.Context, req *types.CompareConfigureRequest) ([]*types.CompareConfigureReply, error) {
	// 获取当前配置
	configure, err := u.repo.GetByEnvAndSrv(ctx, req.EnvID, req.ServerID)

	// 解析生产中的配置
	var oldKeys []string
	otc := map[string]any{}
	if err == nil {
		oe := encoding.GetCodec(configure.Format)
		if err := oe.Unmarshal([]byte(configure.Content), &otc); err != nil {
			return nil, v1.ParseTemplateError()
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
		return nil, v1.ParseTemplateError()
	}
	for key := range ctc {
		curKeys = append(curKeys, key)
	}

	var reply []*types.CompareConfigureReply

	// 新增字段
	addKeys := util.Diff(curKeys, oldKeys)
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

	// 删除的字段
	delKeys := util.Diff(oldKeys, curKeys)
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

		reply = append(reply, &types.CompareConfigureReply{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return reply, nil
}

func (u *ConfigureUseCase) channelKey(srvId, envId uint32) string {
	return fmt.Sprintf("%v:%v", srvId, envId)
}

func (u *ConfigureUseCase) Watch(ctx kratosx.Context, req *types.WatcherConfigRequest, reply types.WatcherConfigReplyFunc) error {
	// 获取配置
	configure, err := u.repo.GetByEnvAndSrv(ctx, req.EnvID, req.ServerID)
	if err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}

	// 第一次链接先返回结果
	if err := reply(&types.WatcherConfigureReply{
		Content: configure.Content,
		Format:  configure.Format,
	}); err != nil {
		return v1.WatchConfigureErrorFormat(err.Error())
	}

	// 注册回调监听
	closer := <-u.registerWatch(req.EnvID, req.ServerID, reply)
	return v1.WatchConfigureErrorFormat(closer)
}

func (u *ConfigureUseCase) registerWatch(srvId, envId uint32, reply types.WatcherConfigReplyFunc) <-chan string {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	key := u.channelKey(srvId, envId)
	closer := make(chan string, 1)
	u.rs[key] = append(u.rs[key], watcher{
		reply: reply,
		close: closer,
	})
	return closer
}

func (u *ConfigureUseCase) SendWatcher(in *Configure) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	key := u.channelKey(in.ServerID, in.EnvID)
	for ind := 0; ind < len(u.rs[key]); ind++ {
		item := u.rs[key][ind]
		if err := item.reply(&types.WatcherConfigureReply{
			Content: in.Content,
			Format:  in.Format,
		}); err != nil {
			item.close <- err.Error()
			u.rs[key] = append(u.rs[key][:ind], u.rs[key][ind+1:]...)
			ind--
		}
	}
}
