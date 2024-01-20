package configure

import (
	"fmt"

	"github.com/limes-cloud/configure/internal/service"

	"github.com/limes-cloud/configure/internal/data"

	"github.com/limes-cloud/kratosx"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/internal/initiator/server"
	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context, service *service.Service) {
	env := data.NewEnvRepo()
	envs, _ := env.All(ctx)
	for _, env := range envs {
		for _, srv := range server.Servers {
			if _, err := service.UpdateConfigure(ctx.Ctx(), &v1.UpdateConfigureRequest{
				ServerId:    srv,
				EnvId:       env.ID,
				Description: "自动初始化",
			}); err != nil {
				pt.Error(fmt.Sprintf("初始化configure失败 srv:%v env:%v err:%s", srv, env.Keyword, err.Error()))
				return
			}
		}
	}
}
