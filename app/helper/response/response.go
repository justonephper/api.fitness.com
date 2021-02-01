package response

import (
	"api.fitness.com/app/helper/request"
	"api.fitness.com/pkg/code"
	"api.fitness.com/pkg/util/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

//校验参数失败回调
func CheckRequestFailed(msg interface{}) gin.H {
	//数据处理
	typeOf := reflect.TypeOf(msg)

	switch typeOf.Kind() {
	case reflect.String:
		return gin.H{"code": code.BadRequestParams, "data": nil, "msg": msg.(string)}
	case reflect.Map:
		msgs := msg.(validator.ValidationErrorsTranslations)
		return gin.H{"code": code.BadRequestParams, "data": nil, "msg": removeTopStruct(msgs)}
	default:
		panic("Unsupported data")
	}
}

//普通失败回调
func Failed(codeNum int, data interface{}) gin.H {
	return gin.H{"code": codeNum, "data": data, "msg": code.LogicCodeText(codeNum)}
}

//成功回调
func Success(data interface{}) gin.H {
	return gin.H{"code": 10000, "data": data, "msg": "success"}
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
