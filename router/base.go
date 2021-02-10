package router

import (
	"fitness/app/controller/api/auth"
	"fitness/app/controller/api/index"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("/api/v1")
	{
		//默认控制器 (调试使用)
		BaseRouter.GET("testPanic", index.TestPanic)
		BaseRouter.GET("response", index.TestResponse)

		//登录 注册 退出登录管理
		BaseRouter.POST("login", auth.Login)
		BaseRouter.POST("register", auth.Register)
		BaseRouter.GET("logout", auth.LogOut)
	}
	return BaseRouter
}
