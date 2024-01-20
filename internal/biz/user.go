package biz

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/forgoer/openssl"
	v1 "github.com/limes-cloud/configure/api/v1"
	"github.com/limes-cloud/configure/config"
	"github.com/limes-cloud/configure/internal/biz/types"
	"github.com/limes-cloud/kratosx"
)

type UserUseCase struct {
	config *config.Config
}

func NewUserUseCase(config *config.Config) *UserUseCase {
	return &UserUseCase{config: config}
}

// Login 用户登录
func (uc *UserUseCase) Login(ctx kratosx.Context, username, password string) (string, error) {
	// 密码解密
	pwByte, _ := base64.StdEncoding.DecodeString(password)
	decryptData, err := openssl.RSADecrypt(pwByte, ctx.Loader("login"))
	if err != nil {
		return "", v1.RsaDecodeErrorFormat(err.Error())
	}

	// 序列化密码
	var pw types.Password
	if json.Unmarshal(decryptData, &pw) != nil {
		return "", v1.PasswordFormatError()
	}

	// 判断当前时间戳是否过期,超过10s则拒绝
	if time.Now().UnixMilli()-pw.Time > 10*1000 {
		return "", v1.PasswordExpireError()
	}

	// 判断用户密码是否正确
	if uc.config.Author.AdminPassword != pw.Password || uc.config.Author.AdminUser != username {
		return "", v1.UserPasswordError()
	}

	// 生成登陆token
	token, err := ctx.JWT().NewToken(map[string]any{})
	if err != nil {
		return "", v1.NewTokenErrorFormat(err.Error())
	}

	return token, nil
}

// RefreshToken 刷新用户token
func (uc *UserUseCase) RefreshToken(ctx kratosx.Context) (string, error) {
	token, err := ctx.JWT().Renewal(ctx)
	if err != nil {
		return "", v1.RefreshTokenErrorFormat(err.Error())
	}
	return token, nil
}
