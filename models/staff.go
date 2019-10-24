package models

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Staff struct {
	Model
	Token         string `json:"token"`
	LoginPassword string `json:"login_password"`
	StaffName     string `json:"staff_name"`
	RoleId        int    `json:"role_id"`
	StaffNum      string `json:"staff_num"`
	StaffEmail    string `json:"staff_email"`
	StaffTel      string `json:"staff_tel"`
	GroupId       int    `json:"group_id"`
	StaffStatus   uint8  `json:"staff_status"`
	Gender        uint8  `json:"gender"`
	HeadImg       string `json:"head_img"`
	CompanyId     int    `json:"company_id"`
}

//根据用户id获取用户信息
func GetStaffById(id int) (Staff, bool) {
	staff := new(Staff)
	db.Where("id = ?", id).First(staff)
	if staff.ID > 0 {
		return *staff, true
	}
	return *staff, false
}

func GetStaffByStaffNum(staffNum string) (Staff, bool) {
	staff := new(Staff)
	db.Where("staff_num  = ?", staffNum).First(staff)
	if staff.ID > 0 {
		return *staff, true
	}
	return *staff, false
}

//生成token
func CreateToken(staff_id int) string {
	orgStr := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(staff_id)
	data := []byte(orgStr)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}

//生成密码
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

//校验密码
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//更新用户token
func (s *Staff) UpdateToken() bool {
	token := CreateToken(s.ID)
	db.Model(s).Where("id = ?", s.ID).Update(Staff{Token: token})
	return true
}

//获取全部用户
func GetStaffs(pageNum int, pageSize int, where interface{}) (staffs []Staff) {
	db.Where(where).Offset(pageNum).Limit(pageSize).Order("id asc").Find(&staffs)
	return
}

//获取数量
func GetStaffTotal(where interface{}) (count int) {
	db.Model(&Staff{}).Where(where).Count(&count)
	return
}
