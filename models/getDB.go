package models

import (
	"gorm.io/gorm"
)

// NOTE: This is your primary database and should remain consistent across any server instance
var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}