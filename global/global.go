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

	PageSize  int
	JwtSecret string

	TimeFormate = "2006-01-02 15:04:05"

	DB *gorm.DB

	//Loger Loger
)
