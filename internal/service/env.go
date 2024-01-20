package service

import (
	"context"

	"github.com/limes-cloud/configure/internal/biz"

	"github.com/jinzhu/copier"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) AllEnv(ctx context.Context, in *emptypb.Empty) (*v1.AllEnvReply, error) {
	envs, err := s.EnvUseCase.All(kratosx.MustContext(ctx))
	if err != nil {
		return nil, err
	}
	reply := v1.AllEnvReply{}
	if err := copier.Copy(&reply.List, envs); err != nil {
		return nil, v1.TransformError()
	}
	return &reply, nil
}

func (s *Service) AddEnv(ctx context.Context, in *v1.AddEnvRequest) (*emptypb.Empty, error) {
	env := &biz.Env{}
	if err := copier.Copy(env, in); err != nil {
		return nil, v1.TransformError()
	}
	_, err := s.EnvUseCase.Add(kratosx.MustContext(ctx), env)
	return nil, err
}

func (s *Service) UpdateEnv(ctx context.Context, in *v1.UpdateEnvRequest) (*emptypb.Empty, error) {
	env := &biz.Env{}
	if err := copier.Copy(env, in); err != nil {
		return nil, v1.TransformError()
	}
	return nil, s.EnvUseCase.Update(kratosx.MustContext(ctx), env)
}

func (s *Service) DeleteEnv(ctx context.Context, in *v1.DeleteEnvRequest) (*emptypb.Empty, error) {
	return nil, s.EnvUseCase.Delete(kratosx.MustContext(ctx), in.Id)
}

func (s *Service) GetEnvToken(ctx context.Context, in *v1.GetEnvTokenRequest) (*v1.GetEnvTokenReply, error) {
	env, err := s.EnvUseCase.Get(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetEnvTokenReply{Token: env.Token}, nil
}

func (s *Service) ResetEnvToken(ctx context.Context, in *v1.ResetEnvTokenRequest) (*v1.ResetEnvTokenReply, error) {
	token, err := s.EnvUseCase.ResetToken(kratosx.MustContext(ctx), in.Id)
	if err != nil {
		return nil, err
	}
	return &v1.ResetEnvTokenReply{Token: token}, nil
}
