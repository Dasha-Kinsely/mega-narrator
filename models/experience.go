package models

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Email string `gorm:"index"`
	Auth Auth `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Title string `gorm:"column:title;notNull"`
	Organization string `gorm:"column:organization;notNull"`
	StartTime time.Time `gorm:"column:start_time"`
	EndTime time.Time `gorm:"column:end_time"`
	Responsibility string `gorm:"column:responsibility;notNull"`
	ReferenceContact string `gorm:"column:reference_contact"`
}