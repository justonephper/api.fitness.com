package cache

import (
	"fmt"
	"fitness/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var len int64

//队列操作
func Lists(c *gin.Context) {
	var err error
	list1 := "list1"

	//rpush
	len = global.RedisClient.LLen(global.RedisClient.Context(), list1).Val()
	if len == 0 {
		num := 0
		for true {
			if num > 9 {
				break
			}
			global.RedisClient.RPush(global.RedisClient.Context(), list1, "haoge00"+strconv.Itoa(num))
			num++
		}
	}

	//获取队列长度
	len = global.RedisClient.LLen(global.RedisClient.Context(), list1).Val()
	fmt.Println("llen():", len)

	//获取队列元素
	vals := global.RedisClient.LRange(global.RedisClient.Context(), list1, 0, -1).Val()
	fmt.Println(vals)

	//出队
	//len = global.RedisClient.LLen(global.RedisClient.Context(), list1).Val()
	//if len > 0 {
	//	//循环出队
	//	for true {
	//		val, err := global.RedisClient.LPop(global.RedisClient.Context(), list1).Result()
	//		if err != nil {
	//			fmt.Println("list1 is empty!")
	//			break
	//		}
	//		fmt.Println("lpop()", val)
	//	}
	//}

	//按照下标获取值
	val := global.RedisClient.LIndex(global.RedisClient.Context(), list1, -1).Val()
	fmt.Println("LIndex():", val)

	//在指定位置插入数据
	//前插入
	err = global.RedisClient.LInsert(global.RedisClient.Context(), list1, "BEFORE", "haoge009", "haoge0090").Err()
	if err != nil {
		fmt.Println("LInsert() failed")
	}
	//后插入
	err = global.RedisClient.LInsert(global.RedisClient.Context(), list1, "AFTER", "haoge009", "haoge0092").Err()
	if err != nil {
		fmt.Println("LInsert() failed")
	}
	fmt.Println(global.RedisClient.LRange(global.RedisClient.Context(), list1, 0, -1))

	//lrem 删除某个值  lrem list num str
	//count > 0: 从头往尾移除值为 value 的元素。
	//count < 0: 从尾往头移除值为 value 的元素。
	//count = 0: 移除所有值为 value 的元素。
	err = global.RedisClient.LRem(global.RedisClient.Context(), list1, -2, "haoge0090").Err()
	if err != nil {
		fmt.Println("lrem() failed")
	}
	fmt.Println("lrem():", global.RedisClient.LRange(global.RedisClient.Context(), list1, 0, -1).Val())


	global.RedisClient.HGetAll(global.RedisClient.Context(),"hash1")

	c.String(http.StatusOK, "ok")
	return
}
