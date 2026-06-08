package client

import (
	"context"
	"os"
	"sync"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	pb "github.com/limes-cloud/configure/api/configure"
)

type source struct {
	host    string
	server  string
	token   string
	request pb.ConfigureClient

	mu     sync.Mutex
	client pb.Configure_WatchConfigureClient
	cancel context.CancelFunc
}

func New(host, name, token string) config.Source {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(host))
	if err != nil {
		panic("configure connect error:" + err.Error())
	}
	return &source{
		request: pb.NewConfigureClient(conn),
		host:    host,
		server:  name,
		token:   token,
	}
}

func NewFromEnv() config.Source {
	host := os.Getenv("CONF_HOST")
	token := os.Getenv("CONF_TOKEN")
	name := os.Getenv("APP_NAME")
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(host))
	if err != nil {
		panic("configure connect error:" + err.Error())
	}
	return &source{
		request: pb.NewConfigureClient(conn),
		host:    host,
		server:  name,
		token:   token,
	}
}

// Load return the config values
func (s *source) Load() ([]*config.KeyValue, error) {
	s.mu.Lock()
	// 取消旧 stream
	if s.cancel != nil {
		s.cancel()
	}
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	s.mu.Unlock()

	client, err := s.request.WatchConfigure(ctx, &pb.WatchConfigureRequest{
		Server: s.server,
		Token:  s.token,
	})
	if err != nil {
		cancel()
		return nil, err
	}

	cf, err := client.Recv()
	if err != nil {
		cancel()
		return nil, err
	}

	s.mu.Lock()
	s.client = client
	s.mu.Unlock()

	return []*config.KeyValue{{
		Key:    s.host,
		Value:  []byte(cf.Content),
		Format: cf.Format,
	}}, nil
}

// Watch return the watcher
func (s *source) Watch() (config.Watcher, error) {
	return newWatcher(s)
}
