package server

import (
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/api/errors"
	"github.com/limes-cloud/configure/internal/config"
)

type UseCase struct {
	config *config.Config
	repo   Repo
}

func NewUseCase(config *config.Config, repo Repo) *UseCase {
	return &UseCase{config: config, repo: repo}
}

// GetServer 获取指定服务信息
func (s *UseCase) GetServer(ctx kratosx.Context, id uint32) (*Server, error) {
	server, err := s.repo.GetServer(ctx, id)
	if err != nil {
		return nil, errors.NotRecordError()
	}
	return server, nil
}

// GetServerByKeyword 获取指定标识的服务信息
func (s *UseCase) GetServerByKeyword(ctx kratosx.Context, keyword string) (*Server, error) {
	server, err := s.repo.GetServerByKeyword(ctx, keyword)
	if err != nil {
		return nil, errors.NotRecordError()
	}
	return server, nil
}

// PageServer 获取分页服务信息
func (s *UseCase) PageServer(ctx kratosx.Context, req *PageServerRequest) ([]*Server, uint32, error) {
	list, total, err := s.repo.PageServer(ctx, req)
	if err != nil {
		return nil, 0, errors.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// AddServer 添加服务信息
func (s *UseCase) AddServer(ctx kratosx.Context, server *Server) (uint32, error) {
	id, err := s.repo.AddServer(ctx, server)
	if err != nil {
		return 0, errors.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// UpdateServer 更新服务信息
func (s *UseCase) UpdateServer(ctx kratosx.Context, server *Server) error {
	if err := s.repo.UpdateServer(ctx, server); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// DeleteServer 删除服务信息
func (s *UseCase) DeleteServer(ctx kratosx.Context, id uint32) error {
	if err := s.repo.DeleteServer(ctx, id); err != nil {
		return errors.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// // GetServerByIds 查询所有环境
// func (s *UseCase) GetServerByIds(ctx kratosx.Context, ids []uint32) ([]*Server, error) {
//	srvs, err := s.repo.GetServerByIds(ctx, ids)
//	if err != nil {
//		return nil, errors.DatabaseErrorFormat(err.Error())
//	}
//	return srvs, nil
// }
