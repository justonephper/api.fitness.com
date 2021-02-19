package bootstrap

import (
	"fitness/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitRedis() *redis.Client {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		global.Logger.Error("redis connect ping failed, err:", zap.Any("err", err))
		return nil
	} else {
		global.Logger.Info("redis connect ping response:", zap.String("pong", pong))
		return client
	}
}
