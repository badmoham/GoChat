package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() {
	dataSourceName := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBUser, DBPassword, DBName)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	DB = db
}
