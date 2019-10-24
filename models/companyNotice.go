package models

type CompanyNotice struct {
	Model
	NoticeTitle   string    `json:"notice_title"`
	NoticeContent string    `json:"notice_content"`
	NoticeStatus  int       `json:"notice_status"`
	CompanyId     int       `json:"company_id" gorm:"index"`
	Company       Companies `json:"company_name"`
	CreatedBy     int       `json:"created_by"`
	UpdatedBy     int       `json:"updated_by"`
}

//公司公告列表
func CompanyNotices(pageNum int, pageSize int, where interface{}) (companyNotices []CompanyNotice) {
	db.Preload("Company").Where(where).Offset(pageNum).Limit(pageSize).Order("id desc").Find(&companyNotices)
	return
}

//公告总数
func CompanyNoticesTotal(where interface{}) (count int) {
	db.Model(&CompanyNotice{}).Where(where).Count(&count)
	return
}

//获取公司公告
func CompanyNoticeOne(id int) (CompanyNotice, bool) {
	var notice CompanyNotice
	db.First(&notice, id)
	if notice.ID > 0 {
		return notice, true
	}
	return notice, false
}

//检测公司公告是否存在
func CompanyNoticeExist(id int) bool {
	var notice CompanyNotice
	db.Where("id = ?", id).First(&notice)
	if notice.ID > 0 {
		return true
	}
	return false
}

//添加公司公告
func AddCompanyNotice(data map[string]interface{}) bool {
	db.Create(&CompanyNotice{
		NoticeTitle:   data["notice_title"].(string),
		NoticeContent: data["notice_content"].(string),
		NoticeStatus:  data["notice_status"].(int),
		CompanyId:     data["company_id"].(int),
		CreatedBy:     data["created_by"].(int),
	})
	return true
}

//更新公司公告
func UpdateCompanyNotice(id int, data map[string]interface{}) bool {
	db.Where("id = ?", id).Updates(data)
	return true
}

//删除公司公告
func DeleteCompanyNotice(id int) bool {
	db.Where("id = ?", id).Delete(&CompanyNotice{})
	return true
}
