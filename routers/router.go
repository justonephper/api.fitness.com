package routers

import (
	"daily.com/pkg/setting"
	v1 "daily.com/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	//中间件
	//r.Use(api.ApiRequestCheck())

	apiv1 := r.Group("/api/v1")

	//登录,退出
	{
		apiv1.POST("/login", v1.Login)
		apiv1.POST("/logout", v1.Logout)
	}

	//公司管理
	{
		apiv1.POST("/companies", v1.Companies)
		apiv1.POST("/company", v1.Company)
		apiv1.POST("/addCompany", v1.AddCompany)
		apiv1.POST("/updateCompany", v1.UpdateCompany)
		apiv1.POST("/deleteCompany", v1.DeleteCompany)
	}

	//部门管理
	{
		apiv1.POST("/department", v1.GetDepartment)
		apiv1.POST("/departments", v1.GetDepartments)
		apiv1.POST("/addDepartments", v1.AddDepartment)
		apiv1.POST("/updateDepartments", v1.UpdateDepartment)
		apiv1.POST("/deleteDepartments", v1.DeleteDepartment)
		apiv1.POST("/departmentGroup", v1.DepartmentGroup)
	}

	//公司公告管理
	{
		apiv1.POST("/companyNotices", v1.CompanyNotices)
		apiv1.POST("/companyNotice", v1.CompanyNoticeOne)
		apiv1.POST("/addCompanyNotice", v1.AddCompanyNotice)
		apiv1.POST("/updateCompanyNotice", v1.UpdateCompanyNotice)
		apiv1.POST("/deleteCompanyNotice", v1.DeleteCompanyNotice)
	}

	//目标管理
	{
		apiv1.GET("/target", v1.Targets)
		apiv1.GET("/target/:id", v1.OneTarget)
		apiv1.POST("/target", v1.AddTarget)
		apiv1.PUT("/target/:id", v1.EditTarget)
		apiv1.DELETE("/target/:id", v1.DeleteTarget)
	}

	//人员管理
	{
		apiv1.POST("/getStaffs", v1.GetStaffs)
		apiv1.POST("/getStaff", v1.GetStaff)
		apiv1.POST("/addStaff", v1.AddStaff)
		apiv1.POST("/updateStaff", v1.UpdateStaff)
		apiv1.POST("/deleteStaff", v1.DeleteStaff)
	}

	return r
}
