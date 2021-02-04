package bootstrap

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



//The whole framework starts here

func Init() {
	//1. 配置文件解析
	InitConfig()

	//2. DB初始化
	InitDB()

	//3. 迁移文件初始化
	InitMigration()
}