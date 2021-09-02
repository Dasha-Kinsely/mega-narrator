package models

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Title string
	StartDate time.Time
	EndDate time.Time
	Responsibility string
}