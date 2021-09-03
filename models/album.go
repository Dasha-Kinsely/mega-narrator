package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Auth Auth
	Title string
	Photo []*string
}