package translator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// 定义一个全局翻译器T
var uni *ut.UniversalTranslator
var v *validator.Validate
var ok bool

func init() {
	//初始化翻译器
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器
	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni = ut.New(enT, zhT, enT)

	// 修改gin框架中的Validator引擎属性，实现自定制
	v, ok = binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func GetTranslator(locale string) (trans ut.Translator) {
	//获取翻译器
	trans, ok = uni.GetTranslator(locale)
	if !ok {
		locale = "en"
		trans, _ = uni.GetTranslator(locale)
	}

	// 注册翻译器
	switch locale {
	case "en":
		_ = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		_ = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		_ = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	return
}
