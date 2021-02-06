package router

import (
	"fitness/app/controller/api/Log"
	"fitness/app/controller/api/auth"
	"fitness/app/controller/api/md5"
	"fitness/app/controller/api/task"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("/api/v1")
	{
		//登录 注册 退出登录管理
		BaseRouter.GET("testPanic", auth.TestPanic)
		BaseRouter.POST("login", auth.Login)
		BaseRouter.POST("register", auth.Register)
		BaseRouter.GET("logout", auth.LogOut)
		BaseRouter.GET("logs", auth.Logs)
		BaseRouter.GET("response", auth.TestResponse)

		//token
		BaseRouter.GET("genToken", auth.GenToken)
		BaseRouter.GET("parseToken", auth.ParseToken)

		//task
		BaseRouter.GET("sendEmail", task.SendEmail)
		BaseRouter.GET("sendSms", task.SendSms)

		//md5
		BaseRouter.GET("md5", md5.Md5)
		BaseRouter.GET("md5Verify", md5.Verify)

		//log接口
		BaseRouter.GET("logTest", Log.LogTest)
	}
	return BaseRouter
}
