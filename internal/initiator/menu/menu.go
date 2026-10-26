package menu

import (
	"github.com/limes-cloud/configure/internal/core"
	managerMenu "github.com/limes-cloud/manager/api/menu"
)

const registerModeAppend = "append"

func Register(ctx core.Context) {
	ctx.BeforeStart("register configure menu", func() {
		conn, err := ctx.GrpcConn("Manager")
		if err != nil {
			ctx.Logger().Errorw("msg", "get manager grpc conn error", "err", err.Error())
			return
		}

		reply, err := managerMenu.NewMenuClient(conn).RegisterMenu(ctx.Ctx(), &managerMenu.RegisterMenuRequest{
			AppKeyword: appKeyword,
			Mode:       registerModeAppend,
			Menus:      menus,
		})
		if err != nil {
			ctx.Logger().Errorw("msg", "register configure menu error", "err", err.Error())
			return
		}

		ctx.Logger().Infow(
			"msg", "register configure menu success",
			"created", reply.Created,
			"updated", reply.Updated,
			"deleted", reply.Deleted,
		)
	})
}
