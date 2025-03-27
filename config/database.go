package config

import (
	"fmt"
	"os"
	"wonderable/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = database.AutoMigrate(
		&models.User{},
		&models.Parent{}, 
		&models.Teacher{},
		&models.Student{},
		&models.Subject{},
	)

	if err != nil {
		fmt.Println(err)
	}

	DB = database
	fmt.Println("Database connected successfully!")
}
