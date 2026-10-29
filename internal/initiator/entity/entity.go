package entity

import (
	"github.com/limes-cloud/configure/internal/core"
	managerEntity "github.com/limes-cloud/manager/api/entity"
)

const registerModeAppend = "append"

// Register 注册 configure 的表结构到 manager（幂等：已存在的实体/字段跳过）
func Register(ctx core.Context) {
	ctx.BeforeStart("register configure entity", func() {
		conn, err := ctx.GrpcConn("Manager")
		if err != nil {
			ctx.Logger().Errorw("msg", "get manager grpc conn error", "err", err.Error())
			return
		}
		client := managerEntity.NewEntityClient(conn)

		entities := registerEntities(ctx)
		reply, err := client.RegisterEntity(ctx.Ctx(), &managerEntity.RegisterEntityRequest{
			AppKeyword: appKeyword,
			Mode:       registerModeAppend,
			Entities:   entities,
		})
		if err != nil {
			ctx.Logger().Errorw("msg", "register configure entity error", "err", err.Error())
			return
		}

		ctx.Logger().Infow(
			"msg", "register configure entity success",
			"entities", len(entities),
			"created", reply.Created,
			"updated", reply.Updated,
			"deleted", reply.Deleted,
		)
	})
}

func registerEntities(ctx core.Context) []*managerEntity.RegisterEntityNode {
	entities := make([]*managerEntity.RegisterEntityNode, 0)
	for _, table := range ctx.Database().Entities() {
		if table == nil || !shouldRegisterTable(table.Name) {
			continue
		}

		node := &managerEntity.RegisterEntityNode{
			Database: table.Database,
			Name:     table.Name,
			Comment:  defaultComment(table.Comment, table.Name),
		}
		for _, field := range table.Fields {
			if !shouldRegisterField(field.Name) {
				continue
			}
			node.Fields = append(node.Fields, &managerEntity.RegisterEntityField{
				Name:    field.Name,
				Comment: defaultComment(field.Comment, field.Name),
				Index:   uint32(len(node.Fields)),
			})
		}
		entities = append(entities, node)
	}
	return entities
}

// defaultComment 注释兜底（manager 校验要求 comment 不能为空）
func defaultComment(comment, fallback string) string {
	if comment == "" {
		return fallback
	}
	return comment
}
