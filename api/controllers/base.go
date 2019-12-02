package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres

	"github.com/peterwade153/ivents/api/middlewares"
	"github.com/peterwade153/ivents/api/models"
	"github.com/peterwade153/ivents/api/responses"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize connect to the database and wire up routes
func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s", DbName)
	}

	a.DB.Debug().AutoMigrate(&models.User{}, &models.Venue{}) //database migration

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json

	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/register", a.UserSignUp).Methods("POST")
	a.Router.HandleFunc("/login", a.Login).Methods("POST")

	s := a.Router.PathPrefix("/api").Subrouter() // subrouter to add auth middleware
	s.Use(middlewares.AuthJwtVerify)

	s.HandleFunc("/users", a.GetAllUsers).Methods("GET")
	s.HandleFunc("/venues", a.CreateVenue).Methods("POST")
	s.HandleFunc("/venues", a.GetVenues).Methods("GET")
	s.HandleFunc("/venues/{id:[0-9]+}", a.GetVenue).Methods("GET")
	s.HandleFunc("/venues/{id:[0-9]+}", a.UpdateVenue).Methods("PUT")
	s.HandleFunc("/venues/{id:[0-9]+}", a.DeleteVenue).Methods("DELETE")
}

func (a *App) RunServer() {
	log.Printf("\nServer starting on port 5000")
	log.Fatal(http.ListenAndServe(":5000", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Ivents")
}
