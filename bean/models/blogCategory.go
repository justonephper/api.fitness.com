package models

import (
	"fitness/global"
	"fitness/pkg/util/request"
)

type BlogCategory struct {
	ModelId
	Name string `json:"name" gorm:"default:'';size:64;comment:'博客分类名称'"`
	Desc string `json:"desc" gorm:"default:'';size:128;comment:'博客分类描述'"`
	ModelTime
}

//实例化对象
func NewBlogCategory() *BlogCategory {
	return &BlogCategory{}
}

//博客分类添加
func (c *BlogCategory) Add() bool {
	if err := global.DB.Create(c).Error; err != nil {
		return false
	}
	return true
}

//查询
func (c *BlogCategory) Find(id interface{}) bool {
	if global.DB.Where("id=?", id).First(c).RecordNotFound() {
		return false
	}
	return true
}

//更新
func (c *BlogCategory) Update(updateData map[string]interface{}) bool {
	if err := global.DB.Model(c).Updates(updateData).Error; err != nil {
		return false
	}
	return true
}

//删除
func (c *BlogCategory) Delete() bool {
	if err := global.DB.Delete(c).Error; err != nil {
		return false
	}
	return true
}

func (c *BlogCategory) Index(info request.PageInfo) (list []BlogCategory, total int64, err error) {
	q := info.Q
	order_key := info.OrderKey
	desc := info.Desc
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.DB.Model(c)

	if q != "" {
		db = db.Where("name LIKE ?", "%"+q+"%").Or("`desc` LIKE ?", "%"+q+"%")
	}

	//总条数
	if err = db.Count(&total).Error; err != nil {
		return list, 0, err
	}

	db = db.Limit(limit).Offset(offset)
	if order_key != "" {
		var OrderStr string
		if desc {
			OrderStr = order_key + " desc"
		} else {
			OrderStr = order_key
		}
		err = db.Order(OrderStr).Find(&list).Error
	} else {
		err = db.Order("id desc").Find(&list).Error
	}
	if err != nil {
		return list, 0, err
	}
	return
}
