package env

import (
	"github.com/google/uuid"
	"github.com/limes-cloud/configure/internal/biz"
	"github.com/limes-cloud/kratosx/types"
	"google.golang.org/protobuf/proto"

	"github.com/limes-cloud/configure/pkg/util"
)

const (
	TEST = uint32(1)
	PRE  = uint32(2)
	PROD = uint32(3)
)

var Envs = []uint32{TEST, PRE, PROD}

var envs = []*biz.Env{
	{
		BaseModel:   types.BaseModel{ID: TEST},
		Keyword:     "TEST",
		Name:        "测试环境",
		Token:       "1025D32F6CA7A2A320FE091B22C5DF3C",
		Status:      proto.Bool(true),
		Description: "用于本地测试",
	},
	{
		BaseModel:   types.BaseModel{ID: PRE},
		Keyword:     "PRE",
		Name:        "预发布环境",
		Token:       util.MD5ToUpper([]byte(uuid.NewString())),
		Status:      proto.Bool(true),
		Description: "用于上线前测试",
	},
	{
		BaseModel:   types.BaseModel{ID: PROD},
		Keyword:     "PROD",
		Name:        "生产环境",
		Token:       "5B655B7D4A51BF79C974C9F27C27D992",
		Status:      proto.Bool(true),
		Description: "用于线上真实环境",
	},
}
