package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"photomanagerapp/models"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "root:@tcp(localhost:3306)/photo_manager_db"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func MigrateDb() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
