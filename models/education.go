package models

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	Degree string
	InstitutePhoto *string
	Institute string
	StartTime time.Time
	EndTime time.Time
}