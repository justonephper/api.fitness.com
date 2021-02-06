package router

import (
	"fitness/app/controller/api/blog"
	"github.com/gin-gonic/gin"
)

func InitBlogRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api/v1")
	{
		//blog相关接口
		ApiRouter.GET("blogs", blog.Index)
		ApiRouter.GET("blogs/:id", blog.Show)
		ApiRouter.POST("blogs", blog.Add)
		ApiRouter.PUT("blogs/:id", blog.Update)
		ApiRouter.DELETE("blogs/:id", blog.Destroy)
	}
}
