package blogCategory

import (
	"api.fitness.com/app/helper/request"
	"api.fitness.com/app/helper/response"
	"api.fitness.com/bean/models"
	"api.fitness.com/bean/requestParams"
	"api.fitness.com/global"
	"api.fitness.com/pkg/code"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var blogCategory *models.BlogCategory

/**
 * @Desc:添加博客分类
 * @Author:haoge
 * @Date:`2021/2/4` 16:32
 **/
func Add(c *gin.Context) {
	var params requestParams.BlogCategoryAddParams

	//validate request params
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusOK, response.UniqueFailedResponse(c.Copy(), err))
		return
	}

	//迁移表
	global.DB.AutoMigrate(models.BlogCategory{})

	blogCategory = models.BlogCategoryNew()
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
	id := c.Param("id")

	blogCategory = models.BlogCategoryNew()
	if !blogCategory.Find(id) {
		c.JSON(http.StatusOK, response.Failed(code.BlogCategoryNotExists, nil))
		return
	}
	c.JSON(http.StatusOK, response.Success(blogCategory))
	return
}

/**
 * @Desc:修改博客分类
 * @Author:haoge
 * @Date:2021/2/4 16:33
 **/
func Update(c *gin.Context) {
	var params requestParams.BlogCategoryUpdateParams
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusOK, response.UniqueFailedResponse(c.Copy(), err))
		return
	}
	id := c.Param("id")

	blogCategory = models.BlogCategoryNew()
	//查询是否存在
	if !blogCategory.Find(id) {
		c.JSON(http.StatusOK, response.Failed(code.BlogCategoryNotExists, nil))
		return
	}

	updateData := map[string]interface{}{
		"name": params.Name,
		"desc": params.Desc,
	}
	if !blogCategory.Update(updateData) {
		c.JSON(http.StatusOK, response.Failed(code.BlogCategoryUpdateFailed, nil))
		return
	}
	c.JSON(http.StatusOK, response.Success(nil))
	return
}

/**
 * @Desc:删除博客分类
 * @Author:haoge
 * @Date:2021/2/4 16:33
 **/
func Destroy(c *gin.Context) {
	id := c.Param("id")

	blogCategory = models.BlogCategoryNew()
	//查询数据是否存在
	if !blogCategory.Find(id) {
		c.JSON(http.StatusOK, response.Failed(code.BlogCategoryNotExists, nil))
		return
	}

	//删除数据
	if !blogCategory.Delete() {
		c.JSON(http.StatusOK, response.Failed(code.BlogCategoryDeleteFailed, nil))
		return
	}
	c.JSON(http.StatusOK, response.Success(nil))
	return
}

/**
 * @Desc:博客分类列表
 * @Author:haoge
 * @Date:2021/2/4 16:34
 **/
func Index(c *gin.Context) {
	var pageInfo request.PageInfo
	//param validate
	if err := c.Bind(&pageInfo); err != nil {
		c.JSON(http.StatusOK, response.UniqueFailedResponse(c.Copy(), err))
		return
	}

	fmt.Printf("%#v", pageInfo)

	//实例化对象
	blogCategory = models.BlogCategoryNew()
	list, total, err := blogCategory.Index(pageInfo)
	if err != nil {
		c.JSON(http.StatusOK, response.PageListNoData(pageInfo))
		return
	}
	c.JSON(http.StatusOK, response.PageListData(list, total, pageInfo))
	return
}
