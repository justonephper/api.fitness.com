package models

type ModelId struct {
	Id uint `gorm:"primary_key"`
}

type ModelTime struct {
	DeletedAt string
	CreatedAt string
	UpdatedAt string
}
