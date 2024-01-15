package logic

import (
	"fmt"
	"sync"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
)

type watcher struct {
	reply v1.Service_WatchConfigureServer
	close chan string
}

type Configure struct {
	conf  *config.Config
	rs    map[string][]watcher
	mutex sync.Mutex
}

func NewConfigure(conf *config.Config) *Configure {
	return &Configure{
		conf: conf,
		rs:   map[string][]watcher{},
	}
}

// Get 查询指定模板
func (t *Configure) Get(ctx kratosx.Context, in *v1.GetConfigureRequest) (*v1.GetConfigureReply, error) {
	env := model.Environment{}
	if err := env.OneByKeyword(ctx, in.EnvKeyword); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}
	configure := model.Configure{}
	if err := configure.OneBySrvAndEnv(ctx, in.ServerId, env.ID); err != nil {
		return nil, v1.DatabaseError()
	}
	reply := v1.GetConfigureReply{}
	if util.Transform(configure, &reply) != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

// Update 更新模配置
func (t *Configure) Update(ctx kratosx.Context, in *v1.UpdateConfigureRequest) (*emptypb.Empty, error) {
	configure := model.Configure{}
	if err := util.Transform(in, &configure); err != nil {
		return nil, v1.TransformError()
	}

	// 获取环境
	env := model.Environment{}
	if err := env.OneByKeyword(ctx, in.EnvKeyword); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取当前正在使用的模板
	tp := model.Template{}
	if err := tp.Current(ctx, in.ServerId); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 进行模板解析
	template := NewTemplate(t.conf)
	reply, err := template.Parse(ctx, &v1.ParseTemplateRequest{ServerId: in.ServerId, EnvKeyword: env.Keyword})
	if err != nil {
		return nil, err
	}

	// 获取解析后的值
	configure.Content = reply.Content
	configure.Format = reply.Format
	configure.Version = util.MD5([]byte(reply.Content))
	configure.EnvironmentID = env.ID

	// 查询往期配置
	old := &model.Configure{}
	if err := old.OneBySrvAndEnv(ctx, in.ServerId, env.ID); err == nil {
		// 校验当前版本是否和之前版本一致
		if old.Version == configure.Version {
			return nil, v1.VersionExistError()
		}
		configure.ID = old.ID
		if err := configure.Update(ctx); err != nil {
			return nil, v1.DatabaseErrorFormat(err.Error())
		}
	} else {
		if err := configure.Create(ctx); err != nil {
			return nil, v1.DatabaseErrorFormat(err.Error())
		}
	}

	t.SendWatcher(configure)
	return nil, nil
}

// Compare 对比配置
func (t *Configure) Compare(ctx kratosx.Context, in *v1.CompareConfigureRequest) (*v1.CompareConfigureReply, error) {
	// 获取环境
	env := model.Environment{}
	if err := env.OneByKeyword(ctx, in.EnvKeyword); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取当前正在使用的模板
	tp := model.Template{}
	if err := tp.Current(ctx, in.ServerId); err != nil {
		return nil, v1.DatabaseErrorFormat(err.Error())
	}

	// 获取生产中的配置
	configure := model.Configure{}
	if err := configure.OneBySrvAndEnv(ctx, in.ServerId, env.ID); err != nil {
		return nil, err
	}

	// 解析当前配置
	template := NewTemplate(t.conf)
	parse, err := template.Parse(ctx, &v1.ParseTemplateRequest{ServerId: in.ServerId, EnvKeyword: in.EnvKeyword})
	if err != nil {
		return nil, err
	}

	// 对比配置
	// 解析生产中的配置
	var oldKeys []string
	otc := map[string]any{}

	oe := encoding.GetCodec(configure.Format)
	if err := oe.Unmarshal([]byte(configure.Content), &otc); err != nil {
		return nil, v1.ParseTemplateError()
	}
	for key := range otc {
		oldKeys = append(oldKeys, key)
	}

	// 解析上传的模板
	var curKeys []string
	ctc := map[string]any{}

	ce := encoding.GetCodec(parse.Format)
	if err := ce.Unmarshal([]byte(parse.Content), &ctc); err != nil {
		return nil, v1.ParseTemplateError()
	}
	for key := range ctc {
		curKeys = append(curKeys, key)
	}

	reply := v1.CompareConfigureReply{
		List: make([]*v1.CompareConfigureInfo, 0),
	}

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
		reply.List = append(reply.List, &v1.CompareConfigureInfo{Type: "add", Key: key, Cur: val})
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
		reply.List = append(reply.List, &v1.CompareConfigureInfo{Type: "del", Key: key, Old: val})
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

		reply.List = append(reply.List, &v1.CompareConfigureInfo{Type: "update", Key: key, Old: oldVal, Cur: curVal})
	}

	return &reply, nil
}

func (t *Configure) channelKey(srvId, envId uint32) string {
	return fmt.Sprintf("%v:%v", srvId, envId)
}

func (t *Configure) Watch(ctx kratosx.Context, req *v1.WatchConfigureRequest, reply v1.Service_WatchConfigureServer) error {
	// 通过token查找环境是否存在
	env := model.Environment{}
	if err := env.OneByToken(ctx, req.Token); err != nil {
		return v1.TokenAuthError()
	}

	srv := model.Server{}
	if err := srv.OneByKeyword(ctx, req.Server); err != nil {
		return v1.ServerNotFound()
	}

	configure := model.Configure{}
	if err := configure.OneBySrvAndEnv(ctx, srv.ID, env.ID); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}

	// 第一次链接先返回结果
	if err := reply.Send(&v1.WatchConfigureReply{
		Content: configure.Content,
		Format:  configure.Format,
	}); err != nil {
		return v1.WatchConfigureErrorFormat(err.Error())
	}

	// 注册回调监听
	closer := <-t.registerWatch(srv.ID, env.ID, reply)
	return v1.WatchConfigureErrorFormat(closer)
}

func (t *Configure) registerWatch(srvId, envId uint32, reply v1.Service_WatchConfigureServer) <-chan string {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	key := t.channelKey(srvId, envId)
	closer := make(chan string, 1)
	t.rs[key] = append(t.rs[key], watcher{
		reply: reply,
		close: closer,
	})

	return closer
}

func (t *Configure) SendWatcher(in model.Configure) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	key := t.channelKey(in.ServerID, in.EnvironmentID)
	for ind := 0; ind < len(t.rs[key]); ind++ {
		item := t.rs[key][ind]

		if err := item.reply.Send(&v1.WatchConfigureReply{
			Content: in.Content,
			Format:  in.Format,
		}); err != nil {
			item.close <- err.Error()
			t.rs[key] = append(t.rs[key][:ind], t.rs[key][ind+1:]...)
			ind--
		}
	}
}
