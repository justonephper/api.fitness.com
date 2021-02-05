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
	jwt.StandardClaims
}

//获取claims数据
func GetStandardClaims() jwt.StandardClaims {
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(global.TokenExpireDuration).Unix(), // 过期时间
		IssuedAt:  time.Now().Unix(),
		Issuer:    global.APP_NAME, // 签名颁发者
		Subject:   "user jwtToken", //签名主题
	}
	return standardClaims
}
