package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	localhost := os.Getenv("DATABASE_HOST")
	database := os.Getenv("DATABASE_NAME")
	password := os.Getenv("DATABASE_PASSWORD")
	user := os.Getenv("DATABASE_USER")
	port := os.Getenv("DATABASE_PORT")
	dsn := "host=" + localhost + " user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error when connecting to database")
	}
}
