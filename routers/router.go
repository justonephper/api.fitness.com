package routers

import (
	"fitness/app/controller/api/Log"
	"fitness/app/controller/api/auth"
	"fitness/app/controller/api/blog"
	"fitness/app/controller/api/blogCategory"
	"fitness/app/controller/api/user"
	"fitness/app/helper/response"
	"fitness/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//404错误处理
func HandleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, response.Failed(10004, "requestParams url not exists!"))
	return
}

func Init() *gin.Engine {
	router := gin.Default()
	//错误处理路由
	router.NoMethod(HandleNotFound)
	router.NoRoute(HandleNotFound)

	//兜底路由
	router.GET("/", func(c *gin.Context) {
		response_str := "欢迎访问登云Api! " + "北京时间:" + time.Now().Format(global.TimeFormate)
		c.JSON(http.StatusOK, response.Success(response_str))
		//c.String(http.StatusOK,response_str)
	})

	//登录 注册 退出登录管理
	router.GET("testPanic", auth.TestPanic)
	router.POST("login", auth.Login)
	router.POST("register", auth.Register)
	router.GET("logout", auth.LogOut)
	router.GET("logs", auth.Logs)
	router.GET("response", auth.TestResponse)

	router.GET("genToken", auth.GenToken)
	router.GET("parseToken", auth.ParseToken)

	//db-test
	router.GET("users", user.Index)
	router.GET("users/:id", user.Show)
	router.POST("users", user.Add)
	router.PUT("users/:id", user.Update)
	router.DELETE("users/:id", user.Destroy)

	//log接口
	router.GET("logTest", Log.LogTest)

	//blog category 相关接口
	router.GET("blogCategories", blogCategory.Index)
	router.GET("blogCategories/:id", blogCategory.Show)
	router.POST("blogCategories", blogCategory.Add)
	router.PUT("blogCategories/:id", blogCategory.Update)
	router.DELETE("blogCategories/:id", blogCategory.Destroy)

	//blog相关接口
	router.GET("blogs", blog.Index)
	router.GET("blogs/:id", blog.Show)
	router.POST("blogs", blog.Add)
	router.PUT("blogs/:id", blog.Update)
	router.DELETE("blogs/:id", blog.Destroy)



	return router
}

//func setupRouter() *gin.Engine {
//
//	//// 初始化 Gin 框架默认实例，该实例包含了路由、中间件以及配置信息
//	//r := gin.Default()
//	//
//	//// Ping 测试路由
//	////r.GET("/ping", func(c *gin.Context) {
//	////	c.String(http.StatusOK, "pong")
//	////})
//	//
//	//// 获取用户数据路由
//	//r.GET("/user/:name", func(c *gin.Context) {
//	//	user := c.Params.ByName("name")
//	//	value, ok := db[user]
//	//	if ok {
//	//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
//	//	} else {
//	//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
//	//	}
//	//})
//	//
//	//// 需要 HTTP 基本授权认证的子路由群组设置
//	//authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
//	//	"foo":  "bar", // 用户名:foo 密码:bar
//	//	"manu": "123", // 用户名:manu 密码:123
//	//}))
//	//
//	//// 保存用户信息路由
//	//authorized.POST("admin", func(c *gin.Context) {
//	//	user := c.MustGet(gin.AuthUserKey).(string)
//	//
//	//	// 解析并验证 JSON 格式请求数据
//	//	var json struct {
//	//		Value string `json:"value" binding:"required"`
//	//	}
//	//
//	//	if c.Bind(&json) == nil {
//	//		db[user] = json.Value
//	//		c.JSON(http.StatusOK, gin.H{"status": "ok"})
//	//	}
//	//})
//
//	return r
//}
