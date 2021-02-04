package models

import "api.fitness.com/global"

type BlogCategory struct {
	ModelId
	Name string `json:"name" gorm:"default:'';size:64;comment:'博客分类名称'"`
	Desc string `json:"desc" gorm:"default:'';size:128;comment:'博客分类描述'"`
	ModelTime
}

//实例化对象
func BlogCategoryNew() *BlogCategory {
	return &BlogCategory{}
}

//博客分类添加
func (c *BlogCategory) Add() bool {
	if err := global.DB.Create(c).Error; err != nil {
		return false
	}
	return true
}
