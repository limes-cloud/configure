package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/crypto"
	"github.com/limes-cloud/kratosx/pkg/valx"

	pb "github.com/limes-cloud/configure/api/configure/env/v1"
	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/biz/env"
	"github.com/limes-cloud/configure/internal/conf"
	"github.com/limes-cloud/configure/internal/data"
)

type EnvService struct {
	pb.UnimplementedEnvServer
	uc *env.UseCase
}

func NewEnvService(conf *conf.Config) *EnvService {
	return &EnvService{
		uc: env.NewUseCase(conf, data.NewEnvRepo()),
	}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewEnvService(c)
		pb.RegisterEnvHTTPServer(hs, srv)
		pb.RegisterEnvServer(gs, srv)
	})
}

// GetEnv 获取指定的环境信息
func (s *EnvService) GetEnv(c context.Context, req *pb.GetEnvRequest) (*pb.GetEnvReply, error) {
	var (
		in  = env.GetEnvRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, err := s.uc.GetEnv(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.GetEnvReply{}
	if err := valx.Transform(result, &reply); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}
	return &reply, nil
}

// ListEnv 获取环境信息列表
func (s *EnvService) ListEnv(c context.Context, req *pb.ListEnvRequest) (*pb.ListEnvReply, error) {
	var (
		in  = env.ListEnvRequest{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	result, total, err := s.uc.ListEnv(ctx, &in)
	if err != nil {
		return nil, err
	}

	reply := pb.ListEnvReply{Total: total}
	if err := valx.Transform(result, &reply.List); err != nil {
		ctx.Logger().Warnw("msg", "reply transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	return &reply, nil
}

// CreateEnv 创建环境信息
func (s *EnvService) CreateEnv(c context.Context, req *pb.CreateEnvRequest) (*pb.CreateEnvReply, error) {
	var (
		in = env.Env{
			Token: crypto.MD5ToUpper([]byte(uuid.NewString())),
		}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	id, err := s.uc.CreateEnv(ctx, &in)
	if err != nil {
		return nil, err
	}

	return &pb.CreateEnvReply{Id: id}, nil
}

// UpdateEnv 更新环境信息
func (s *EnvService) UpdateEnv(c context.Context, req *pb.UpdateEnvRequest) (*pb.UpdateEnvReply, error) {
	var (
		in  = env.Env{}
		ctx = kratosx.MustContext(c)
	)

	if err := valx.Transform(req, &in); err != nil {
		ctx.Logger().Warnw("msg", "req transform err", "err", err.Error())
		return nil, errors.TransformError()
	}

	if err := s.uc.UpdateEnv(ctx, &in); err != nil {
		return nil, err
	}

	return &pb.UpdateEnvReply{}, nil
}

// DeleteEnv 删除环境信息
func (s *EnvService) DeleteEnv(c context.Context, req *pb.DeleteEnvRequest) (*pb.DeleteEnvReply, error) {
	total, err := s.uc.DeleteEnv(kratosx.MustContext(c), req.Ids)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteEnvReply{Total: total}, nil
}

// UpdateEnvStatus 更新环境信息状态
func (s *EnvService) UpdateEnvStatus(c context.Context, req *pb.UpdateEnvStatusRequest) (*pb.UpdateEnvStatusReply, error) {
	return &pb.UpdateEnvStatusReply{}, s.uc.UpdateEnvStatus(kratosx.MustContext(c), req.Id, req.Status)
}

func (s *EnvService) GetEnvToken(c context.Context, req *pb.GetEnvTokenRequest) (*pb.GetEnvTokenReply, error) {
	res, err := s.uc.GetEnv(kratosx.MustContext(c), &env.GetEnvRequest{Id: &req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.GetEnvTokenReply{Token: res.Token}, nil
}

func (s *EnvService) ResetEnvToken(c context.Context, req *pb.ResetEnvTokenRequest) (*pb.ResetEnvTokenReply, error) {
	in := env.Env{
		Id:    req.Id,
		Token: crypto.MD5ToUpper([]byte(uuid.NewString())),
	}
	return &pb.ResetEnvTokenReply{Token: in.Token}, s.uc.UpdateEnv(kratosx.MustContext(c), &in)
}
