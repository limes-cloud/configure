package logic

import (
	"github.com/google/uuid"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratosx"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Environment struct {
	conf *config.Config
}

func NewEnvironment(conf *config.Config) *Environment {
	return &Environment{
		conf: conf,
	}
}

// All 获取全部环境
func (e *Environment) All(ctx kratosx.Context, in *emptypb.Empty) (*v1.AllEnvironmentReply, error) {
	env := model.Environment{}
	list, err := env.All(ctx)
	if err != nil {
		return nil, v1.DatabaseError()
	}

	reply := v1.AllEnvironmentReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

// Add 新增环境
func (e *Environment) Add(ctx kratosx.Context, in *v1.AddEnvironmentRequest) (*emptypb.Empty, error) {
	env := &model.Environment{
		Token: util.MD5ToUpper([]byte(uuid.NewString()))}

	// 查询keyword是否存在
	if env.OneByKeyword(ctx, in.Keyword) == nil {
		return nil, v1.AlreadyExistsError()
	}

	if util.Transform(in, env) != nil {
		return nil, v1.TransformError()
	}

	env.Status = proto.Bool(true)

	if err := env.Create(ctx); err != nil {
		return nil, v1.DatabaseError()
	}

	return nil, nil
}

// Update 修改环境
func (e *Environment) Update(ctx kratosx.Context, in *v1.UpdateEnvironmentRequest) (*emptypb.Empty, error) {
	env := &model.Environment{}
	if util.Transform(in, env) != nil {
		return nil, v1.TransformError()
	}

	if env.UpdateByID(ctx, in.Id) != nil {
		return nil, v1.DatabaseError()
	}

	return nil, nil
}

// Delete 删除环境
func (e *Environment) Delete(ctx kratosx.Context, in *v1.DeleteEnvironmentRequest) (*emptypb.Empty, error) {
	env := &model.Environment{}

	if env.DeleteByID(ctx, in.Id) != nil {
		return nil, v1.DatabaseError()
	}

	return nil, nil
}

// GetToken 获取环境token
func (e *Environment) GetToken(ctx kratosx.Context, in *v1.GetEnvironmentTokenRequest) (*v1.GetEnvironmentTokenReply, error) {
	env := &model.Environment{}

	if env.OneByID(ctx, in.Id) != nil {
		return nil, v1.DatabaseError()
	}

	reply := v1.GetEnvironmentTokenReply{}
	if util.Transform(env, &reply) != nil {
		return nil, v1.TransformError()
	}

	return &reply, nil
}

// ResetToken 重置环境token
func (e *Environment) ResetToken(ctx kratosx.Context, in *v1.ResetEnvironmentTokenRequest) (*v1.ResetEnvironmentTokenReply, error) {
	env := &model.Environment{
		Token: util.MD5ToUpper([]byte(uuid.NewString()))}

	if env.UpdateByID(ctx, in.Id) != nil {
		return nil, v1.DatabaseError()
	}

	return &v1.ResetEnvironmentTokenReply{
		Token: env.Token,
	}, nil
}
