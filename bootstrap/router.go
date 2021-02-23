package bootstrap

import (
	"fitness/app/middleware"
	"fitness/global"
	"fitness/pkg/code"
	"fitness/pkg/util/response"
	"fitness/router"
	"github.com/gin-gonic/gin"
	"time"
)

//404错误处理
func HandleNotFound(c *gin.Context) {
	response.Failed(c,code.RequestUrlNotFound, "requestParams url not exists!")
	return
}

//注册错误路由处理函数
func injectGlobalDefaultUrl(router *gin.Engine) {
	//错误处理路由
	router.NoMethod(HandleNotFound)
	router.NoRoute(HandleNotFound)
}

//注册全局中间件
func injectGlobalMiddleware(router *gin.Engine) {
	router.Use(middleware.Cors())
}

func InitRouter() *gin.Engine {
	Router := gin.Default()
	//注入全局中间件
	injectGlobalMiddleware(Router)

	//注入错误处理路由
	injectGlobalDefaultUrl(Router)

	//兜底路由
	Router.GET("/", func(c *gin.Context) {
		response_str := "欢迎访问登云Api! " + "北京时间:" + time.Now().Format(global.TimeFormate)
		response.Success(c,response_str)
	})

	//未登陆状态下的路由组
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	//登陆状态下的路由组
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitUserRouter(PrivateGroup)
		router.InitBlogCategoryRouter(PrivateGroup)
		router.InitBlogRouter(PrivateGroup)
	}
	return Router
}
