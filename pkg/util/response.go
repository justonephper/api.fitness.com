package util

import (
	"github.com/gin-gonic/gin"
)

func Success(msg string, data interface{}) gin.H {
	if msg == "" {
		msg = "请求成功"
	}
	return gin.H{
		"code":   200,
		"status": "success",
		"msg":    msg,
		"data":   data,
	}
}

func Fail(msg string) gin.H {
	if msg == "" {
		msg = "请求失败"
	}
	return gin.H{
		"code":   400,
		"status": "fail",
		"msg":    msg,
		"data":   nil,
	}
}
