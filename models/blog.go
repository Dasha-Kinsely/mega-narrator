package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title string ``
	Image *string ``
	Text string ``
	Link string
}