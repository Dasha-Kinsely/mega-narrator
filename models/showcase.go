package models

import "gorm.io/gorm"

type Showcase struct {
	gorm.Model
	Auth Auth
	Title string
	Video *string 
}