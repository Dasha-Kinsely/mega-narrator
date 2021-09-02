package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Title string
	Icon *string
	AcquiredYears string
	Fluency int8
	Description string
}