package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Email string `gorm:"index"`
	Auth Auth `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Title string `gorm:"column:title;notNull"`
	Icon *string `gorm:"column:icon"`
	AcquiredYears string `gorm:"column:acquired_years"`
	Fluency int8 `gorm:"column:fluency;notNull"`
	Description string `gorm:"column:description;notNull"`
}