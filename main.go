package main

import (
	"api.fitness.com/bootstrap"
	"api.fitness.com/global"
	"api.fitness.com/routers"
	"fmt"
)

func main() {
	//框架初始化
	bootstrap.Init()
	// 路由信息初始化
	r := routers.Init()
	// 启动服务器并监听 8080 端口
	if err := r.Run();err != nil {
		fmt.Println("start up service failed,err:%v\n",err)
	}

	defer global.DB.Close()
}

