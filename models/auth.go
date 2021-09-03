package models

type Auth struct {
	Email string `gorm:"column:email;primaryKey;unique"`
	PasswordHash string `gorm:"column:password_hash;notNull"`
}

func FindOneUser(condition interface{}) (Auth, error) {
	db := GetDB()
	var auth Auth
	err := db.First(&auth, condition).Error
	return auth, err
}