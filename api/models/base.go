package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	fmt.Println("Setting up Database")

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbName := os.Getenv("DB_NAME")
	DbPassword := os.Getenv("DB_PASSWORD")

	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	db, err = gorm.Open("postgres", DBURI)

	if err != nil {
		fmt.Printf("Cannot connect to database %s", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s", DbName)
	}

	db.Debug().AutoMigrate(&User{}, &Venue{}) //database migration
}

// GetDb will return the db instance
func GetDb() *gorm.DB {
	return db
}
