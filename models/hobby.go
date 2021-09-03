package models

import "gorm.io/gorm"

type Hobby struct {
	gorm.Model
	Email string `gorm:"index"`
	Auth Auth `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Title string `gorm:"column:title;notNull"`
	AltDescription string `gorm:"column:alt_description"`
}