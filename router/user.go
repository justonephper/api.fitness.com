package router

import (
	"fitness/app/controller/api/user"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api/v1")
	{
		//db-test
		ApiRouter.GET("users", user.Index)
		ApiRouter.GET("users/:id", user.Show)
		ApiRouter.POST("users", user.Add)
		ApiRouter.PUT("users/:id", user.Update)
		ApiRouter.DELETE("users/:id", user.Destroy)
	}
}
