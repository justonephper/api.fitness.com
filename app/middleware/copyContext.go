package middleware

import (
	"fitness/global"
	"github.com/gin-gonic/gin"
)

//拷贝请求的上下文
func CopyContext()  gin.HandlerFunc {
	return func(c *gin.Context) {
		//global.C = c.Copy()
		//将上下文放入全局请求变量，避免上下文在方法中传递
		global.C = c
		c.Next()
	}
}
