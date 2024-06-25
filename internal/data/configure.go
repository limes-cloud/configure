package data

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/kratosx/pkg/valx"

	biz "github.com/limes-cloud/configure/internal/biz/configure"
	"github.com/limes-cloud/configure/internal/data/model"
)

type configureRepo struct {
}

func NewConfigureRepo() biz.Repo {
	return &configureRepo{}
}

// ToConfigureEntity model转entity
func (r configureRepo) ToConfigureEntity(m *model.Configure) *biz.Configure {
	e := &biz.Configure{}
	_ = valx.Transform(m, e)
	return e
}

// ToConfigureModel entity转model
func (r configureRepo) ToConfigureModel(e *biz.Configure) *model.Configure {
	m := &model.Configure{}
	_ = valx.Transform(e, m)
	return m
}

// GetConfigureByEnvAndSrv 获取指定版本的模板
func (r configureRepo) GetConfigureByEnvAndSrv(ctx kratosx.Context, envId, srvId uint32) (*biz.Configure, error) {
	var (
		m  = model.Configure{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.Where("env_id=? and server_id=?", envId, srvId).First(&m).Error; err != nil {
		return nil, err
	}

	return r.ToConfigureEntity(&m), nil
}

// GetConfigure 获取指定的数据
func (r configureRepo) GetConfigure(ctx kratosx.Context, id uint32) (*biz.Configure, error) {
	var (
		m  = model.Configure{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}
	return r.ToConfigureEntity(&m), nil
}

// ListConfigure 获取列表
func (r configureRepo) ListConfigure(ctx kratosx.Context, req *biz.ListConfigureRequest) ([]*biz.Configure, uint32, error) {
	var (
		bs    []*biz.Configure
		ms    []*model.Configure
		total int64
		fs    = []string{"id", "format", "version", "description", "created_at", "updated_at"}
	)

	db := ctx.DB().Model(model.Configure{}).Select(fs).
		Where("server_id=?", req.ServerId).
		Where("env_id=?", req.EnvId)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	db = db.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))

	if err := db.Find(&ms).Error; err != nil {
		return nil, 0, err
	}

	for _, m := range ms {
		bs = append(bs, r.ToConfigureEntity(m))
	}
	return bs, uint32(total), nil
}

// CreateConfigure 创建数据
func (r configureRepo) CreateConfigure(ctx kratosx.Context, req *biz.Configure) (uint32, error) {
	m := r.ToConfigureModel(req)
	return m.Id, ctx.DB().Create(m).Error
}

// UpdateConfigure 更新数据
func (r configureRepo) UpdateConfigure(ctx kratosx.Context, req *biz.Configure) error {
	return ctx.DB().Updates(r.ToConfigureModel(req)).Error
}

// DeleteConfigure 删除数据
func (r configureRepo) DeleteConfigure(ctx kratosx.Context, id uint32) error {
	db := ctx.DB().Where("id=?", id).Delete(&model.Configure{})
	return db.Error
}

func (r configureRepo) CurrentTemplate(ctx kratosx.Context, srvId uint32) (string, string, error) {
	var (
		m  = model.Template{}
		fs = []string{"*"}
	)
	db := ctx.DB().Select(fs)
	if err := db.First(&m, "is_use=true and server_id=?", srvId).Error; err != nil {
		return "", "", err
	}
	return m.Format, m.Content, nil
}

func (r configureRepo) GetServerIdByKeyword(ctx kratosx.Context, keyword string) (uint32, error) {
	var id uint32
	if err := ctx.DB().Model(model.Server{}).
		Select("id").
		Where("keyword = ?", keyword).
		Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}

func (r configureRepo) GetEnvIdByToken(ctx kratosx.Context, token string) (uint32, error) {
	var id uint32
	if err := ctx.DB().Model(model.Env{}).
		Select("id").
		Where("token = ?", token).
		Scan(&id).Error; err != nil {
		return 0, err
	}
	return id, nil
}
