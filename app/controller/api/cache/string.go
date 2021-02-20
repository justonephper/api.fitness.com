package cache

import (
	"fmt"
	"fitness/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

//string相关操作
func Strings(c *gin.Context) {

	var err error

	key := "name"

	//设置值
	global.RedisClient.Set(global.RedisClient.Context(), key, "haoge", time.Duration(10)*time.Second)

	//判读存在不
	res, _ := global.RedisClient.Exists(global.RedisClient.Context(), key).Result()
	fmt.Println("Exists():", res)

	//获取值,带错误返回值
	str, _ := global.RedisClient.Get(global.RedisClient.Context(), key).Result()
	fmt.Println("get-result():", str)

	//获取值
	val := global.RedisClient.Get(global.RedisClient.Context(), "haha").Val()
	fmt.Println("get-val():", val)

	//append追加字符串
	err = global.RedisClient.Append(global.RedisClient.Context(), key, "111").Err()
	if err != nil {
		fmt.Println("append failed")
	}
	//获取追加后的值
	val1 := global.RedisClient.Get(global.RedisClient.Context(), key).Val()
	fmt.Println("append():", val1)

	//bit设置
	global.RedisClient.SetBit(global.RedisClient.Context(), "peter", 100, 1)
	global.RedisClient.SetBit(global.RedisClient.Context(), "peter", 101, 1)
	global.RedisClient.SetBit(global.RedisClient.Context(), "peter", 102, 1)

	//bitCount 统计次数
	all := global.RedisClient.BitCount(global.RedisClient.Context(), "peter", nil).Val()
	some := global.RedisClient.BitCount(global.RedisClient.Context(), "peter", &redis.BitCount{
		Start: 0,
		End:   -1,
	}).Val()
	fmt.Println("all:", all, "some:", some)

	//mset(一次性设置多个值)
	err = global.RedisClient.MSet(global.RedisClient.Context(), map[string]interface{}{
		"name":    "haoge",
		"sex":     "nan",
		"age":     28,
		"address": "beijingshi chongwenmen",
	}).Err()
	if err != nil {
		fmt.Println("Mset() failed")
	}
	fmt.Println("Mset() success")

	//mget 一次性获取多个值
	vals := global.RedisClient.MGet(global.RedisClient.Context(), "name", "sex", "age").Val()
	fmt.Println("mget():", vals)
}
