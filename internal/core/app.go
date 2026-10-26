package core

import (
	"github.com/limes-cloud/kratosx"
)

func InitApp(opts ...kratosx.Option) *kratosx.App {
	defOpts := []kratosx.Option{
		kratosx.WithConfigSource(configSource()),
	}
	return kratosx.New(append(defOpts, opts...)...)
}
