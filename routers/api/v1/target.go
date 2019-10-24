package v1

import (
	"daily.com/models"
	"daily.com/pkg/e"
	"daily.com/pkg/setting"
	"daily.com/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取目标列表
func Targets(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.GetTargets(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTargetTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func OneTarget(c *gin.Context) {
}

//目标创建
func AddTarget(c *gin.Context) {
}

//编辑目标
func EditTarget(c *gin.Context) {
}

//删除目标
func DeleteTarget(c *gin.Context) {
}
