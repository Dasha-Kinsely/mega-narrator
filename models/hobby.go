package models

import "gorm.io/gorm"

type Hobby struct {
	gorm.Model
	Auth Auth
	Title string
	AltDescription string
}