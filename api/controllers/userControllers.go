package controllers

 import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/peterwade153/ivents/api/models"
	"github.com/peterwade153/ivents/api/responses"
)

// CreateUser controller for creating new users
func CreateUser(w http.ResponseWriter, r *http.Request){
	user := &models.User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = user.UserExists()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare() // here strip the text

	err = user.Validate("") // default were all fields(email, lastname, firstname, password, profileimage) are validated
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userCreated, err := user.SaveUser()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}
	responses.JSON(w, http.StatusCreated, userCreated)
}

// GetAllUsers returns all users
func GetAllUsers(w http.ResponseWriter, r *http.Request){
	user := &models.User{}
	
	users, err := user.GetAllUsers()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, users)
}

