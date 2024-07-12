package permission

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"
	"github.com/limes-cloud/manager/api/manager/auth"
	resourcev1 "github.com/limes-cloud/manager/api/manager/resource/v1"

	"github.com/limes-cloud/configure/api/configure/errors"
)

const (
	Manager = "Manager"
	Env     = "cfg_env"
	Server  = "cfg_server"
)

func NewClient(ctx kratosx.Context) (resourcev1.ResourceClient, error) {
	conn, err := kratosx.MustContext(ctx).GrpcConn(Manager)
	if err != nil {
		return nil, errors.ManagerServerError()
	}
	return resourcev1.NewResourceClient(conn), nil
}

func GetPermission(ctx kratosx.Context, keyword string) (bool, []uint32, error) {
	info, err := auth.GetAuthInfo(ctx)
	if err != nil {
		return false, nil, err
	}
	if info.UserId == 1 || info.RoleId == 1 {
		return true, nil, nil
	}

	client, err := NewClient(ctx)
	if err != nil {
		return false, nil, err
	}
	reply, err := client.GetResourceScopes(ctx, &resourcev1.GetResourceScopesRequest{
		Keyword: keyword,
		UserId:  info.UserId,
	})

	if err != nil {
		return false, nil, err
	}
	return reply.All, reply.Scopes, nil
}

func GetEnv(ctx kratosx.Context) (bool, []uint32, error) {
	all, ids, err := GetPermission(ctx, Env)
	if ids == nil {
		ids = []uint32{}
	}
	return all, ids, err
}

func HasEnv(ctx kratosx.Context, id uint32) bool {
	all, ids, err := GetPermission(ctx, Env)
	if err != nil {
		return false
	}
	if all {
		return true
	}
	return valx.InList(ids, id)
}

func GetServer(ctx kratosx.Context) (bool, []uint32, error) {
	all, ids, err := GetPermission(ctx, Server)
	if ids == nil {
		ids = []uint32{}
	}
	return all, ids, err
}

func HasServer(ctx kratosx.Context, id uint32) bool {
	all, ids, err := GetPermission(ctx, Server)
	if err != nil {
		return false
	}
	if all {
		return true
	}
	return valx.InList(ids, id)
}
