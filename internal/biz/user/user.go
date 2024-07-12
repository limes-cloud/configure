package user

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/forgoer/openssl"
	"github.com/limes-cloud/kratosx"

	"github.com/limes-cloud/configure/api/configure/errors"
	"github.com/limes-cloud/configure/internal/conf"
)

type UseCase struct {
	config *conf.Config
}

func NewUseCase(config *conf.Config) *UseCase {
	return &UseCase{config: config}
}

type Password struct {
	Password string `json:"password"`
	Time     int64  `json:"time"`
}

// Login 用户登录
func (uc *UseCase) Login(ctx kratosx.Context, username, password string) (string, error) {
	// 密码解密
	pwByte, _ := base64.StdEncoding.DecodeString(password)
	decryptData, err := openssl.RSADecrypt(pwByte, ctx.Loader("login"))
	if err != nil {
		return "", errors.PasswordError(err.Error())
	}

	// 序列化密码
	var pw Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return "", errors.PasswordError()
	}

	// 判断当前时间戳是否过期,超过10s则拒绝
	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		return "", errors.PasswordExpireError()
	}

	// 判断用户密码是否正确
	if uc.config.Author.AdminPassword != pw.Password || uc.config.Author.AdminUser != username {
		return "", errors.PasswordError()
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(map[string]any{"userId": 1, "roleKeyword": "superAdmin"})
	if err != nil {
		return "", errors.SystemError(err.Error())
	}

	return token, nil
}

// RefreshToken 刷新用户token
func (uc *UseCase) RefreshToken(ctx kratosx.Context) (string, error) {
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return "", errors.RefreshTokenError(err.Error())
	}
	return token, nil
}
