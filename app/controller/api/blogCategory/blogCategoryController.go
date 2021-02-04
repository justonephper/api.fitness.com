package blogCategory

import (
	"api.fitness.com/app/helper/response"
	"api.fitness.com/bean/models"
	"api.fitness.com/bean/requestParams"
	"api.fitness.com/pkg/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

var err error
var blogCategory *models.BlogCategory

func init() {
	blogCategory = models.BlogCategoryNew()
}

/**
 * @Desc:添加博客分类
 * @Author:haoge
 * @Date:`2021/2/4` 16:32
 **/
func Add(c *gin.Context) {
	var params requestParams.BlogCategoryAddParams

	//validate request params
	if err = c.Bind(&params); err != nil {
		c.JSON(http.StatusOK, response.UniqueFailedResponse(c.Copy(), err))
		return
	}

	blogCategory.Name = params.Name
	blogCategory.Desc = params.Desc
	ok := blogCategory.Add()
	if !ok {
		c.JSON(http.StatusOK, response.Failed(code.BlogCategoryAddFailed, nil))
		return
	}
	c.JSON(http.StatusOK, response.Success(blogCategory))
	return
}

/**
 * @Desc:博客详情
 * @Author:haoge
 * @Date:2021/2/4 16:33
 **/
func Show(c *gin.Context) {

}

/**
 * @Desc:修改博客分类
 * @Author:haoge
 * @Date:2021/2/4 16:33
 **/
func Update(c *gin.Context) {

}

/**
 * @Desc:删除博客分类
 * @Author:haoge
 * @Date:2021/2/4 16:33
 **/
func Destroy(c *gin.Context) {

}

/**
 * @Desc:博客分类列表
 * @Author:haoge
 * @Date:2021/2/4 16:34
 **/
func Index(c *gin.Context) {

}
