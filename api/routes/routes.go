package routes

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/peterwade153/ivents/api/controllers"
	"github.com/peterwade153/ivents/api/middlewares"
	"github.com/peterwade153/ivents/api/responses"
)

func home(w http.ResponseWriter, r *http.Request){
	responses.JSON(w, http.StatusOK, "Welcome To Ivents")
} 

// Handlers routes
func Handlers() *mux.Router{

	r := mux.NewRouter().StrictSlash(true)
	r.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/users/register", controllers.UserSignUp).Methods("POST")
	r.HandleFunc("/users/login", controllers.Login).Methods("POST")
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	
	return r
}

