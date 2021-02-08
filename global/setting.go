package global

import "time"

//本文件存放业务元素和框架加载配置
var (
	TimeFormate = "2006-01-02 15:04:05"
	ConfigFile  = "config.yaml"
)

const (
	TokenExpireDuration = time.Hour * 2
	RoleAdmin           = "admin"
	RoleUser            = "user"
	RoleStaff           = "staff"
)
