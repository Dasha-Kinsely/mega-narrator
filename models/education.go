package models

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	Email string `gorm:"index"`
	Auth Auth `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Degree string `gorm:"column:title;notNull"`
	InstitutePhoto *string `gorm:"column:institute_photo"`
	Institute string `gorm:"column:institute;notNull"`
	StartTime time.Time `gorm:"column:start_time"`
	EndTime time.Time `gorm:"column:end_time"`
}