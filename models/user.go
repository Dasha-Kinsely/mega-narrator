package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string
	Photo *string
	FirstName string
	LastName string
	TitlePhrase string
	Birthdate string
	Languages []string
	Gender string
	ResumeFile *string
}
