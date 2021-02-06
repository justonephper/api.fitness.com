package bootstrap

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//The whole framework starts here

func Init() *gin.Engine {
	//分步骤初始化，注意初始化顺序

	//1. 配置文件解析
	InitConfig()

	//2. DB初始化
	InitDB()

	//3. 迁移文件初始化
	InitMigration()

	//4. 路由初始化
	return InitRouter()
}
