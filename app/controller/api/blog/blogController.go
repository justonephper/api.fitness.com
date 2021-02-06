package blog

import (
	"fitness/bean/models"
	"fitness/bean/requestParams"
	"fitness/global"
	"fitness/pkg/code"
	"fitness/pkg/util/request"
	"fitness/pkg/util/response"
	"github.com/gin-gonic/gin"
)

var blog *models.Blog

/**
 * @Desc:添加博客
 * @Author:haoge
 * @Date:2021/2/2 18:16
 **/
func Add(c *gin.Context) {
	var params requestParams.BlogAddParams
	//parse requestParams params
	if err := c.ShouldBind(&params); err != nil {
		//参数校验失败统一处理函数
		response.UniqueFailedResponse(err)
		return
	}

	//执行迁移文件
	global.DB.AutoMigrate(&models.Blog{})

	//查询是否存在日志记录（主键为空返回`true`）
	//res := global.DB.NewRecord(insert)

	//blog := &models.Blog{
	//	Name:    params.Name,
	//	Title:   params.Title,
	//	Content: params.Content,
	//}

	blog = models.NewBlog()
	blog.Name = params.Name
	blog.Title = params.Title
	blog.Content = params.Content

	if ok := blog.Create(); !ok {
		response.Failed(code.BlogAddFailed, nil)
		return
	}
	response.Success(blog)
	return
}

/**
 * @Desc:编辑博客
 * @Author:haoge
 * @Date:2021/2/2 18:24
 **/
func Update(c *gin.Context) {
	var params requestParams.BlogUpdateParams
	//parse requestParams params
	if err := c.Bind(&params); err != nil {
		response.UniqueFailedResponse(err)
		return
	}

	//get requestParams id
	id := c.Param("id")

	//query db-data
	blog = models.NewBlog()
	ok := blog.Find(id)
	if !ok {
		response.Failed(code.BlogNotExists, nil)
		return
	}

	//update db-data
	updateData := map[string]interface{}{
		"name":    params.Name,
		"title":   params.Title,
		"content": params.Content,
	}

	if ok = blog.Update(updateData); !ok {
		response.Failed(code.BlogUpdateFailed, nil)
		return
	}

	response.Success(blog)
	return
}

/**
 * @Desc:删除博客
 * @Author:haoge
 * @Date:2021/2/2 18:25
 **/
func Destroy(c *gin.Context) {
	//get requestParams id (string)
	id := c.Param("id")

	//quert db-data
	blog = models.NewBlog()
	ok := blog.Find(id)
	if !ok {
		response.Failed(code.BlogNotExists, nil)
		return
	}

	//delete db-data
	if !blog.Delete() {
		response.Failed(code.BlogDeleteFailed, nil)
		return
	}
	response.Success(nil)
	return
}

/**
 * @Desc:博客详情
 * @Author:haoge
 * @Date:2021/2/2 18:26
 **/
func Show(c *gin.Context) {
	id := c.Param("id")

	blog = models.NewBlog()
	ok := blog.Find(id)
	if !ok {
		response.Failed(code.BlogNotExists, nil)
		return
	}
	response.Success(blog)
	return
}

/**
 * @Desc:博客列表
 * @Author:haoge
 * @Date:2021/2/2 18:25
 **/
func Index(c *gin.Context) {
	var pageInfo request.PageInfo

	//params validate
	if err := c.Bind(&pageInfo); err != nil {
		response.UniqueFailedResponse(err)
		return
	}

	blog = models.NewBlog()
	list, total, err := blog.Index(pageInfo)
	if err != nil {
		response.PageListNoData(pageInfo)
		return
	}

	response.PageListData(list, total, pageInfo)
	return
}
