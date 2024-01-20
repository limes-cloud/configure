package client

import (
	"testing"
	"time"

	"github.com/limes-cloud/kratosx/config"
)

func TestSource_Watch(t *testing.T) {
	source := New("localhost:6082", "Resource", "220151E8F2B6AB8A82BA3A9AD576883F")
	conf := config.New(source)
	if err := conf.Load(); err != nil {
		t.Error(err.Error())
		return
	}
	conf.Watch("file", func(value config.Value) {
		t.Log("配置变更")
	})

	time.Sleep(100 * time.Second)
}
