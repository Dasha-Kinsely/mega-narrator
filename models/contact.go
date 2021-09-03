package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Email string `gorm:"index"`
	Auth Auth `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Title string `gorm:"column:title;notNull"`
	LinkedIn string `gorm:"column:linkedin"`
	Facebook string `gorm:"column:facebook"`
	Twitter string `gorm:"column:twitter"`
	Github string `gorm:"column:github"`
	Phone string `gorm:"column:phone"`
}