package biz

import (
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/biz/types"
	"github.com/limes-cloud/kratosx"
	ktypes "github.com/limes-cloud/kratosx/types"
)

type Server struct {
	ktypes.BaseModel
	Keyword     string `json:"keyword" gorm:"not null;type:char(32) binary;comment:服务标识"`
	Name        string `json:"name" gorm:"not null;size:64;comment:服务名称"`
	Description string `json:"description" gorm:"not null;size:128;comment:服务描述"`
}

type ServerRepo interface {
	Get(ctx kratosx.Context, id uint32) (*Server, error)
	GetByKeyword(ctx kratosx.Context, key string) (*Server, error)
	PageServer(ctx kratosx.Context, req *types.PageServerRequest) ([]*Server, uint32, error)
	All(ctx kratosx.Context, options ktypes.Scopes) ([]*Server, error)
	Create(ctx kratosx.Context, c *Server) (uint32, error)
	Update(ctx kratosx.Context, c *Server) error
	Delete(ctx kratosx.Context, uint322 uint32) error
}

type ServerUseCase struct {
	config *config.Config
	repo   ServerRepo
}

func NewServerUseCase(config *config.Config, repo ServerRepo) *ServerUseCase {
	return &ServerUseCase{config: config, repo: repo}
}

// Get 获取指定服务信息
func (u *ServerUseCase) Get(ctx kratosx.Context, id uint32) (*Server, error) {
	server, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return server, nil
}

// GetByKeyword 获取指定标识的服务信息
func (u *ServerUseCase) GetByKeyword(ctx kratosx.Context, keyword string) (*Server, error) {
	server, err := u.repo.GetByKeyword(ctx, keyword)
	if err != nil {
		return nil, v1.NotRecordError()
	}
	return server, nil
}

// Page 获取分页服务信息
func (u *ServerUseCase) Page(ctx kratosx.Context, req *types.PageServerRequest) ([]*Server, uint32, error) {
	list, total, err := u.repo.PageServer(ctx, req)
	if err != nil {
		return nil, 0, v1.DatabaseErrorFormat(err.Error())
	}
	return list, total, nil
}

// Add 添加服务信息
func (u *ServerUseCase) Add(ctx kratosx.Context, server *Server) (uint32, error) {
	id, err := u.repo.Create(ctx, server)
	if err != nil {
		return 0, v1.DatabaseErrorFormat(err.Error())
	}
	return id, nil
}

// Update 更新服务信息
func (u *ServerUseCase) Update(ctx kratosx.Context, server *Server) error {
	if err := u.repo.Update(ctx, server); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}

// Delete 删除服务信息
func (u *ServerUseCase) Delete(ctx kratosx.Context, id uint32) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		return v1.DatabaseErrorFormat(err.Error())
	}
	return nil
}
