package cache

import (
	"fitness/global"
	"fmt"
	"github.com/gin-gonic/gin"
)

var redisClient = global.RedisClient
var ctx = global.RedisClient.Context()

//hash操作
func Hash(c *gin.Context) {
	key := "myhash"
	//设置key指定的哈希集中指定字段的值
	for i := 0; i < 1000; i++ {
		redisClient.HSet(ctx, key, fmt.Sprintf("key%d", i), "hello")
	}
}
