package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {

	err := godotenv.Load("..\\..\\.env")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	HOST := os.Getenv("HOST")
	USER := os.Getenv("USER")
	PASSWORD := os.Getenv("PASSWORD")
	DATABASE := os.Getenv("DATABASE")
	PORT := os.Getenv("PORT")
	SSLMODE := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", HOST, USER, PASSWORD, DATABASE, PORT, SSLMODE)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
