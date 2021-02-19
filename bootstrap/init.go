package bootstrap

import (
	"fitness/global"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//The whole framework starts here

func Init() *gin.Engine {
	//分步骤初始化，注意初始化顺序

	//1. 配置文件解析
	global.Viper = InitConfig()

	//2. 初始化zap日志库
	global.Logger = InitZap()

	//2.1 DB初始化
	global.DB = InitDB()

	//2.2 redis初始化
	global.RedisClient = InitRedis()

	//3. 迁移文件初始化
	InitMigration()

	//4. 路由初始化
	return InitRouter()
}
