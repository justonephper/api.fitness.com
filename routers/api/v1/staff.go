package v1

import (
	"daily.com/models"
	"daily.com/pkg/setting"
	"encoding/json"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"

	//"daily.com/models"
	//"daily.com/pkg/setting"
	"daily.com/pkg/util"
	//"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
)

//获取多个用户
func GetStaffs(c *gin.Context) {
	where := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	var staff_status int = -1
	if arg := c.DefaultPostForm("staff_status", "all"); arg != "all" {
		staff_status = com.StrTo(arg).MustInt()
		where["staff_status"] = staff_status
		valid.Range(staff_status, 1, 2, "staff_status").Message("用户状态错误")
	}

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	}
	data["list"] = models.GetStaffs(util.GetPage(c), setting.PageSize, where)
	data["page"] = c.PostForm("page")
	data["count"] = models.GetStaffTotal(where)
	c.JSON(http.StatusOK, util.Success("请求成功", data))
}

//获取单个用户
func GetStaff(c *gin.Context) {
	valid := validation.Validation{}

	var StaffId int = 0
	if arg := c.PostForm("staff_id"); arg != "" {
		StaffId = com.StrTo(arg).MustInt()
	}
	valid.Min(StaffId, 1, "staff_id").Message("staffId必须大于0")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	}
	if data, ok := models.GetStaffById(StaffId); !ok {
		c.JSON(http.StatusOK, util.Fail("用户信息不存在"))
		return
	} else {
		response := make(map[string]interface{})
		by_data, _ := json.Marshal(data)
		json.Unmarshal(by_data, &response)
		delete(response, "login_password")
		c.JSON(http.StatusOK, util.Success("请求成功", response))
	}

}

//添加用户
func AddStaff(c *gin.Context) {

}

//修改用户
func UpdateStaff(c *gin.Context) {

}

//删除用户
func DeleteStaff(c *gin.Context) {

}
