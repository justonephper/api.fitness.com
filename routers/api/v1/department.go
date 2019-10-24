package v1

import (
	"daily.com/models"
	"daily.com/pkg/setting"
	"daily.com/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取部门列表
func GetDepartments(c *gin.Context) {
	where := make(map[string]interface{})
	data := make(map[string]interface{})
	departmentStatus := com.StrTo(c.PostForm("department_status")).MustInt()
	page := com.StrTo(c.PostForm("page")).MustInt()
	valid := validation.Validation{}

	valid.Required(departmentStatus, "department_stauts").Message("department_status 数据不能为空")
	valid.Required(page, "page").Message("page 数据不能为空")
	valid.Range(departmentStatus, 1, 2, "department_status").Message("department_status 数据不合法")
	valid.Min(page, 1, "page").Message("page 数据不合法")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	} else {
		where["department_status"] = departmentStatus
	}
	data["list"] = models.GetDepartments(util.GetPage(c), setting.PageSize, where)
	data["count"] = models.GetDepartmentTotal(where)
	data["page"] = c.PostForm("page")
	c.JSON(http.StatusOK, util.Success("请求成功", data))
}

//获取部门分组列表
func DepartmentGroup(c *gin.Context) {
	where := make(map[string]interface{})
	data := make(map[string]interface{})
	departmentStatus := com.StrTo(c.PostForm("department_status")).MustInt()
	page := com.StrTo(c.PostForm("page")).MustInt()
	valid := validation.Validation{}

	valid.Required(departmentStatus, "department_stauts").Message("department_status 数据不能为空")
	valid.Required(page, "page").Message("page 数据不能为空")
	valid.Range(departmentStatus, 1, 2, "department_status").Message("department_status 数据不合法")
	valid.Min(page, 1, "page").Message("page 数据不合法")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	} else {
		where["department_status"] = departmentStatus
	}
	data["list"] = models.DepartmentGroup(where)
	data["count"] = models.GetDepartmentTotal(where)
	data["page"] = c.PostForm("page")
	c.JSON(http.StatusOK, util.Success("请求成功", data))
}

//获取部门信息
func GetDepartment(c *gin.Context) {
	departmentId := com.StrTo(c.PostForm("department_id")).MustInt()
	valid := validation.Validation{}

	valid.Required(departmentId, "department_id").Message("department_id为必填项")
	valid.Min(departmentId, 1, "department_id").Message("department_id必须大于0")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	}

	data := models.GetDepartmentById(departmentId)
	if data.ID > 0 {
		c.JSON(http.StatusOK, util.Success("请求成功!", data))
	} else {
		c.JSON(http.StatusOK, util.Success("请求成功", make(map[string]string)))
	}
}

//添加部门
func AddDepartment(c *gin.Context) {
	companyId := com.StrTo(c.PostForm("company_id")).MustInt()
	departmentName := c.PostForm("department_name")
	pid := com.StrTo(c.PostForm("pid")).MustInt()
	status := com.StrTo(c.PostForm("department_status")).MustInt()
	staffId := com.StrTo(c.PostForm("staff_id")).MustInt()
	valid := validation.Validation{}

	valid.Required(companyId, "company_id").Message("company_id 不能为空")
	valid.Required(departmentName, "department_name").Message("department_name 不能为空")
	valid.Required(pid, "pid").Message("pid 不能为空")
	valid.Required(status, "department_status").Message("department_status 不能为空")
	valid.Required(staffId, "staff_id").Message("staff_id 不能为空")
	valid.Min(companyId, 1, "company_id").Message("company_id 必须大于0")
	valid.Min(pid, 1, "pid").Message("pid 必须大于0")
	valid.Range(status, 1, 2, "department_status").Message("department_status 必须为1或2")
	valid.Min(staffId, 1, "staff_id").Message("staff_id 必须大于0")

	params := make(map[string]interface{})
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	} else {
		params["pid"] = pid
		params["department_name"] = departmentName
		params["department_status"] = status
		params["company_id"] = companyId
		params["created_by"] = staffId
	}
	if ok := models.AddDepartment(params); ok {
		c.JSON(http.StatusOK, util.Success("添加成功", make(map[string]string)))
		return
	} else {
		c.JSON(http.StatusOK, util.Success("添加失败", make(map[string]string)))
		return
	}
}

//更新部门
func UpdateDepartment(c *gin.Context) {
	companyId := com.StrTo(c.PostForm("company_id")).MustInt()
	departmentName := c.PostForm("department_name")
	pid := com.StrTo(c.PostForm("pid")).MustInt()
	status := com.StrTo(c.PostForm("department_status")).MustInt()
	staffId := com.StrTo(c.PostForm("staff_id")).MustInt()
	Id := com.StrTo(c.PostForm("id")).MustInt()
	valid := validation.Validation{}

	valid.Required(companyId, "company_id").Message("company_id 不能为空")
	valid.Required(departmentName, "department_name").Message("department_name 不能为空")
	valid.Required(pid, "pid").Message("pid 不能为空")
	valid.Required(status, "department_status").Message("department_status 不能为空")
	valid.Required(staffId, "staff_id").Message("staff_id 不能为空")
	valid.Required(Id, "id").Message("id 不能为空")
	valid.Min(companyId, 1, "company_id").Message("company_id 必须大于0")
	valid.Min(pid, 1, "pid").Message("pid 必须大于0")
	valid.Range(status, 1, 2, "department_status").Message("department_status 必须为1或2")
	valid.Min(staffId, 1, "staff_id").Message("staff_id 必须大于0")
	valid.Min(Id, 1, "id").Message("id 必须大于0")

	data := make(map[string]interface{})
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	} else {
		data["pid"] = pid
		data["department_name"] = departmentName
		data["department_status"] = status
		data["company_id"] = companyId
		data["created_by"] = staffId
	}

	if ok := models.UpdateDepartment(Id, data); !ok {
		c.JSON(http.StatusOK, util.Fail("更新失败"))
		return
	} else {
		c.JSON(http.StatusOK, util.Success("更新成功", make(map[string]string)))
		return
	}
}

//删除部门
func DeleteDepartment(c *gin.Context) {
	id := com.StrTo(c.PostForm("department_id")).MustInt()
	valid := validation.Validation{}

	valid.Required(id, "id").Message("department_id 不能为空")
	valid.Min(id, 1, "department_id").Message("department_id 必须大于0")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	}
	if ok := models.DeleteDepartment(id); !ok {
		c.JSON(http.StatusOK,util.Fail("删除失败"))
		return
	} else {
		c.JSON(http.StatusOK,util.Success("删除成功",make(map[string]string)))
		return
	}
}
