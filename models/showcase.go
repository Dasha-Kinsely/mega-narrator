package models

import "gorm.io/gorm"

type Showcase struct {
	gorm.Model
	Title string
	Video *string 
}