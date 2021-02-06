package auth

import (
	"fitness/bean/requestParams"
	"fitness/pkg/util/response"
	"fmt"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

//login
func Login(c *gin.Context) {
	var loginParam requestParams.LoginParam
	//var header bean.Header
	if err := c.ShouldBind(&loginParam); err != nil {
		//参数校验失败统一处理函数
		response.UniqueFailedResponse(err)
		return
	}

	//查询数据库
	response.Success("login successful")
}

//测试异常的使用
func TestPanic(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			response.CheckRequestFailed(err)
		}
	}()

	fmt.Println("start")
	panic("the type is string!")
}

//测试响应
func TestResponse(c *gin.Context) {
	response.CheckRequestFailed("this is a string")

	msgMap := map[string]interface{}{
		"code": 10000,
		"data": nil,
		"msg": map[string]string{
			"name":     "haoge",
			"password": "111111",
		},
	}

	//dataMap := response.CheckRequestFailed(msgMap)
	//fmt.Println(dataMap)

	response.CheckRequestFailed(msgMap)
	return
}

//异步任务逻辑
func asyncJob() {
	//异步发送邮件
	//task.SendRegisterEmail(nil)
}

//register
func Register(c *gin.Context) {
	response.Success("注册成功")
}

//logOut
func LogOut(c *gin.Context) {
	response.Success("退出成功")
	return
}

func Logs(c *gin.Context) {
	logs.WithFields(logs.Fields{
		"animal": "dog",
	}).Info("大黄")
}
