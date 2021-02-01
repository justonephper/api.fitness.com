package user

import (
	"api.fitness.com/app/helper/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

/**
 * @Desc:用户列表
 * @Author:haoge
 * @Date:2021/1/29 17:58
 **/
func Index(c *gin.Context) {
	log.Println("this is a log")
	c.JSON(http.StatusOK,response.Success("index"))
}

/**
 * @Desc:查询单条
 * @Author:haoge
 * @Date:2021/1/29 17:58
 **/
func Show(c *gin.Context) {

}

/**
 * @Desc:添加用户
 * @Author:haoge
 * @Date:2021/1/29 17:56
 **/
func Add(c *gin.Context) {

}

/**
 * @Desc:编辑用户
 * @Author:haoge
 * @Date:2021/1/29 17:59
 **/
func Update(c *gin.Context) {

}

/**
 * @Desc:删除用户
 * @Author:haoge
 * @Date:2021/1/29 17:59
 **/
func Destroy(c *gin.Context) {

}
