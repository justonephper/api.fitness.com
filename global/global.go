package global

import (
	"github.com/gin-gonic/gin"
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

//请求上下文变量
var (
	C *gin.Context
)

const (
	TokenExpireDuration = time.Hour * 2
	RoleAdmin           = "admin"
	RoleUser            = "user"
	RoleStaff           = "staff"
)
