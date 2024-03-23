package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/configure/api/configure/v1"
	"github.com/limes-cloud/configure/api/errors"
	biz "github.com/limes-cloud/configure/internal/biz/configure"
	envbiz "github.com/limes-cloud/configure/internal/biz/env"
	srvbiz "github.com/limes-cloud/configure/internal/biz/server"
	tpbiz "github.com/limes-cloud/configure/internal/biz/template"
	"github.com/limes-cloud/configure/internal/config"
	data "github.com/limes-cloud/configure/internal/data/configure"
)

type ConfigureService struct {
	v1.UnimplementedServiceServer
	uc    *biz.UseCase
	tpuc  *tpbiz.UseCase
	envuc *envbiz.UseCase
	srvuc *srvbiz.UseCase
}

func NewConfigureService(conf *config.Config, tpuc *tpbiz.UseCase, envuc *envbiz.UseCase, srvuc *srvbiz.UseCase) *ConfigureService {
	return &ConfigureService{
		uc:    biz.NewUseCase(conf, data.NewRepo()),
		tpuc:  tpuc,
		envuc: envuc,
		srvuc: srvuc,
	}
}

func (s *ConfigureService) CompareConfigure(ctx context.Context, in *v1.CompareConfigureRequest) (*v1.CompareConfigureReply, error) {
	kCtx := kratosx.MustContext(ctx)
	res, err := s.tpuc.ParseTemplate(kCtx, &tpbiz.ParseTemplateRequest{EnvId: in.EnvId, ServerId: in.ServerId})
	if err != nil {
		return nil, err
	}

	list, err := s.uc.CompareConfigure(kratosx.MustContext(ctx), &biz.CompareConfigureRequest{
		EnvId:    in.EnvId,
		ServerId: in.ServerId,
		Content:  res.Content,
		Format:   res.Format,
	})
	if err != nil {
		return nil, err
	}
	reply := v1.CompareConfigureReply{}
	if err := copier.Copy(&reply.List, list); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *ConfigureService) GetConfigure(ctx context.Context, in *v1.GetConfigureRequest) (*v1.GetConfigureReply, error) {
	configure, err := s.uc.GetConfigureByEnvAndSrv(kratosx.MustContext(ctx), in.EnvId, in.ServerId)
	if err != nil {
		return nil, err
	}
	reply := v1.GetConfigureReply{}
	if err := copier.Copy(&reply, configure); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *ConfigureService) UpdateConfigure(ctx context.Context, in *v1.UpdateConfigureRequest) (*emptypb.Empty, error) {
	kCtx := kratosx.MustContext(ctx)
	res, err := s.tpuc.ParseTemplate(kCtx, &tpbiz.ParseTemplateRequest{EnvId: in.EnvId, ServerId: in.ServerId})
	if err != nil {
		return nil, err
	}
	return nil, s.uc.UpdateConfigure(kratosx.MustContext(ctx), &biz.Configure{
		ServerId:    in.ServerId,
		EnvId:       in.EnvId,
		Description: &in.Description,
		Content:     res.Content,
		Format:      res.Format,
	})
}

func (s *ConfigureService) WatchConfigure(in *v1.WatchConfigureRequest, reply v1.Service_WatchConfigureServer) error {
	kCtx := kratosx.MustContext(reply.Context())

	env, err := s.envuc.GetEnvByToken(kCtx, in.Token)
	if err != nil {
		return errors.TokenAuthError()
	}
	server, err := s.srvuc.GetServerByKeyword(kCtx, in.Server)
	if err != nil {
		return errors.TokenAuthError()
	}

	return s.uc.Watch(kCtx, &biz.WatcherConfigRequest{
		ServerId: server.ID,
		EnvId:    env.ID,
	}, func(data *biz.WatcherConfigureReply) error {
		return reply.Send(&v1.WatchConfigureReply{
			Format:  data.Format,
			Content: data.Content,
		})
	})
}
