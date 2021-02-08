package others

import (
	"fitness/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//token中存在的用户信息
type TokenUserInfo struct {
	StoreId     uint   `json:"store_id"`
	StoreUserId uint   `json:"store_user_id"`
	LoginUserId uint   `json:"login_user_id"`
	UserType    string `json:"user_type"`
	ClientType  string `json:"client_type"`
	Currency    string `json:"currency"`
}

//生成token的信息
type Claims struct {
	TokenUserInfo
	BufferTime int64
	jwt.StandardClaims
}

//获取claims数据
func GetStandardClaims() jwt.StandardClaims {
	standardClaims := jwt.StandardClaims{
		NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
		ExpiresAt: time.Now().Unix() + global.Config.JWT.ExpiresTime, // 过期时间 7天  配置文件
		IssuedAt:  time.Now().Unix(),
		Issuer:    global.Config.GinConfig.AppName, // 签名颁发者
		Subject:   global.Config.JWT.Subject,       //签名主题
	}
	return standardClaims
}
