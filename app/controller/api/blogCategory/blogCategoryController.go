package blogCategory

import (
	"fitness/app/helper/request"
	"fitness/app/helper/response"
	"fitness/bean/models"
	"fitness/bean/requestParams"
	"fitness/global"
	"fitness/pkg/code"
	"github.com/gin-gonic/gin"
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
		response.UniqueFailedResponse(c, err)
		return
	}

	//迁移表
	global.DB.AutoMigrate(models.BlogCategory{})

	blogCategory = models.NewBlogCategory()
	blogCategory.Name = params.Name
	blogCategory.Desc = params.Desc
	ok := blogCategory.Add()
	if !ok {
		response.Failed(c, code.BlogCategoryAddFailed, nil)
		return
	}
	response.Success(c, blogCategory)
	return
}

/**
 * @Desc:博客详情
 * @Author:haoge
 * @Date:2021/2/4 16:33
 **/
func Show(c *gin.Context) {
	id := c.Param("id")

	blogCategory = models.NewBlogCategory()
	if !blogCategory.Find(id) {
		response.Failed(c, code.BlogCategoryNotExists, nil)
		return
	}
	response.Success(c, blogCategory)
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
		response.UniqueFailedResponse(c, err)
		return
	}
	id := c.Param("id")

	blogCategory = models.NewBlogCategory()
	//查询是否存在
	if !blogCategory.Find(id) {
		response.Failed(c, code.BlogCategoryNotExists, nil)
		return
	}

	updateData := map[string]interface{}{
		"name": params.Name,
		"desc": params.Desc,
	}
	if !blogCategory.Update(updateData) {
		response.Failed(c, code.BlogCategoryUpdateFailed, nil)
		return
	}
	response.Success(c, nil)
	return
}

/**
 * @Desc:删除博客分类
 * @Author:haoge
 * @Date:2021/2/4 16:33
 **/
func Destroy(c *gin.Context) {
	id := c.Param("id")

	blogCategory = models.NewBlogCategory()
	//查询数据是否存在
	if !blogCategory.Find(id) {
		response.Failed(c, code.BlogCategoryNotExists, nil)
		return
	}

	//删除数据
	if !blogCategory.Delete() {
		response.Failed(c, code.BlogCategoryDeleteFailed, nil)
		return
	}
	response.Success(c, nil)
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
		response.UniqueFailedResponse(c, err)
		return
	}

	//实例化对象
	blogCategory = models.NewBlogCategory()
	list, total, err := blogCategory.Index(pageInfo)
	if err != nil {
		response.PageListNoData(c, pageInfo)
		return
	}
	response.PageListData(c, list, total, pageInfo)
	return
}
