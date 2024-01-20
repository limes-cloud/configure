package service

import (
	"context"

	"github.com/limes-cloud/configure/internal/biz"

	"github.com/limes-cloud/configure/internal/biz/types"

	"github.com/jinzhu/copier"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CompareConfigure(ctx context.Context, in *v1.CompareConfigureRequest) (*v1.CompareConfigureReply, error) {
	kCtx := kratosx.MustContext(ctx)
	parseReply, err := s.TemplateUseCase.Parse(kCtx, &types.ParseRequest{EnvId: in.EnvId, ServerId: in.ServerId})
	if err != nil {
		return nil, err
	}

	list, err := s.ConfigureUseCase.Compare(kCtx, &types.CompareConfigureRequest{
		EnvID:    in.EnvId,
		ServerID: in.ServerId,
		Format:   parseReply.Format,
		Content:  parseReply.Content,
	})
	reply := v1.CompareConfigureReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) GetConfigure(ctx context.Context, in *v1.GetConfigureRequest) (*v1.GetConfigureReply, error) {
	configure, err := s.ConfigureUseCase.GetByEnvAndSrc(kratosx.MustContext(ctx), in.EnvId, in.ServerId)
	if err != nil {
		return nil, err
	}
	reply := v1.GetConfigureReply{}
	if err := copier.Copy(&reply, configure); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) UpdateConfigure(ctx context.Context, in *v1.UpdateConfigureRequest) (*emptypb.Empty, error) {
	kCtx := kratosx.MustContext(ctx)
	parseReply, err := s.TemplateUseCase.Parse(kCtx, &types.ParseRequest{EnvId: in.EnvId, ServerId: in.ServerId})
	if err != nil {
		return nil, err
	}

	return nil, s.ConfigureUseCase.Update(kratosx.MustContext(ctx), &biz.Configure{
		ServerID:    in.ServerId,
		EnvID:       in.EnvId,
		Description: &in.Description,
		Format:      parseReply.Format,
		Content:     parseReply.Content,
	})
}

func (s *Service) WatchConfigure(in *v1.WatchConfigureRequest, reply v1.Service_WatchConfigureServer) error {
	kCtx := kratosx.MustContext(reply.Context())

	env, err := s.EnvUseCase.GetByToken(kCtx, in.Token)
	if err != nil {
		return v1.TokenAuthError()
	}
	server, err := s.ServerUseCase.GetByKeyword(kCtx, in.Server)
	if err != nil {
		return v1.TokenAuthError()
	}

	return s.ConfigureUseCase.Watch(kratosx.MustContext(reply.Context()), &types.WatcherConfigRequest{
		EnvID:    env.ID,
		ServerID: server.ID,
	}, func(data *types.WatcherConfigureReply) error {
		return reply.Send(&v1.WatchConfigureReply{
			Format:  data.Format,
			Content: data.Content,
		})
	})
}
