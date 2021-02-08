package bootstrap

import (
	"fitness/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"os"
)

//初始化数据库并产生数据库全局变量
func InitDB() *gorm.DB {
	switch global.Config.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

//初始化数据库对象
func GormMysql() *gorm.DB {
	m := global.Config.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	if db, err := gorm.Open("mysql", dsn); err != nil {
		global.Logger.Error("MySQL启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		db.SingularTable(true)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		return db
	}
}
