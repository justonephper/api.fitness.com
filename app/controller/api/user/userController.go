package user

import (
	"fitness/bean/models"
	"fitness/global"
	"fitness/pkg/util/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 * @Desc:用户列表
 * @Author:haoge
 * @Date:2021/1/29 17:58
 **/
func Index(c *gin.Context) {
	//c.JSON(http.StatusOK,responseParams.Success("success"))

	//删除表结构
	//global.DB.DropTableIfExists(&models.Users{})

	//自动迁移表结构
	global.DB.AutoMigrate(&models.Users{})

	//添加数据

	response.Success("migrate successful")
}

/**
 * @Desc:博客操作
 * @Author:haoge
 * @Date:2021/2/2 17:31
 **/
func Blog(c *gin.Context) {

	//global.DB.DropTableIfExists(models.Blog{})
	//生成表
	global.DB.AutoMigrate(models.Blog{})

	blog := models.Blog{
		Name:    "第一篇博客222",
		Title:   "这是title",
		Content: "这是 desc",
	}
	//创建数据  自动填充
	res := global.DB.NewRecord(blog)
	fmt.Println("res:%v", res)
	global.DB.Create(&blog)
	res = global.DB.NewRecord(blog)
	fmt.Println("res:%v", res)
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
