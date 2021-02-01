package auth

import (
	"api.fitness.com/bean/models"
	"api.fitness.com/bean/request"
	"api.fitness.com/global"
)

/**
 * @Desc:登录接口
 * @Author:haoge
 * @Date:2021/1/29 13:37
 **/
func Login(params request.LoginParam)  {
	//查询数据库
	global.DB.First(&models.Users{},map[string]string{"email":"15350732156@163.com"})
}