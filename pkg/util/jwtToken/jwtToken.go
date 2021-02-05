package jwtToken

import (
	"errors"
	"fitness/bean/others"
	"fitness/global"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.JwtSecret),
	}
}

//生成token
func (c *JWT) GenToken(claims *others.Claims) (string, error) {
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(global.JwtSecret)
}

//解析token
func (c *JWT) ParseToken(tokenString string) (*others.Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &others.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return global.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*others.Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid jwtToken")
}
