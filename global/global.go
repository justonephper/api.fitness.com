package global

import (
	"fitness/config"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//全局常用变量，类似于php中的超全局变量
var (
	DB          *gorm.DB //数据库对象
	RedisClient *redis.Client
	Config      config.Config
	Viper       *viper.Viper
	Logger      *zap.Logger
)
