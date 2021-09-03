package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Email string `gorm:"index"`
	Auth Auth `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Title string `gorm:"column:title;notNull"`
	Image *string `gorm:"column:image;notNull"`
	Text string `gorm:"column:text;notNull"`
	Link string `gorm:"column:link"`
}