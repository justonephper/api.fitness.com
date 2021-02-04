package models

import (
	"api.fitness.com/app/helper/request"
	"api.fitness.com/global"
	"fmt"
)

type Blog struct {
	ModelId
	Name    string `json:"name" gorm:"default:'';size:64;comment:'博客名称'"`
	Title   string `json:"title" gorm:"default:'';size:128;comment:'标题'"`
	Content string `json:"content" gorm:"type:text;not null;comment:'内容'"`
	ModelTime
}

//实例化对象
func BlogNew() *Blog {
	return &Blog{}
}

//添加
func (blog *Blog) Create() bool {
	if err := global.DB.Create(blog).Error; err != nil {
		return false
	}
	return true
}

//根据id查询是否存在
func (blog *Blog) Find(id interface{}) bool {
	if global.DB.Where("id=?", id).First(blog).RecordNotFound() {
		return false
	}
	return true
}

//删除记录
func (blog *Blog) Delete() bool {
	err := global.DB.Delete(blog).Error
	if err != nil {
		return false
	}
	return true
}

//修改记录
func (blog *Blog) Update(updateData map[string]interface{}) bool {
	if err := global.DB.Model(blog).Updates(updateData).Error; err != nil {
		return false
	}
	return true
}

//列表
func (blog *Blog) Index(info request.PageInfo) (list []Blog, total int64, err error) {
	q := info.Q
	order := info.Order
	desc := info.Desc
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	fmt.Println(limit)
	fmt.Println(offset)

	db := global.DB.Model(blog)

	//搜索关键词
	if q != "" {
		db = db.Where("name LIKE ?", "%"+q+"%").
			Or("title LIKE ?", "%"+q+"%").
			Or("content LIKE ?", "%"+q+"%")
	}

	if err = db.Count(&total).Error; err != nil {
		return list, total, err
	}

	db = db.Limit(limit).Offset(offset)
	//排序
	if order != "" {
		var OrderStr string
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
		err = db.Order(OrderStr).Find(&list).Error
	} else {
		err = db.Order("id desc").Find(&list).Error
	}
	return list, total, err
}
