package global

import (
	"fitness/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//全局常用变量，类似于php中的超全局变量
var (
	C  *gin.Context //请求上下文变量
	DB *gorm.DB     //数据库对象
	Config config.Config
	Viper     *viper.Viper
	Logger    *zap.Logger
)
