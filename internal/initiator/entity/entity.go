package entity

import (
	"github.com/limes-cloud/configure/internal/core"
	"github.com/limes-cloud/kratosx/library/db"
	managerEntity "github.com/limes-cloud/manager/api/entity"
)

// registerPageSize 拉取已注册实体时的分页大小（manager 限制最大 50）
const registerPageSize = uint32(50)

// Register 注册 configure 的表结构到 manager（幂等：已存在的实体/字段跳过）
func Register(ctx core.Context) {
	ctx.BeforeStart("register configure entity", func() {
		conn, err := ctx.GrpcConn("Manager")
		if err != nil {
			ctx.Logger().Errorw("msg", "get manager grpc conn error", "err", err.Error())
			return
		}
		client := managerEntity.NewEntityClient(conn)

		// 拉取当前应用已注册的实体，作为幂等判断依据
		exists, err := loadRegisteredEntities(ctx, client)
		if err != nil {
			ctx.Logger().Errorw("msg", "load registered entity error", "err", err.Error())
			return
		}

		// 逐个注册表结构：实体已存在则复用其 id，字段按名幂等创建
		created, fields := 0, 0
		for _, item := range ctx.Database().Entities() {
			if item == nil || !shouldRegisterTable(item.Name) {
				continue
			}

			var id uint32
			if existing, ok := exists[item.Name]; ok {
				id = existing
			} else {
				reply, err := client.CreateEntity(ctx.Ctx(), &managerEntity.CreateEntityRequest{
					App:      appKeyword,
					Database: item.Database,
					Name:     item.Name,
					Comment:  defaultComment(item.Comment, item.Name),
				})
				if err != nil {
					ctx.Logger().Warnw("msg", "create entity error", "table", item.Name, "err", err.Error())
					continue
				}
				exists[item.Name] = reply.Id
				id = reply.Id
				created++
			}

			n, err := registerFields(ctx, client, id, item.Fields)
			if err != nil {
				ctx.Logger().Warnw("msg", "register entity fields error", "table", item.Name, "err", err.Error())
				continue
			}
			fields += n
		}

		ctx.Logger().Infow(
			"msg", "register configure entity success",
			"created", created,
			"fields", fields,
		)
	})
}

// loadRegisteredEntities 分页拉取当前应用已注册的实体，返回 name -> id 映射
func loadRegisteredEntities(ctx core.Context, client managerEntity.EntityClient) (map[string]uint32, error) {
	exists := map[string]uint32{}
	for page := uint32(1); ; page++ {
		reply, err := client.ListEntity(ctx.Ctx(), &managerEntity.ListEntityRequest{
			Page:     page,
			PageSize: registerPageSize,
			App:      appKeyword,
		})
		if err != nil {
			return nil, err
		}
		for _, item := range reply.List {
			if item != nil {
				exists[item.Name] = item.Id
			}
		}
		if len(reply.List) < int(registerPageSize) {
			break
		}
	}
	return exists, nil
}

// registerFields 注册实体字段（幂等：已存在的字段跳过），返回新建字段数
func registerFields(ctx core.Context, client managerEntity.EntityClient, entityId uint32, fields []db.Field) (int, error) {
	reply, err := client.ListEntityField(ctx.Ctx(), &managerEntity.ListEntityFieldRequest{EntityId: entityId})
	if err != nil {
		return 0, err
	}
	exists := map[string]struct{}{}
	for _, item := range reply.List {
		if item != nil {
			exists[item.Name] = struct{}{}
		}
	}

	created, index := 0, uint32(0)
	for _, field := range fields {
		if !shouldRegisterField(field.Name) {
			continue
		}
		index++
		if _, ok := exists[field.Name]; ok {
			continue
		}
		if _, err := client.CreateEntityField(ctx.Ctx(), &managerEntity.CreateEntityFieldRequest{
			EntityId: entityId,
			Name:     field.Name,
			Comment:  defaultComment(field.Comment, field.Name),
			Index:    index - 1,
		}); err != nil {
			return created, err
		}
		created++
	}
	return created, nil
}

// defaultComment 注释兜底（manager 校验要求 comment 不能为空）
func defaultComment(comment, fallback string) string {
	if comment == "" {
		return fallback
	}
	return comment
}
