package request

import "github.com/gin-gonic/gin"

// Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"page_size" form:"page_size" binding:"required"`
	Q        string `json:"q" form:"q"`
	Order    string `json:"order"`
	Desc     bool   `json:"desc"`
}

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
