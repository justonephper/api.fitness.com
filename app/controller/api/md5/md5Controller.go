package md5

import (
	"fitness/pkg/util/md5"
	"fitness/pkg/util/response"
	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string
}

var str string = "haoge"

//生成md5加密串
func Md5(c *gin.Context) {
	t := md5.NewMd5()
	md5Str := t.Encrypt([]byte(str))
	response.Success(md5Str)
	return
}

//检验密码是否正确
func Verify(c *gin.Context) {
	t := md5.NewMd5()
	md5str := t.Encrypt([]byte(str))
	res := t.Verify([]byte(str), md5str)
	response.Success(res)
	return
}
