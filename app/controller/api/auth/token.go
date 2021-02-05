package auth

import (
	"fitness/app/helper/response"
	"fitness/bean/others"
	"fitness/pkg/code"
	"fitness/pkg/util/jwtToken"
	"github.com/gin-gonic/gin"
)

//生成token
func GenToken(c *gin.Context) {
	userInfo := others.TokenUserInfo{
		StoreId:     10,
		StoreUserId: 9,
		LoginUserId: 8,
		UserType:    "admin",
		ClientType:  "web",
		Currency:    "CN",
	}

	standardClaims := others.GetStandardClaims()
	claims := &others.Claims{
		TokenUserInfo: userInfo,
		StandardClaims: standardClaims,
	}

	j := jwtToken.NewJWT()
	token, err := j.GenToken(claims)
	if err != nil {
		response.Failed(c, code.Failed, nil)
		return
	}
	response.Success(c, token)
	return
}

//解析token
func ParseToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		response.Failed(c, code.BadRequestParams, nil)
		return
	}
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdG9yZV9pZCI6MTAsInN0b3JlX3VzZXJfaWQiOjksImxvZ2luX3VzZXJfaWQiOjgsInVzZXJfdHlwZSI6IkFkbWluIiwiY2xpZW50X3R5cGUiOiJXZWIiLCJjdXJyZW5jeSI6IkNOIiwiZXhwIjoxNjEyNTE4MDcyLCJpYXQiOjE2MTI1MTA4NzIsImlzcyI6ImZpdG5lc3MiLCJzdWIiOiJ1c2VyIHRva2VuIn0.uRKQSatube3XaNG3SkodvskXJuS1ei9IPrVD8vJW5QQ"

	j := jwtToken.NewJWT()
	//解析
	tokenData, err := j.ParseToken(token)
	if err != nil {
		response.Failed(c, code.Failed, "Token parse failed")
		return
	}
	response.Success(c, tokenData)
	return
}
