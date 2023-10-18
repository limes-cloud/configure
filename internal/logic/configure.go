package logic

import (
	"fmt"
	"sync"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/md"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
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
func (t *Configure) Get(ctx kratos.Context, in *v1.GetConfigureRequest) (*v1.GetConfigureReply, error) {
	configure := model.Configure{}
	if err := configure.OneBySrvAndEnv(ctx, in.ServerId, in.EnvironmentId); err != nil {
		return nil, v1.ErrorDatabase()
	}
	reply := v1.GetConfigureReply{}
	if util.Transform(configure, &reply) != nil {
		return nil, v1.ErrorTransform()
	}
	return &reply, nil
}

// Update 更新模配置
func (t *Configure) Update(ctx kratos.Context, in *v1.UpdateConfigureRequest) (*emptypb.Empty, error) {
	configure := model.Configure{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if err := util.Transform(in, &configure); err != nil {
		return nil, v1.ErrorTransform()
	}

	// 获取当前正在使用的模板
	tp := model.Template{}
	if err := tp.Current(ctx, in.ServerId); err != nil {
		return nil, v1.ErrorDatabaseFormat(err.Error())
	}

	// 进行模板解析
	template := NewTemplate(t.conf)
	reply, err := template.Parse(ctx, &v1.ParseTemplateRequest{ServerId: in.ServerId, EnvironmentId: in.EnvironmentId})
	if err != nil {
		return nil, err
	}

	// 获取解析后的值
	configure.Content = reply.Content
	configure.Format = reply.Format
	configure.Version = util.MD5([]byte(reply.Content))

	// 查询往期配置
	old := &model.Configure{}
	if old.OneBySrvAndEnv(ctx, in.ServerId, in.EnvironmentId) == nil {
		//校验当前版本是否和之前版本一致
		if old.Version == configure.Version {
			return nil, v1.ErrorVersionExist()
		}
		configure.ID = old.ID
		if err := configure.Update(ctx); err != nil {
			return nil, v1.ErrorDatabaseFormat(err.Error())
		}
	} else {
		if err := configure.Create(ctx); err != nil {
			return nil, v1.ErrorDatabaseFormat(err.Error())
		}
	}

	t.SendWatcher(configure)
	return nil, nil
}

func (t *Configure) channelKey(srvId, envId int64) string {
	return fmt.Sprintf("%v:%v", srvId, envId)
}

func (t *Configure) Watch(ctx kratos.Context, req *v1.WatchConfigureRequest, reply v1.Service_WatchConfigureServer) error {
	// 通过token查找环境是否存在
	env := model.Environment{}
	if err := env.OneByToken(ctx, req.Token); err != nil {
		return v1.ErrorDatabaseFormat(err.Error())
	}

	srv := model.Server{}
	if err := srv.OneByKeyword(ctx, req.Server); err != nil {
		return v1.ErrorDatabaseFormat(err.Error())
	}

	configure := model.Configure{}
	if err := configure.OneBySrvAndEnv(ctx, srv.ID, env.ID); err != nil {
		return v1.ErrorDatabaseFormat(err.Error())
	}

	// 第一次链接先返回结果
	if err := reply.Send(&v1.WatchConfigureReply{
		Content: configure.Content,
		Format:  configure.Format,
	}); err != nil {
		return v1.ErrorWatchConfigureFormat(err.Error())
	}

	// 注册回调监听
	closer := <-t.registerWatch(srv.ID, env.ID, reply)
	return v1.ErrorWatchConfigureFormat(closer)
}

func (t *Configure) registerWatch(srvId, envId int64, reply v1.Service_WatchConfigureServer) <-chan string {
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
