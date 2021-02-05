package bootstrap

import (
	"fitness/bean/models"
	"fitness/global"
)

func InitMigration() {
	global.DB.AutoMigrate(
		models.Users{},
		models.BlogCategory{},
		models.Blog{},
	)
	return
}
