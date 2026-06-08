package client

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/config"
)

type watcher struct {
	source *source
	ch     chan []*config.KeyValue
	ctx    context.Context
	cancel context.CancelFunc
}

func newWatcher(s *source) (*watcher, error) {
	ctx, cancel := context.WithCancel(context.Background())
	w := &watcher{
		source: s,
		ch:     make(chan []*config.KeyValue, 1),
		ctx:    ctx,
		cancel: cancel,
	}

	go func() {
		for {
			select {
			case <-w.ctx.Done():
				return
			default:
			}

			s.mu.Lock()
			stream := s.client
			s.mu.Unlock()

			if stream == nil {
				time.Sleep(time.Second)
				continue
			}

			cf, err := stream.Recv()
			if err != nil {
				// stream 断开，触发重连
				for {
					select {
					case <-w.ctx.Done():
						return
					default:
					}
					kvs, er := s.Load()
					if er != nil {
						time.Sleep(5 * time.Second)
						continue
					}
					select {
					case w.ch <- kvs:
					case <-w.ctx.Done():
						return
					}
					break
				}
				continue
			}

			select {
			case w.ch <- []*config.KeyValue{{
				Key:    s.host,
				Value:  []byte(cf.Content),
				Format: cf.Format,
			}}:
			case <-w.ctx.Done():
				return
			}
		}
	}()

	return w, nil
}

func (w *watcher) Next() ([]*config.KeyValue, error) {
	select {
	case res := <-w.ch:
		return res, nil
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	}
}

func (w *watcher) Stop() error {
	w.cancel()
	return nil
}
