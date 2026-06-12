package core

import (
	"os"

	kconfig "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/limes-cloud/kratosx/config"

	"github.com/limes-cloud/configure/api/configure/client"
)

var conf = &Conf{}

type Conf struct {
	Test string
	Res  string
}

func configSource() kconfig.Source {
	host := os.Getenv("CONF_HOST")
	token := os.Getenv("CONF_TOKEN")
	name := os.Getenv("APP_NAME")
	if host != "" && token != "" && name != "" {
		return client.New(host, name, token)
	}
	return file.NewSource("conf/")
}

func configScanWatch(watch config.Watcher) {
	watch("business", func(value config.Value) {
		if err := value.Scan(conf); err != nil {
			panic(err)
		}
	})
}
