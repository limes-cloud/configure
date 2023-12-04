package handler

import (
	"context"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/kratosx"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) GetConfigure(ctx context.Context, in *v1.GetConfigureRequest) (*v1.GetConfigureReply, error) {
	return s.Configure.Get(kratosx.MustContext(ctx), in)
}

func (s *Service) UpdateConfigure(ctx context.Context, in *v1.UpdateConfigureRequest) (*emptypb.Empty, error) {
	return s.Configure.Update(kratosx.MustContext(ctx), in)
}

func (s *Service) WatchConfigure(in *v1.WatchConfigureRequest, reply v1.Service_WatchConfigureServer) error {
	return s.Configure.Watch(kratosx.MustContext(reply.Context()), in, reply)
}
