package cache

import (
	"fitness/global"
	"fitness/pkg/code"
	"fitness/pkg/util/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

//设置字符串
func SetString(c *gin.Context) {
	cacheKey := c.Query("cacheKey")
	cacheVal := c.Query("cacheVal")
	if cacheKey == "" || cacheVal == "" {
		response.Failed(c, code.Failed, "cacheKey and cacheVal is required")
		return
	}
	timer := time.Duration(5) * time.Second
	err := global.RedisClient.Set(global.RedisClient.Context(), cacheKey, cacheVal, timer).Err()
	if err != nil {
		response.Failed(c, code.Failed, fmt.Sprintf("Cache save failed:err:", err))
		return
	}
	response.Success(c, "set success")
	return
}

//获取字符串
func GetString(c *gin.Context) {
	cacheKey := c.Query("cacheKey")
	if cacheKey == "" {
		response.Failed(c, code.Failed, "request params not enough!")
		return
	}

	res, err := global.RedisClient.Get(global.RedisClient.Context(), cacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			response.Failed(c, code.Failed, "the value not exists!")
			return
		} else {
			response.Failed(c, code.Failed, fmt.Sprintf("get failed,res:%s", err))
			return
		}
		return
	}
	response.Success(c, res)
	return
}
