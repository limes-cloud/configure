package rpc

import (
	"sync"

	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/auth"
	v1 "github.com/limes-cloud/manager/api/manager/auth/v1"
	resourcev1 "github.com/limes-cloud/manager/api/manager/resource/v1"

	"github.com/limes-cloud/configure/api/configure/errors"
)

const (
	Manager = "Manager"
	Env     = "cfg_env"
	Server  = "cfg_server"
)

type PermissionInfra struct {
}

var (
	permissionIns  *PermissionInfra
	permissionOnce sync.Once
)

func NewPermissionInfra() *PermissionInfra {
	permissionOnce.Do(func() {
		permissionIns = &PermissionInfra{}
	})
	return permissionIns
}

func client(ctx kratosx.Context) (resourcev1.ResourceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Manager)
	if err != nil {
		return nil, errors.ManagerServerError()
	}
	return resourcev1.NewResourceClient(conn), nil
}

// GetPermission 获取当前用户，指定key的权限
func (infra *PermissionInfra) GetPermission(ctx kratosx.Context, keyword string) (bool, []uint32, error) {
	var (
		info = &v1.AuthReply{}
		err  error
	)
	if ctx.Token() != "" {
		err = ctx.JWT().Parse(ctx, info)
	} else {
		info, err = auth.GetAuthInfo(ctx)
	}

	if err != nil {
		return false, nil, err
	}
	if info.UserId == 1 || info.RoleId == 1 {
		return true, nil, nil
	}

	c, err := client(ctx)
	if err != nil {
		return false, nil, err
	}
	reply, err := c.GetResourceScopes(ctx, &resourcev1.GetResourceScopesRequest{
		Keyword: keyword,
	})

	if err != nil {
		return false, nil, err
	}
	return reply.All, reply.Scopes, nil
}

// GetEnv 获取当前用户对于env的权限
func (infra *PermissionInfra) GetEnv(ctx kratosx.Context) (bool, []uint32, error) {
	all, ids, err := infra.GetPermission(ctx, Env)
	if ids == nil {
		ids = []uint32{}
	}
	return all, ids, err
}

// HasEnv 获取当前用户是否具有指定env的权限
func (infra *PermissionInfra) HasEnv(ctx kratosx.Context, id uint32) bool {
	all, ids, err := infra.GetPermission(ctx, Env)
	if err != nil {
		return false
	}
	if all {
		return true
	}
	return valx.InList(ids, id)
}

// GetServer 获取当前用户是对于server的权限
func (infra *PermissionInfra) GetServer(ctx kratosx.Context) (bool, []uint32, error) {
	all, ids, err := infra.GetPermission(ctx, Server)
	if ids == nil {
		ids = []uint32{}
	}
	return all, ids, err
}

// HasServer 获取当前用户是具有指定server的权限
func (infra *PermissionInfra) HasServer(ctx kratosx.Context, id uint32) bool {
	all, ids, err := infra.GetPermission(ctx, Server)
	if err != nil {
		return false
	}
	if all {
		return true
	}
	return valx.InList(ids, id)
}
