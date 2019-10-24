package api

import (
	"daily.com/models"
	"daily.com/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	ID    int    `json:"staff_id"`
	Token string `json:"token"`
}

func ApiRequestCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Copy()
		var params Request
		if err := ctx.BindJSON(&params); err != nil {
			c.JSON(http.StatusOK, util.Fail("请求失败"))
			return
		}
		requestURI := c.Request.RequestURI
		if requestURI != "/api/v1/login" && requestURI != "/api/v1/logout" {
			//参数检测
			if ok := checkParams(params); !ok {
				c.JSON(http.StatusOK, util.Fail("必要参数不足!"))
				c.Abort()
				return
			}
			if ok := checkPower(params); !ok {
				c.JSON(http.StatusOK, util.Fail("您的账号已在其他设备登录，请重新登录!"))
				c.Abort()
				return
			}
		}
		// before request
		c.Next()
	}
}

//检验参数
func checkParams(params Request) bool {
	valid := validation.Validation{}

	valid.Required(params.ID, "id")
	valid.Required(params.Token, "token")

	if valid.HasErrors() {
		return false
	}
	return true
}

//鉴权
func checkPower(params Request) bool {
	staff, ok := models.GetStaffById(params.ID)
	if !ok {
		return false
	}
	if staff.Token != params.Token {
		return false
	}
	return true
}
