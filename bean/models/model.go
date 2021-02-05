package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ModelId struct {
	Id uint `gorm:"primary_key" json:"id"`
}

//自定义的数据类型必须实现 Scanner 和 Valuer 接口，以便让 GORM 知道如何将该类型接收、保存到数据库
//https://gorm.io/zh_CN/docs/data_types.html

//MyTime 自定义时间
type MyTime time.Time

type ModelTime struct {
	CreatedAt MyTime  `json:"created_at" gorm:"type:datetime"`
	UpdatedAt MyTime  `json:"updated_at" gorm:"type:datetime"`
	DeletedAt *MyTime `sql:"index" json:"deleted_at" gorm:"type:datetime"`
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 MyTime
func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

// 实现 driver.Valuer 接口，Value 返回 MyTime value
func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}