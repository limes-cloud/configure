package entity

import (
	"context"

	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/kratosx/library/db"
	managerEntity "github.com/limes-cloud/manager/api/entity"
)

const registerModeAppend = "append"

func Register(ctx core.Context) {
	ctx.BeforeStart("register configure entity", func() {
		conn, err := ctx.GrpcConn("Manager")
		if err != nil {
			ctx.Logger().Errorw("msg", "get manager grpc conn error", "err", err.Error())
			return
		}

		reply, err := managerEntity.NewEntityClient(conn).RegisterEntity(context.Background(), &managerEntity.RegisterEntityRequest{
			AppKeyword: appKeyword,
			Mode:       registerModeAppend,
			Entities:   loadEntities(ctx),
		})
		if err != nil {
			ctx.Logger().Errorw("msg", "register configure entity error", "err", err.Error())
			return
		}

		ctx.Logger().Infow(
			"msg", "register configure entity success",
			"created", reply.Created,
			"updated", reply.Updated,
			"deleted", reply.Deleted,
		)
	})
}

func loadEntities(ctx core.Context) []*managerEntity.RegisterEntityNode {
	var list []*managerEntity.RegisterEntityNode
	for _, item := range ctx.Database().Entities() {
		if item == nil || !shouldRegisterTable(item.Name) {
			continue
		}

		list = append(list, &managerEntity.RegisterEntityNode{
			Database: item.Database,
			Name:     item.Name,
			Comment:  item.Comment,
			Fields:   loadFields(item.Fields),
		})
	}
	return list
}

func loadFields(fields []db.Field) []*managerEntity.RegisterEntityField {
	var list []*managerEntity.RegisterEntityField
	for _, item := range fields {
		if !shouldRegisterField(item.Name) {
			continue
		}
		list = append(list, registerField(item.Name, item.Comment, uint32(len(list))))
	}
	return list
}
