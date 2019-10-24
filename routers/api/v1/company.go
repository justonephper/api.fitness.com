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

//公司列表
func Companies(c *gin.Context) {
	where := make(map[string]interface{})
	data := make(map[string]interface{})
	page := com.StrTo(c.PostForm("page")).MustInt()
	valid := validation.Validation{}

	if arg := c.DefaultPostForm("company_status", "all"); arg != "all" {
		company_status := com.StrTo(arg).MustInt()
		valid.Range(company_status, 1, 2, "company_id").Message("company_status必须是1,2或者all")
		where["company_status"] = company_status
	}
	valid.Min(page, 1, "page").Message("page 必须为大于0的整数")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Error()))
			return
		}
	}
	data["list"] = models.GetCompanies(util.GetPage(c), setting.PageSize, where)
	data["count"] = models.GetCompanayTotal(where)
	data["page"] = c.PostForm("page")
	c.JSON(http.StatusOK, util.Success("请求成功", data))
	return
}

//获取公司信息
func Company(c *gin.Context) {
	valid := validation.Validation{}
	var company_id int = -1
	if arg := c.PostForm("company_id"); arg != "" {
		company_id = com.StrTo(arg).MustInt()
	}
	valid.Min(company_id, 1, "company_id").Message("company_id必须存在且大于0")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	}
	data := models.GetCompanyById(company_id)
	c.JSON(http.StatusOK, util.Success("请求成功", data))
	return
}

//添加公司
func AddCompany(c *gin.Context) {
	companyName := c.PostForm("company_name")
	companyStatus := com.StrTo(c.PostForm("company_status")).MustInt()
	valid := validation.Validation{}

	valid.Required(companyName, "company_name").Message("公司名称不能为空")
	valid.Range(companyStatus, 1, 2, "company_status").Message("公司状态必须为1或者2")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Error()))
			return
		}
	}
	insertData := make(map[string]interface{})
	insertData["company_name"] = companyName
	insertData["company_status"] = companyStatus
	if ok := models.AddCompany(insertData); ok {
		c.JSON(http.StatusOK, util.Success("添加成功", make(map[string]interface{})))
		return
	} else {
		c.JSON(http.StatusOK, util.Fail("添加失败"))
	}
}

//更新公司
func UpdateCompany(c *gin.Context) {
	where := make(map[string]interface{})
	companyId := com.StrTo(c.PostForm("company_id")).MustInt()
	companyName := c.PostForm("company_name")
	companyStatus := com.StrTo(c.PostForm("company_status")).MustInt()
	valid := validation.Validation{}

	valid.Required(companyId, "company_id").Message("公司id不能为空")
	valid.Min(companyId, 1, "company_id").Message("公司id必须大于0")
	valid.Required(companyName, "company_name").Message("公司名称不能为空")
	valid.Required(companyStatus, "company_status").Message("公司状态不能为空")
	valid.Range(companyStatus, 1, 2, "company_status").Message("公司状态不合法")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	} else {
		where["company_name"] = companyName
		where["company_status"] = companyStatus
	}

	if ok := models.CompanyExists(companyId); !ok {
		c.JSON(http.StatusOK, util.Fail("公司信息不存在"))
		return
	}

	if ok := models.UpdateCompany(companyId, where); ok {
		c.JSON(http.StatusOK, util.Success("修改成功", make(map[string]string)))
	} else {
		c.JSON(http.StatusOK, util.Fail("更新失败"))
	}
}

//删除公司
func DeleteCompany(c *gin.Context) {
	companyId := com.StrTo(c.PostForm("company_id")).MustInt()
	companyStatus := com.StrTo(c.PostForm("company_status")).MustInt()
	valid := validation.Validation{}

	valid.Required(companyId, "company_id").Message("公司id不能为空")
	valid.Required(companyStatus, "company_status").Message("公司状态不能为空")
	valid.Min(companyId, 1, "company_id").Message("公司id必须大于0的整数")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	}
	if ok := models.DeleteCompany(companyId); ok {
		c.JSON(http.StatusOK, util.Success("删除成功", make(map[string]string)))
		return
	} else {
		c.JSON(http.StatusOK, util.Fail("删除失败"))
		return
	}
}
