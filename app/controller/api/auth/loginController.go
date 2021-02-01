package auth

import (
	"api.fitness.com/app/helper/response"
	"api.fitness.com/app/service/auth"
	"api.fitness.com/bean/request"
	"fmt"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"net/http"
)

//login
func Login(c *gin.Context) {
	var loginParam request.LoginParam
	//var header bean.Header
	if err := c.ShouldBind(&loginParam); err != nil {
		//参数校验失败统一处理函数
		c.JSON(http.StatusOK, response.UniqueFailedResponse(c.Copy(), err))
		return
	}

	//查询数据库
	auth.Login(loginParam)
	c.JSON(http.StatusOK, response.Success("login successful"))
}

//测试异常的使用
func TestPanic(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusOK, response.CheckRequestFailed(err))
		}
	}()

	fmt.Println("start")
	panic("the type is string!")
}

//异步任务逻辑
func asyncJob() {
	//异步发送邮件
	//job.SendRegisterEmail(nil)
}

//register
func Register(c *gin.Context) {
	c.JSON(http.StatusOK,response.Success("注册成功"))
}

//logOut
func LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("退出成功"))
	return
}

func Logs(c *gin.Context)  {
	logs.WithFields(logs.Fields{
		"animal":"dog",
	}).Info("大黄")
}
