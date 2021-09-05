package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"index"`
	Auth Auth `gorm:"foreignKey:Email;references:Email;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Photo *string `gorm:"column:photo"` // Path
	FirstName string `gorm:"column:firstname;notNull"`
	LastName string `gorm:"column:lastname;notNull"`
	TitlePhrase string `gorm:"column:title_phrase"`
	Birthdate string `gorm:"column:birthdate"`
	Citizenship string `gorm:"column:citizenship"`
	Languages string `gorm:"column:languages"`
	Gender string `gorm:"column:gender"`
	ResumeFile *string `gorm:"column:resume_file"` // Path
}

func UpdateOrCreateUser(condition interface{}, claim string) (error) {
	db := GetDB()
	var user User
	if db.Model(&user).Where("email=?", claim).Updates(condition).RowsAffected == 0 {
		return db.Save(condition).Error
	} else {
		return errors.New("Failed to either update or save...")
	}
}
