package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title string
	Photo []*string
}