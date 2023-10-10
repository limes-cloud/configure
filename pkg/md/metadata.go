package md

import (
	"strconv"

	"github.com/limes-cloud/kratos"
)

const (
	uid  = "user-id"
	name = "user-name"
)

func GetUserID(ctx kratos.Context) int64 {
	res := ctx.GetMetadata(uid)
	if res == "" {
		return 0
	}
	id, _ := strconv.ParseInt(res, 10, 64)
	return id
}
func GetUserName(ctx kratos.Context) string {
	return ctx.GetMetadata(name)

}
