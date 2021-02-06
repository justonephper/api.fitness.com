package main

import (
	"fitness/bootstrap"
	"fitness/global"
	"fmt"
)

func main() {
	//框架初始化
	r := bootstrap.Init()
	// 启动服务器并监听 8080 端口
	if err := r.Run();err != nil {
		fmt.Println("start up service failed,err:%v\n",err)
	}

	defer global.DB.Close()
}

