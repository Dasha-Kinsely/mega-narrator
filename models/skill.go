package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Auth Auth
	Title string
	Icon *string
	AcquiredYears string
	Fluency int8
	Description string
}