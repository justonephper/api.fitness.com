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

	//获取分数最高者  (5.0版本以上支持)
	//err = global.RedisClient.ZPopMax(global.RedisClient.Context(), key).Err()
	//if err != nil {
	//	response.Success(c, fmt.Sprintf("ZPopMax() failed,err:", err))
	//	return
	//}

	//获取一定范围你的元素
	child, err := global.RedisClient.ZRange(global.RedisClient.Context(), key, 0, -1).Result()
	if err != nil {
		response.Success(c, fmt.Sprintf("ZRang() failed,err:", err))
		return
	}
	fmt.Println("ZRang():", child)

	//获取成员的排名
	member := global.RedisClient.ZRank(global.RedisClient.Context(), key, "wangwu").Val()
	fmt.Println("ZRank():", member)

	//移除成员
	rem_res := global.RedisClient.ZRem(global.RedisClient.Context(), key, "zhangsan").Val()
	fmt.Println("ZRem():", rem_res)

	//查询成员
	fmt.Println("all member:", global.RedisClient.ZRange(global.RedisClient.Context(), key, 0, -1).Val())

	//查询成员分数
	score := global.RedisClient.ZScore(global.RedisClient.Context(), key, "liliu").Val()
	fmt.Println("ZScore():", score)

	//扫描键值对
	keys, _ := global.RedisClient.ZScan(global.RedisClient.Context(), key, 0, "", 0).Val()
	fmt.Println("ZScan():", keys)

	response.Success(c, "sorted set")
	return
}
