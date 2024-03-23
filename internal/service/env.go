package service

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/limes-cloud/configure/api/env/v1"
	"github.com/limes-cloud/configure/api/errors"
	biz "github.com/limes-cloud/configure/internal/biz/env"
	"github.com/limes-cloud/configure/internal/config"
	data "github.com/limes-cloud/configure/internal/data/env"
)

type EnvService struct {
	v1.UnimplementedServiceServer
	uc *biz.UseCase
}

func NewEnvService(conf *config.Config) *EnvService {
	return &EnvService{
		uc: biz.NewUseCase(conf, data.NewRepo()),
	}
}

func (s *EnvService) AllEnv(ctx context.Context, in *emptypb.Empty) (*v1.AllEnvReply, error) {
	envs, err := s.uc.AllEnv(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	reply := v1.AllEnvReply{}
	if err := copier.Copy(&reply.List, envs); err != nil {
		return nil, errors.TransformError()
	}
	return &reply, nil
}

func (s *EnvService) AddEnv(ctx context.Context, in *v1.AddEnvRequest) (*v1.AddEnvReply, error) {
	env := &biz.Env{}
	if err := copier.Copy(env, in); err != nil {
		return nil, errors.TransformError()
	}
	id, err := s.uc.AddEnv(kratosx.MustContext(ctx), env)
	return &v1.AddEnvReply{Id: id}, err
}

func (s *EnvService) UpdateEnv(ctx context.Context, in *v1.UpdateEnvRequest) (*emptypb.Empty, error) {
	env := &biz.Env{}
	if err := copier.Copy(env, in); err != nil {
		return nil, errors.TransformError()
	}
	return nil, s.uc.UpdateEnv(kratosx.MustContext(ctx), env)
}

func (s *EnvService) DeleteEnv(ctx context.Context, in *v1.DeleteEnvRequest) (*emptypb.Empty, error) {
	return nil, s.uc.DeleteEnv(kratosx.MustContext(ctx), in.Id)
}

func (s *EnvService) GetEnvToken(ctx context.Context, in *v1.GetEnvTokenRequest) (*v1.GetEnvTokenReply, error) {
	env, err := s.uc.GetEnv(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetEnvTokenReply{Token: env.Token}, nil
}

func (s *EnvService) ResetEnvToken(ctx context.Context, in *v1.ResetEnvTokenRequest) (*v1.ResetEnvTokenReply, error) {
	token, err := s.uc.ResetEnvToken(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ResetEnvTokenReply{Token: token}, nil
}
