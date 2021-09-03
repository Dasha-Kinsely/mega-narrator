package setup

import "github.com/Dasha-Kinsely/mega-narrator/models"

func MysqlMigration(){
	db := models.GetDB()
	db.AutoMigrate(&models.Auth{})
	db.AutoMigrate(&models.User{})
}