package Log

import (
	"fitness/app/helper/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func LogTest(c *gin.Context) {
	//c.JSON(http.StatusOK,responseParams.Success("log test!"))

	////设置日志字段，使用info报错级别
	//logrus.WithFields(logrus.Fields{
	//	"name":"haoge",
	//	"age":24,
	//	"address":"beijingshi chongwenmen",
	//}).Info("支付失败:")

	//设置日志输出位置
	//log.Out = os.Stdout

	log.WithFields(logrus.Fields{
		"name":    "haoge",
		"age":     24,
		"address": "beijingshi chongwenmen",
	}).Error("系统异常：")

	response.Success(c, "handle ok")
}
