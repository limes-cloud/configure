package configure

import (
	"github.com/limes-cloud/kratosx"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/initiator/server"
	"github.com/limes-cloud/configure/internal/logic"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/pt"
)

func Init(ctx kratosx.Context, config *config.Config) {
	lc := logic.NewConfigure(config)

	env := model.Environment{}
	envs, _ := env.All(ctx)

	for _, env := range envs {
		for _, srv := range server.Servers {
			if _, err := lc.Update(ctx, &v1.UpdateConfigureRequest{
				EnvKeyword: env.Keyword,
				ServerId:   srv,
			}); err != nil {
				pt.Error("初始化configure失败：" + err.Error())
				return
			}
		}
	}
}
