package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Users struct {
	ModelId
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	HeadPhoto         string `json:"head_photo"`
	Password          string `json:"password"`
	PhonePrefix       string `json:"phone_prefix"`
	PhoneNumber       string `json:"phone_number"`
	BusinessName      string `json:"business_name"`
	CountryCode       string `json:"country_code"`
	Country           int    `json:"country"`
	Province          int    `json:"province"`
	City              int    `json:"city"`
	PostCode          string `json:"post_code"`
	EmailVerifiedAt   string `json:"email_verified_at"`
	VerificationToken string `json:"verification_token"`
	Verified          uint8  `json:"verified"`
	NewUser           uint8  `json:"new_user"`
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
