package logic

import (
	"github.com/google/uuid"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/md"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
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
func (e *Environment) All(ctx kratos.Context, in *emptypb.Empty) (*v1.AllEnvironmentReply, error) {
	env := model.Environment{}
	list, err := env.All(ctx)
	if err != nil {
		return nil, v1.ErrorDatabase()
	}

	reply := v1.AllEnvironmentReply{}
	if util.Transform(list, &reply.List) != nil {
		return nil, v1.ErrorTransform()
	}

	return &reply, nil
}

// Add 新增环境
func (e *Environment) Add(ctx kratos.Context, in *v1.AddEnvironmentRequest) (*emptypb.Empty, error) {
	env := &model.Environment{
		Token:      util.MD5ToUpper([]byte(uuid.NewString())),
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}

	// 查询keyword是否存在
	if env.OneByKeyword(ctx, in.Keyword) == nil {
		return nil, v1.ErrorAlreadyExists()
	}

	if util.Transform(in, env) != nil {
		return nil, v1.ErrorTransform()
	}

	if err := env.Create(ctx); err != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}

// Update 修改环境
func (e *Environment) Update(ctx kratos.Context, in *v1.UpdateEnvironmentRequest) (*emptypb.Empty, error) {
	env := &model.Environment{
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if util.Transform(in, env) != nil {
		return nil, v1.ErrorTransform()
	}

	if env.UpdateByID(ctx, in.Id) != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}

// Delete 删除环境
func (e *Environment) Delete(ctx kratos.Context, in *v1.DeleteEnvironmentRequest) (*emptypb.Empty, error) {
	env := &model.Environment{}

	if env.DeleteByID(ctx, in.Id) != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}

// GetToken 获取环境token
func (e *Environment) GetToken(ctx kratos.Context, in *v1.GetEnvironmentTokenRequest) (*v1.GetEnvironmentTokenReply, error) {
	env := &model.Environment{}

	if env.OneByID(ctx, in.Id) != nil {
		return nil, v1.ErrorDatabase()
	}

	reply := v1.GetEnvironmentTokenReply{}
	if util.Transform(env, &reply) != nil {
		return nil, v1.ErrorTransform()
	}

	return &reply, nil
}

// ResetToken 重置环境token
func (e *Environment) ResetToken(ctx kratos.Context, in *v1.ResetEnvironmentTokenRequest) (*emptypb.Empty, error) {
	env := &model.Environment{
		Token:      util.MD5ToUpper([]byte(uuid.NewString())),
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}

	if env.UpdateByID(ctx, in.Id) != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}
