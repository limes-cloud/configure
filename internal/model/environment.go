package model

import (
	"github.com/limes-cloud/kratos"
	"gorm.io/gorm"
)

// 1
//type Environment struct {
//	Keyword     string `json:"keyword" gorm:"unique;not null;size:32;comment:环境关键字"`
//	Title       string `json:"title" gorm:"not null;size:32;comment:环境标题"`
//	Description string `json:"description" gorm:"not null;size:128;comment:环境描述"`
//	Token       string `json:"token,omitempty" gorm:"not null;size:32;comment:连接token"`
//	Status      *bool  `json:"status" gorm:"not null;comment:启用状态"`
//	Operator    string `json:"operator,omitempty"`
//	OperatorId  int64  `json:"operator_id,omitempty"`
//}

func (e *Environment) Create(ctx kratos.Context) error {
	return ctx.DB().Model(e).Create(&e).Error
}

func (u *Environment) OneByKeyword(ctx kratos.Context, keyword string) error {
	return ctx.DB().First(u, "keyword = ?", keyword).Error
}

// 2
type Environment struct {
	Keyword     string `json:"keyword" gorm:"unique;not null;size:32;comment:环境关键字"`
	Title       string `json:"title" gorm:"not null;size:32;comment:环境标题"`
	Description string `json:"description" gorm:"not null;size:128;comment:环境描述"`
	Token       string `json:"token,omitempty" gorm:"not null;size:32;comment:连接token"`
	Status      *bool  `json:"status" gorm:"not null;comment:启用状态"`
	Operator    string `json:"operator,omitempty"`
	OperatorId  int64  `json:"operator_id,omitempty"`
}

type EnvironmentFactory struct {
	db *gorm.DB
}

func NewEnvironmentFactory(db *gorm.DB) *EnvironmentFactory {
	return &EnvironmentFactory{
		db: db,
	}
}

func (e *EnvironmentFactory) Create(in *Environment) error {
	return e.db.Create(in).Error
}

func (e *EnvironmentFactory) OneByKeyword(keyword string) (*Environment, error) {
	res := &Environment{}
	return res, e.db.First(res, "keyword = ?", keyword).Error
}

func (u *Environment) AllFilter(ctx *gin.Context) ([]Environment, error) {
	var list []Environment
	db := database(ctx).Table(u.Table()).Select("id,env_keyword")
	return list, transferErr(db.Find(&list).Error)
}

func (u *Environment) All(ctx *gin.Context, m any) ([]Environment, error) {
	var list []Environment
	db := database(ctx).Table(u.Table())
	db = gin.GormWhere(db, u.Table(), m)
	return list, transferErr(db.Find(&list).Error)
}

func (u *Environment) UpdateByID(ctx *gin.Context) error {
	user := meta.User(ctx)
	u.OperatorId = user.UserId
	u.Operator = user.UserName
	return transferErr(database(ctx).Table(u.Table()).Updates(u).Error)
}

func (u *Environment) DeleteByID(ctx *gin.Context, id int64) error {
	return transferErr(database(ctx).Table(u.Table()).Where("id = ?", id).Delete(&u).Error)
}
