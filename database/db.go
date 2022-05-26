package database

import (
	"fmt"
	"log"
	"notification/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		utils.ApiConfig.DBHost,
		utils.ApiConfig.DBUser,
		utils.ApiConfig.DBPassword,
		utils.ApiConfig.DBName,
		utils.ApiConfig.DBPort,
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	MigrateDatabase(DB)

	log.Println("connect to database")
	return DB, err
}

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&Notification{})
	if err != nil {
		log.Fatal(err.Error())
	}
}