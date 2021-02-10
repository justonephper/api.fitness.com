package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Users struct {
	ModelId
	FirstName         string    `json:"first_name" gorm:"type:varchar(64);comment:'姓氏'"`
	LastName          string    `json:"last_name" gorm:"size:64;comment:'名字'"`
	Name              string    `json:"name" gorm:"size:64;comment:'姓名'"`
	Email             string    `json:"email" gorm:"size:64;comment:'邮箱'"`
	HeadPhoto         string    `json:"head_photo" gorm:"comment:'头像地址'"`
	Password          string    `json:"password" gorm:"comment:'密码'"`
	PhonePrefix       string    `json:"phone_prefix" gorm:"size:16;comment:'手机号前缀'"`
	PhoneNumber       string    `json:"phone_number" gorm:"size:32;comment:'手机号'"`
	BusinessName      string    `json:"business_name" gorm:"size:64;comment:'公司名称'"`
	CountryCode       string    `json:"country_code" gorm:"size:4;comment:'国家code'"`
	Country           int64     `json:"country" gorm:"comment:'国家id'"`
	Province          int64     `json:"province" gorm:"comment:'省份id'"`
	City              int64     `json:"city" gorm:"comment:'城市id'"`
	PostCode          string    `json:"post_code" gorm:"size:16;comment:'邮编'"`
	EmailVerifiedAt   MyTime `json:"email_verified_at" gorm:"comment:'邮箱验证时间'"`
	VerificationToken string    `json:"verification_token" gorm:"comment:'jwt'"`
	Verified          *bool     `json:"verified" gorm:"comment:'是否激活,1:激活，0：未激活'"`
	NewUser           *bool     `json:"new_user" gorm:"comment:'是否新用户，1：新用户，0：非新用户'"`
	ModelTime
}

func (user *Users) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", time.Now().Unix())
	return nil
}

func (user *Users) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updated_at", time.Now().Unix())
	return nil
}
