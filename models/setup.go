package models

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"table_top_api/utils"
)

var DB *gorm.DB

func InitDB() {
	dbStr, err := utils.OpenDBStr()
	if err != nil {
		log.Fatal("Failed to open .env file:", err)
	}
	database, err := gorm.Open(mysql.Open(dbStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(&Game{}, &User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	DB = database
}
