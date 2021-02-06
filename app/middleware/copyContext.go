package middleware

import (
	"fitness/global"
	"github.com/gin-gonic/gin"
)

//拷贝请求的上线文
func CopyContext()  gin.HandlerFunc {
	return func(c *gin.Context) {
		//global.C = c.Copy()
		global.C = c
		c.Next()
	}
}
