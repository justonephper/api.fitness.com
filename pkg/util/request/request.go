package request

import (
	"fitness/global"
)

// Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page" binding:"required,"`
	PageSize int    `json:"page_size" form:"page_size" binding:"required,"`
	Q        string `json:"q" form:"q"`
	OrderKey string `json:"order_key" form:"order_key"`
	Desc     bool   `json:"desc" form:"desc"`
}

//获取header头中语言设置
func GetLocale() (locale string) {
	header := global.C.Request.Header
	if lang, ok := header["Api-Lang"]; ok {
		locale = lang[0]
	} else {
		locale = "en"
	}
	return locale
}
