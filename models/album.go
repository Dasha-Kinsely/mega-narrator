package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	CollectionName string
	Photo []*string
}