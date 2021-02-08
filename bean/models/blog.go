package models

import (
	"fitness/global"
	"fitness/pkg/util/request"
)

type Blog struct {
	ModelId

	Name    string `json:"name" gorm:"default:'';size:64;comment:'博客名称'"`
	Title   string `json:"title" gorm:"default:'';size:128;comment:'标题'"`
	Content string `json:"content" gorm:"type:text;not null;comment:'内容'"`
	ModelTime
}

//实例化对象
func NewBlog() *Blog {
	return &Blog{}
}

//添加
func (c *Blog) Create() bool {
	if err := global.DB.Create(c).Error; err != nil {
		return false
	}
	return true
}

//根据id查询是否存在
func (c *Blog) Find(id interface{}) bool {
	if global.DB.Where("id=?", id).First(c).RecordNotFound() {
		return false
	}
	return true
}

//删除记录
func (c *Blog) Delete() bool {
	err := global.DB.Delete(c).Error
	if err != nil {
		return false
	}
	return true
}

//修改记录
func (c *Blog) Update(updateData map[string]interface{}) bool {
	if err := global.DB.Model(c).Updates(updateData).Error; err != nil {
		return false
	}
	return true
}

//列表
func (c *Blog) Index(info request.PageInfo) (list []Blog, total int64, err error) {
	q := info.Q
	order_key := info.OrderKey
	desc := info.Desc
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.DB.Model(c)

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
	return list, total, err
}
