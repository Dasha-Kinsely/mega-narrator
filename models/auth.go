package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Email string ``
	Password string ``
}