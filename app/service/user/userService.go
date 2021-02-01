package user

import (
	"api.fitness.com/bean/models"
	"api.fitness.com/bean/request"
	"api.fitness.com/global"
)

/**
 * @Desc:用户列表
 * @Author:haoge
 * @Date:2021/2/1 11:40
 **/
func Index(params request.LogicParam) {

}

/**
 * @Desc:用户详情
 * @Author:haoge
 * @Date:2021/2/1 11:45
 **/
func Show(params request.LogicParam) {
	err := global.DB.Where("email=?","15350732156@163.com").First(&models.Users{}).Error
}

/**
 * @Desc:添加用户
 * @Author:haoge
 * @Date:2021/2/1 11:45
 **/
func Add(params request.LogicParam) {

}

/**
 * @Desc:用户编辑
 * @Author:haoge
 * @Date:2021/2/1 11:46
 **/
func Update(params request.LogicParam) {

}

/**
 * @Desc:用户删除
 * @Author:haoge
 * @Date:2021/2/1 11:46
 **/
func Destroy(params request.LogicParam) {

}
