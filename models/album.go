package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Email string `gorm:"index"`
	Auth Auth `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Title string `gorm:"column:title;notNull"`
	Photo1 *string `gorm:"column:photo_1;notNull;index:collection"`
	Photo2 *string `gorm:"column:photo_2;notNull;index:collection"`
	Photo3 *string `gorm:"column:photo_3;notNull;index:collection"`
	Photo4 *string `gorm:"column:photo_4;notNull;index:collection"`
}