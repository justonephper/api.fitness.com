package router

import (
	"fitness/app/controller/api/blogCategory"
	"github.com/gin-gonic/gin"
)

func InitBlogCategoryRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("api/v1")
	{
		//blog category 相关接口
		ApiRouter.GET("blogCategories", blogCategory.Index)
		ApiRouter.GET("blogCategories/:id", blogCategory.Show)
		ApiRouter.POST("blogCategories", blogCategory.Add)
		ApiRouter.PUT("blogCategories/:id", blogCategory.Update)
		ApiRouter.DELETE("blogCategories/:id", blogCategory.Destroy)
	}
}
