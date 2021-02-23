package cache

import (
	"fitness/global"
	"fitness/pkg/util/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

//无序集合
func Set(c *gin.Context) {

	key1 := "firstSet"
	//key2 := "secondSet"

	//添加单个元素
	global.RedisClient.SAdd(global.RedisClient.Context(), key1, "name")
	//添加多个元素
	global.RedisClient.SAdd(global.RedisClient.Context(), key1, "age", "address")

	//查询集合元素
	vals := global.RedisClient.SMembers(global.RedisClient.Context(), key1).Val()
	fmt.Println(vals)

	//查询key下的基数 (集合元素的数量).
	num := global.RedisClient.SCard(global.RedisClient.Context(), key1)
	fmt.Println("Scard():", num)

	//测试 diff()
	key2 := "key2"
	key3 := "key3"
	global.RedisClient.SAdd(global.RedisClient.Context(), key2, "a", "b", "c")
	global.RedisClient.SAdd(global.RedisClient.Context(), key3, "b", "c", "d")
	sdiff1 := global.RedisClient.SDiff(global.RedisClient.Context(), key2, key3).Val()
	fmt.Println("key2-key3-diff:", sdiff1)

	//SDiffStore() 差集
	key4 := "key4"
	global.RedisClient.SDiffStore(global.RedisClient.Context(), key4, key1, key2)
	fmt.Println("key1-key2-diff:", global.RedisClient.SMembers(global.RedisClient.Context(), key4))

	//SInter()   交集
	sinter := global.RedisClient.SInter(global.RedisClient.Context(), key2, key3).Val()
	fmt.Println(sinter)

	//SInterStore 交集并存储
	sinterStoreVal := "sinterStoreVal"
	global.RedisClient.SInterStore(global.RedisClient.Context(), sinterStoreVal, key2, key3)
	fmt.Println("SInterStore():", global.RedisClient.SMembers(global.RedisClient.Context(), sinterStoreVal))

	//检查集合中，某个键是否存在
	key_exists := global.RedisClient.SIsMember(global.RedisClient.Context(), key1, "name").Val()
	fmt.Println("name is exists in key1:", key_exists)

	//SMove 从某一个集合移动到另一个集合
	fmt.Println("before move:", global.RedisClient.SMembers(global.RedisClient.Context(), key1), global.RedisClient.SMembers(global.RedisClient.Context(), key2))
	global.RedisClient.SMove(global.RedisClient.Context(), key1, key2, "name")
	fmt.Println("after move:", global.RedisClient.SMembers(global.RedisClient.Context(), key1), global.RedisClient.SMembers(global.RedisClient.Context(), key2))

	rand1 := global.RedisClient.SRandMember(global.RedisClient.Context(), key2).Val()
	rand2 := global.RedisClient.SRandMemberN(global.RedisClient.Context(), key2, 2).Val()
	fmt.Println("rand1,rand2:", rand1, rand2)

	//从结合中移除某个元素
	fmt.Println("before Srem():", global.RedisClient.SMembers(global.RedisClient.Context(), key2).Val())
	global.RedisClient.SRem(global.RedisClient.Context(), key2, "name")
	fmt.Println("after Srem():", global.RedisClient.SMembers(global.RedisClient.Context(), key2).Val())

	//获取并集
	sunion := global.RedisClient.SUnion(global.RedisClient.Context(), key1, key2).Val()
	fmt.Println("union():", sunion)

	//求并集并存储
	global.RedisClient.SUnionStore(global.RedisClient.Context(), "sunionStore", key1, key2)
	fmt.Println(global.RedisClient.SMembers(global.RedisClient.Context(),"sunionStore"))

	//删除key
	//global.RedisClient.Del(global.RedisClient.Context(), key1)

	response.Success(c, "ok")
	return
}
