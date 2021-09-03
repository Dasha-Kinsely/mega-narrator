package setup

import "github.com/Dasha-Kinsely/mega-narrator/models"

func MysqlMigration() {
	db := models.GetDB()
	db.AutoMigrate(&models.Auth{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Skill{})
	db.AutoMigrate(&models.Education{})
	db.AutoMigrate(&models.Experience{})
	db.AutoMigrate(&models.Blog{})
	db.AutoMigrate(&models.Album{})
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Hobby{})
	db.AutoMigrate(&models.Showcase{})
	return
}