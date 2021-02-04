package bootstrap

import (
	"api.fitness.com/bean/models"
	"api.fitness.com/global"
)

func InitMigration() {
	global.DB.AutoMigrate(
		models.Users{},
	)
	return
}
