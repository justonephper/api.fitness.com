package index

import (
	"fitness/pkg/util/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

//测试异常的使用
func TestPanic(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			response.CheckRequestFailed(c, err)
		}
	}()

	fmt.Println("start")
	panic("the type is string!")
}

//测试响应
func TestResponse(c *gin.Context) {
	response.CheckRequestFailed(c, "this is a string")

	msgMap := map[string]interface{}{
		"code": 10000,
		"data": nil,
		"msg": map[string]string{
			"name":     "haoge",
			"password": "111111",
		},
	}

	//dataMap := response.CheckRequestFailed(c,msgMap)
	//fmt.Println(dataMap)

	response.CheckRequestFailed(c, msgMap)
	return
}
