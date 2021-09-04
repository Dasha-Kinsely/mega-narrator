package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Email string `gorm:"index"`
	Auth Auth `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Title string `gorm:"column:title;notNull"`
	Photo1 *string `gorm:"column:photo_1;notNull"`
	Photo2 *string `gorm:"column:photo_2;notNull"`
	Photo3 *string `gorm:"column:photo_3;notNull"`
	Photo4 *string `gorm:"column:photo_4;notNull"`
}