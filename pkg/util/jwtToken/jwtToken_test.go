package jwtToken

import (
	"fitness/bean/others"
	"fitness/global"
	"testing"
)

func TestGetToken(t *testing.T) {
	obj := NewJWT()
	token, err := _genToken(obj)
	if err != nil {
		t.Log("gen token failed")
	}
	t.Logf("token:%s", token)
}

//生成token
func _genToken(jwt *JWT) (token string, err error) {
	userInfo := others.TokenUserInfo{
		StoreId:     1,
		StoreUserId: 1,
		LoginUserId: 1,
		UserType:    global.RoleAdmin,
		ClientType:  global.ClientTypeWeb,
		Currency:    "CN",
	}
	tokenUserInfo := others.GetTokenUserInfo(userInfo, 500)
	return jwt.GenToken(tokenUserInfo)
}

func TestParseToken(t *testing.T) {
	obj := NewJWT()
	token, _ := _genToken(obj)
	data, _ := obj.ParseToken(token)
	t.Log(*data)
}
