package models

import (
	"daily.com/pkg/setting"
	"github.com/jinzhu/gorm"
	"time"
)

type Companies struct {
	Model
	CompanyName   string `json:"company_name"`
	CompanyStatus int    `json:"company_status"`
}

//获取公司列表
func GetCompanies(pageNum int, pageSize int, where interface{}) (company []Companies) {
	db.Where(where).Offset(pageNum).Limit(pageSize).Find(&company)
	return
}

//获取公司条数
func GetCompanayTotal(where interface{}) (count int) {
	db.Model(&Companies{}).Where(where).Count(&count)
	return
}

//获取公司信息
func GetCompanyById(company_id int) (company Companies) {
	db.Where("id = ?", company_id).First(&company)
	return
}

//添加公司
func AddCompany(data map[string]interface{}) bool {
	db.Create(&Companies{
		CompanyName:   data["company_name"].(string),
		CompanyStatus: data["company_status"].(int),
	})
	return true
}

//检测公司是否存在
func CompanyExists(company_id int) bool {
	var company Companies
	db.Where("id = ?", company_id).First(&company)
	if company.ID > 0 {
		return true
	}
	return false
}

//修改公司信息
func UpdateCompany(id int, data map[string]interface{}) bool {
	db.Model(&Companies{}).Where("id = ?", id).Updates(data)
	return true
}

//删除公司
func DeleteCompany(id int) bool {
	db.Where("id = ?", id).Delete(&Companies{})
	return true
}

func (company *Companies) BeforeCreate(scope *gorm.Scope) error {
	now := time.Now().Format(setting.TimeFormat)
	scope.SetColumn("CreatedAt", now)
	scope.SetColumn("UpdatedAt", now)

	return nil
}

func (company *Companies) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Format(setting.TimeFormat))

	return nil
}
