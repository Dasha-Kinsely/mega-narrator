package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	LinkedIn string
	Facebook string
	Twitter string
	Github string
	Phone string
}