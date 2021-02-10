package auth

import (
	"fitness/bean/requestParams"
	"fitness/pkg/util/response"
	"github.com/gin-gonic/gin"
)

//login
func Login(c *gin.Context) {
	var loginParam requestParams.LoginParam
	//var header bean.Header
	if err := c.ShouldBind(&loginParam); err != nil {
		//参数校验失败统一处理函数
		response.UniqueFailedResponse(c, err)
		return
	}

	//查询数据库
	response.Success(c, "login successful")
}

//register
func Register(c *gin.Context) {
	response.Success(c, "注册成功")
}

//logOut
func LogOut(c *gin.Context) {
	response.Success(c, "退出成功")
	return
}
