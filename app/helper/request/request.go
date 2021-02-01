package request

import "github.com/gin-gonic/gin"

//获取header头中语言设置
func GetLocale(c *gin.Context) (locale string) {
	header := c.Request.Header
	if lang, ok := header["Api-Lang"]; ok {
		locale = lang[0]
	} else {
		locale = "en"
	}
	return locale
}


