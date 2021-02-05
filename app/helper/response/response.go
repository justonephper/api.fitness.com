package response

import (
	"api.fitness.com/app/helper/request"
	"api.fitness.com/pkg/code"
	"api.fitness.com/pkg/util/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

type PageList struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

//校验参数失败回调     string | map[string]string
func CheckRequestFailed(msg interface{}) gin.H {
	//数据处理
	switch msgs := msg.(type) {
	case string:
		return gin.H{"code": code.BadRequestParams, "data": nil, "msg": msgs}
	case map[string]string:
		return gin.H{"code": code.BadRequestParams, "data": nil, "msg": removeTopStruct(msgs)}
	default:
		return gin.H{"code": code.BadRequestParams, "data": nil, "msg": code.LogicCodeText(code.BadRequestParams)}
	}
}

//普通失败回调
func Failed(codeNum int, data interface{}) gin.H {
	return gin.H{"code": codeNum, "data": data, "msg": code.LogicCodeText(codeNum)}
}

//成功回调
func Success(data interface{}) gin.H {
	return gin.H{"code": code.Success, "data": data, "msg": "success"}
}

//请求的失败总方法
func UniqueFailedResponse(c *gin.Context, err error) gin.H {
	//检测是否是参数校验错误
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		return Failed(code.BadRequestParams, err.Error())
	}
	// validator.ValidationErrors类型错误则进行翻译
	locale := request.GetLocale(c)
	trans := translator.GetTranslator(locale)
	return CheckRequestFailed(errs.Translate(trans))
}

//分页数据
func PageListData(list interface{}, total int64, info request.PageInfo) PageList {
	return PageList{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}
}

//分页无数据
func PageListNoData(info request.PageInfo) PageList {
	return PageList{
		List:     nil,
		Total:    0,
		Page:     info.Page,
		PageSize: info.PageSize,
	}
}
