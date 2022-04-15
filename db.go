package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() *gorm.DB {
	
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=ctf port=5432 TimeZone=Asia/Calcutta", user, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})

	return db
}
