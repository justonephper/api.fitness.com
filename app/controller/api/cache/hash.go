package cache

import (
	"fitness/global"
	"fitness/pkg/util/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

//hash操作
func Hash(c *gin.Context) {
	key := "myhash"

	//检测key中某个字段是否存在  return：bool
	res := global.RedisClient.HExists(global.RedisClient.Context(), key, "name").Val()
	fmt.Println("name-exists-check:", res)

	//检测key的属性数量
	hlen := global.RedisClient.HLen(global.RedisClient.Context(), key).Val()
	fmt.Println("hlen():", hlen)

	//获取指定键的所有属性值 hkeys
	allKey := global.RedisClient.HKeys(global.RedisClient.Context(), key).Val()
	fmt.Println("hkeys():", allKey)

	////如果不存在，则设置
	if !res {
		global.RedisClient.HSet(global.RedisClient.Context(), key, "name", "haoge")
		//global.RedisClient.HMSet(global.RedisClient.Context(), key, map[string]interface{}{
		//	"age":     28,
		//	"address": "beijingshi",
		//})
	}

	keys := global.RedisClient.HGetAll(global.RedisClient.Context(), key).Val()
	fmt.Println("HGetAll():", keys)

	//获取hash值
	val := global.RedisClient.HGet(global.RedisClient.Context(), key, "name").Val()
	fmt.Println("name:", val)

	vals := global.RedisClient.HMGet(global.RedisClient.Context(), key, "name", "age", "address").Val()
	fmt.Println("HMGet():", vals)

	//删除key所对应的值
	//global.RedisClient.HDel(global.RedisClient.Context(), key, "name")

	//获取键所对应的所有属性的值 hvals
	all_vals := global.RedisClient.HVals(global.RedisClient.Context(), key).Val()
	fmt.Println("HVals():", all_vals)

	//设置key指定的哈希集中指定字段的值
	//global.RedisClient.HMSet(global.RedisClient.Context(), key, map[string]interface{}{
	//	"child1": map[string]interface{}{
	//		"name": "child1-name",
	//		"age":  "child1-age",
	//	},
	//	"child2": map[string]interface{}{
	//		"name": "child2-name",
	//		"age":  "child2-age",
	//	},
	//	"child3": map[string]interface{}{
	//		"name": "child3-name",
	//		"age":  "child3-age",
	//	},
	//})

	response.Success(c, "ok")
	return
}
