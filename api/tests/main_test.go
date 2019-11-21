package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
	"github.com/joho/godotenv"

	"github.com/peterwade153/ivents/api/controllers"
	"github.com/peterwade153/ivents/api/models"
)

var a controllers.App

func TestMain(m *testing.M) {

	if err := godotenv.Load(os.ExpandEnv("../../.env")); err != nil {
		os.Setenv("DB_HOST", "localhost")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_TEST_NAME", "postgres")
		os.Setenv("DB_PASSWORD", "postgres")
	}

	a = controllers.App{}
	a.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_TEST_NAME"),
		os.Getenv("DB_PASSWORD"))

	tearDown(a.DB)
	seedUsers(a.DB)
	seedVenues(a.DB)

	os.Exit(m.Run())
}

func tearDown(db *gorm.DB) error {
	fmt.Println("\n This is test db teardown")
	err := db.DropTableIfExists(&models.User{}, &models.Venue{}).Error
	if err != nil {
		fmt.Printf("This is the drop db error: %s", err)
		return err
	}
	err = db.AutoMigrate(&models.User{}, &models.Venue{}).Error
	if err != nil {
		fmt.Printf("This is the automigrate error: %s", err)
		return err
	}
	log.Println("Successfully refreshed table")
	return nil
}

func seedUsers(db *gorm.DB) ([]models.User, error) {

	var err error
	users := []models.User{
		models.User{
			FirstName: "Steven",
			LastName:  "victor",
			Email:     "steven@gmail.com",
			Password:  "password",
		},
		models.User{
			FirstName: "Kenny",
			LastName:  "Morris",
			Email:     "kenny@gmail.com",
			Password:  "password",
		},
	}
	for i := range users {
		err = db.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}
	return users, nil
}

func seedVenues(db *gorm.DB) ([]models.Venue, error) {
	var err error
	venues := []models.Venue{
		models.Venue{
			Name:        "weconnect",
			Description: "building houses",
			Location:    "kasese",
			Capacity:    29,
			Category:    "indoor",
		},
		models.Venue{
			Name:        "mandela",
			Description: "location found",
			Location:    "kampala",
			Capacity:    400,
			Category:    "outdoor",
		},
	}
	for i := range venues {
		err = db.Model(&models.Venue{}).Create(&venues[i]).Error
		if err != nil {
			return []models.Venue{}, err
		}
	}
	return venues, nil
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
