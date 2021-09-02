package models

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Title string
	Organization string
	StartTime time.Time
	EndTime time.Time
	Hilight []string
	Responsibility string
	ReferenceContact string
}