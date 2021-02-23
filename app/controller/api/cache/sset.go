package cache

import (
	"fitness/global"
	"fitness/pkg/util/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

//有序集合
func SortSet(c *gin.Context) {
	key := "sset1"
	//添加元素
	global.RedisClient.ZAdd(global.RedisClient.Context(), key, &redis.Z{
		Score:  100,
		Member: "zhangsan",
	}, &redis.Z{
		Score:  99,
		Member: "lisi",
	}, &redis.Z{
		Score:  101,
		Member: "wangwu",
	}, &redis.Z{
		Score:  97,
		Member: "liliu",
	})

	_, err := global.RedisClient.ZAdd(global.RedisClient.Context(), key, &redis.Z{
		Score:  105,
		Member: "zhangsan",
	}).Result()
	if err != nil {
		fmt.Println("insert the same value failed,err:", err)
		return
	}

	vals := global.RedisClient.ZRevRangeByScore(global.RedisClient.Context(), key, &redis.ZRangeBy{Max: "+inf", Min: "-inf"}).Val()

	fmt.Println("ZRevRangeByScore():", vals)

	//返回key的有序集元素个数 ZCard()
	num := global.RedisClient.ZCard(global.RedisClient.Context(), key).Val()
	fmt.Println("Zcard():", num)

	//区间统计
	count := global.RedisClient.ZCount(global.RedisClient.Context(), key, "100", "110").Val()
	fmt.Println("ZCount():", count)

	//删除分数最高者
	pop, _ := global.RedisClient.ZRangeByLex(global.RedisClient.Context(), key, &redis.ZRangeBy{
		Min: "-",
		Max: "+",
	}).Result()
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>%#v", pop)

	response.Success(c, "sorted set")
	return
}
