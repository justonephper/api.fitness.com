package v1

import (
	"daily.com/models"
	"daily.com/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var login struct {
		StaffNum string `json:"staff_num"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusOK, util.Fail(err.Error()))
		return
	}

	valid := validation.Validation{}
	valid.Required(login.StaffNum, "staff_num").Message("账号不能为空")
	valid.Required(login.Password, "password").Message("密码不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	}

	staff, ok := models.GetStaffByStaffNum(login.StaffNum)
	if !ok {
		c.JSON(http.StatusOK, util.Fail("用户信息不存在"))
		return
	}
	if ok := models.PasswordVerify(login.Password, staff.LoginPassword); ok {
		c.JSON(http.StatusOK, util.Success("登录成功", staff))
		return
	} else {
		c.JSON(http.StatusOK, util.Fail("账号或密码错误"))
		return
	}
}

//退出登录
func Logout(c *gin.Context) {
	var params struct {
		StaffId int    `json:"staff_id" form:"staff_id" binding:"required"`
		Token   string `json:"token" form:"token" binding:"required"`
	}

	if c.BindJSON(&params) != nil {
		c.JSON(http.StatusOK, util.Fail("请求参数不合法"))
		return
	}

	staff, ok := models.GetStaffById(params.StaffId)
	if !ok {
		c.JSON(http.StatusOK, util.Fail("用户信息不存在"))
		return
	}
	//更新用户token
	if ok := staff.UpdateToken(); !ok {
		c.JSON(http.StatusOK, util.Fail("退出失败!"))
		return
	}
	c.JSON(http.StatusOK, util.Success("退出成功", nil))
	return
}
