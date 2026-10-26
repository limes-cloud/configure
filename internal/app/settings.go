package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/configure/internal/core"

	pb "github.com/limes-cloud/configure/api/settings"
	"github.com/limes-cloud/configure/internal/domain/service"
	"github.com/limes-cloud/configure/internal/infra/dbs"
	"github.com/limes-cloud/configure/internal/types"
)

type Settings struct {
	pb.UnimplementedSettingsServer
	srv *service.Settings
}

func NewSettings() *Settings {
	return &Settings{
		srv: service.NewSettings(
			dbs.NewBusiness(),
			dbs.NewTemplate(),
			getConfigureSrv(), // 与配置模块共享单例，避免重复订阅配置变更广播
		),
	}
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		srv := NewSettings()
		pb.RegisterSettingsHTTPServer(hs, srv)
		pb.RegisterSettingsServer(gs, srv)
	})
}

// GetSettings 获取系统配置（当前模板引用的业务变量 + 指定环境的值）
func (s *Settings) GetSettings(c context.Context, req *pb.GetSettingsRequest) (*pb.GetSettingsResponse, error) {
	ctx := core.MustContext(c)
	reply, err := s.srv.GetSettings(ctx, resolveApp(ctx, req.GetApp()), req.EnvId)
	if err != nil {
		return nil, err
	}

	resp := &pb.GetSettingsResponse{App: reply.App, EnvId: reply.EnvId, Format: reply.Format}
	for _, item := range reply.Items {
		resp.Items = append(resp.Items, &pb.SettingsItem{
			Keyword:     item.Keyword,
			Type:        item.Type,
			Description: item.Description,
			Extra:       item.Extra,
			Value:       item.Value,
		})
	}
	return resp, nil
}

// UpdateSettings 保存系统配置（类型校验 + 落库 + 重新渲染 + 广播，保存即生效）
func (s *Settings) UpdateSettings(c context.Context, req *pb.UpdateSettingsRequest) (*pb.UpdateSettingsResponse, error) {
	items := make([]*types.UpdateSettingsItem, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, &types.UpdateSettingsItem{
			Keyword: item.Keyword,
			Value:   item.Value,
		})
	}
	if err := s.srv.UpdateSettings(core.MustContext(c), resolveApp(core.MustContext(c), req.GetApp()), req.EnvId, items); err != nil {
		return nil, err
	}
	return &pb.UpdateSettingsResponse{}, nil
}