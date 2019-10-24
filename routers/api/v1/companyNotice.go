package v1

import (
	"daily.com/models"
	"daily.com/pkg/e"
	"daily.com/pkg/setting"
	"daily.com/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//公告列表
func CompanyNotices(c *gin.Context) {
	page := com.StrTo(c.PostForm("page")).MustInt()
	status := com.StrTo(c.PostForm("status")).MustInt()
	valid := validation.Validation{}

	valid.Min(page, 1, "page").Message("page必须为大于1的证整数")
	valid.Range(status, 1, 2, "notice_status").Message("notice_status只能为1或者2")

	where := make(map[string]interface{})
	data := make(map[string]interface{})
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	} else {
		where["notice_status"] = status
	}
	data["list"] = models.CompanyNotices(util.GetPage(c), setting.PageSize, where)
	data["count"] = models.CompanyNoticesTotal(where)
	data["page"] = c.PostForm("page")
	c.JSON(http.StatusOK, util.Success("请求成功", data))
}

//获取单条公告
func CompanyNoticeOne(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	valid := validation.Validation{}

	valid.Required(id, "id").Message("id不能为空")
	valid.Min(id, 1, "id").Message("id必须为大于0的整数")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	}

	if data, ok := models.CompanyNoticeOne(id); !ok {
		c.JSON(http.StatusOK, util.Fail(e.GetMsg(e.ERROR_NOT_EXIST)))
		return
	} else {
		c.JSON(http.StatusOK, util.Success("查询成功", data))
		return
	}
}

//添加公告
func AddCompanyNotice(c *gin.Context) {
	params := make(map[string]interface{})
	title := c.PostForm("title")
	content := c.PostForm("content")
	status := com.StrTo(c.PostForm("status")).MustInt()
	companyId := com.StrTo(c.PostForm("company_id")).MustInt()
	createdBy := com.StrTo(c.PostForm("staff_id")).MustInt()
	valid := validation.Validation{}

	valid.Required(title, "title").Message("公告标题不能为空")
	valid.Required(content, "content").Message("公告内容不能为空")
	valid.Range(status, 1, 2, "status").Message("公告状态必须为1或者2")
	valid.Min(companyId, 1, "company_id").Message("company_id必须为大于0的整数")
	valid.Min(createdBy, 1, "created_by").Message("创建人必须为大于0的整数")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	} else {
		params["notice_title"] = title
		params["notice_content"] = content
		params["notice_status"] = status
		params["company_id"] = companyId
		params["created_by"] = createdBy
	}
	if ok := models.AddCompanyNotice(params); !ok {
		c.JSON(http.StatusOK, util.Fail("添加失败"))
		return
	} else {
		c.JSON(http.StatusOK, util.Success("添加成功", make(map[string]string)))
		return
	}
}

//更新公告
func UpdateCompanyNotice(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	title := c.PostForm("title")
	content := c.PostForm("content")
	status := c.PostForm("status")
	companyId := c.PostForm("company_id")
	staffId := c.PostForm("staff_id")
	valid := validation.Validation{}

	valid.Required(id, "id").Message("id不能空")
	valid.Required(title, "title").Message("title不能空")
	valid.Required(content, "content").Message("content不能空")
	valid.Required(status, "status").Message("status不能空")
	valid.Required(companyId, "company_id").Message("company_id不能空")
	valid.Required(staffId, "staff_id").Message("staff_id不能空")
}

//删除公告
func DeleteCompanyNotice(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("id 必须为大于0的整数")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.Fail(err.Message))
			return
		}
	} else {
		c.JSON(http.StatusOK, util.Success("", make(map[string]string)))
		return
	}
}
