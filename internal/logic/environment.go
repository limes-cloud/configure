package logic

import (
	"github.com/google/uuid"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/md"
	"github.com/limes-cloud/configure/pkg/util"
	"github.com/limes-cloud/kratos"
	"google.golang.org/protobuf/types/known/emptypb"
)

// AllEnvironment 获取全部环境
func (l *Logic) AllEnvironment(ctx kratos.Context, in *emptypb.Empty) (*v1.AllEnvironmentReply, error) {
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

// AddEnvironment 新增环境
func (l *Logic) AddEnvironment(ctx kratos.Context, in *v1.AddEnvironmentRequest) (*emptypb.Empty, error) {
	env := &model.Environment{
		Token:      uuid.NewString(),
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}
	if util.Transform(in, env) != nil {
		return nil, v1.ErrorTransform()
	}

	if env.Create(ctx) != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}

// UpdateEnvironment 修改环境
func (l *Logic) UpdateEnvironment(ctx kratos.Context, in *v1.UpdateEnvironmentRequest) (*emptypb.Empty, error) {
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

// DeleteEnvironment 删除环境
func (l *Logic) DeleteEnvironment(ctx kratos.Context, in *v1.DeleteEnvironmentRequest) (*emptypb.Empty, error) {
	env := &model.Environment{}

	if env.DeleteByID(ctx, in.Id) != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}

// GetEnvironmentToken 获取环境token
func (l *Logic) GetEnvironmentToken(ctx kratos.Context, in *v1.GetEnvironmentTokenRequest) (*v1.GetEnvironmentTokenReply, error) {
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

// ResetEnvironmentToken 重置环境token
func (l *Logic) ResetEnvironmentToken(ctx kratos.Context, in *v1.ResetEnvironmentTokenRequest) (*emptypb.Empty, error) {
	env := &model.Environment{
		Token:      uuid.NewString(),
		Operator:   md.GetUserName(ctx),
		OperatorID: md.GetUserID(ctx),
	}

	if env.UpdateByID(ctx, in.Id) != nil {
		return nil, v1.ErrorDatabase()
	}

	return nil, nil
}
