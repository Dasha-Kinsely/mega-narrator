package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Auth Auth
	LinkedIn string
	Facebook string
	Twitter string
	Github string
	Phone string
}