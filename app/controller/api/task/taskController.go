package task

import (
	"fitness/app/job"
	"fitness/pkg/util/response"
	"github.com/gin-gonic/gin"
)

//发送邮件
func SendEmail(c *gin.Context) {
	j := job.NewEmailSender()
	j.Handle(nil)
	response.Success("send email successful")
	return
}

//发送短信
func SendSms(c *gin.Context) {
	j := job.SmsSender{}
	j.Handle(nil)
	response.Success("send sms successful")
	return
}
