package models

import "gorm.io/gorm"

type Hobby struct {
	gorm.Model
	Title string
	AltDescription string
}