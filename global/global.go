package global

import (
	"github.com/jinzhu/gorm"
	"time"
)

var (
	RunMode string

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	APP_NAME  string
	JwtSecret []byte

	TimeFormate = "2006-01-02 15:04:05"

	DB *gorm.DB

	//Loger Loger
)

const (
	TokenExpireDuration = time.Hour * 2
	RoleAdmin           = "admin"
	RoleUser            = "user"
	RoleStaff           = "staff"
)
