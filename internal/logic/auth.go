package logic

import (
	"encoding/base64"
	"encoding/json"
	"time"

	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/model"
	"github.com/limes-cloud/configure/pkg/md"

	"github.com/forgoer/openssl"
	"github.com/limes-cloud/kratos"
)

type Auth struct {
	conf *config.Config
}

func NewAuth(conf *config.Config) *Auth {
	return &Auth{
		conf: conf,
	}
}

const (
	imageCaptchaTp = "login"
	passwordCert   = "login"
)

type Password struct {
	Password string `json:"password"`
	Time     int64  `json:"time"`
}

var adminUser = &model.User{
	ID:                1,
	DepartmentKeyword: "company",
	RoleID:            1,
	RoleKeyword:       "super_admin",
}

// Login 用户登录
func (a *Auth) Login(ctx kratos.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	// 密码解密
	passByte, _ := base64.StdEncoding.DecodeString(in.Password)
	decryptData, err := openssl.RSADecrypt(passByte, ctx.Loader(passwordCert))
	if err != nil {
		return nil, v1.ErrorRsaDecodeFormat(err.Error())
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return nil, v1.ErrorPasswordFormat()
	}

	// 判断当前时间戳是否过期,超过10s则拒绝
	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		return nil, v1.ErrorPasswordExpire()
	}

	if a.conf.AdminPassword != pw.Password || a.conf.AdminUser != in.Username {
		return nil, v1.ErrorUserPassword()
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(md.New(adminUser))
	if err != nil {
		return nil, v1.ErrorNewTokenFormat(err.Error())
	}

	return &v1.LoginReply{Token: token}, nil
}

// RefreshToken 刷新用户token
func (a *Auth) RefreshToken(ctx kratos.Context) (*v1.RefreshTokenReply, error) {
	// 进行token续期
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return nil, v1.ErrorRefreshTokenFormat(err.Error())
	}

	return &v1.RefreshTokenReply{Token: token}, nil
}
