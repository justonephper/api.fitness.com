package response

import (
	"fitness/pkg/code"
	"fitness/pkg/util/request"
	"fitness/pkg/util/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type PageList struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  interface{} `json:"msg"`
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func Result(c *gin.Context, code int, data interface{}, msg interface{}) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

//校验参数失败回调     string | map[string]string
func CheckRequestFailed(c *gin.Context, msg interface{}) {
	//数据处理
	switch msgs := msg.(type) {
	case string:
		Result(c, code.BadRequestParams, nil, msgs)
	case map[string]string:
		Result(c, code.BadRequestParams, nil, removeTopStruct(msgs))
	default:
		Result(c, code.BadRequestParams, nil, code.LogicCodeText(code.BadRequestParams))
	}
}

//普通失败回调
func Failed(c *gin.Context, codeNum int, data interface{}) {
	if data != nil {
		Result(c, codeNum, nil, data)
	} else {
		Result(c, codeNum, data, code.LogicCodeText(codeNum))
	}
}

//成功回调
func Success(c *gin.Context, data interface{}) {
	Result(c, code.Success, data, code.LogicCodeText(code.Success))
}

//请求的失败总方法
func UniqueFailedResponse(c *gin.Context, err error) {
	//检测是否是参数校验错误
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		Failed(c, code.BadRequestParams, err.Error())
		return
	}
	// validator.ValidationErrors类型错误则进行翻译
	locale := request.GetLocale()
	trans := translator.GetTranslator(locale)
	CheckRequestFailed(c, errs.Translate(trans))
}

//分页数据
func PageListData(c *gin.Context, list interface{}, total int64, info request.PageInfo) {
	data := PageList{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}
	Result(c, code.Success, data, code.LogicCodeText(code.Success))
}

//分页无数据
func PageListNoData(c *gin.Context, info request.PageInfo) {
	data := PageList{
		List:     nil,
		Total:    0,
		Page:     info.Page,
		PageSize: info.PageSize,
	}
	Result(c, code.Success, data, code.LogicCodeText(code.Success))
}
