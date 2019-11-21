package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/peterwade153/ivents/api/models"
	"github.com/peterwade153/ivents/api/responses"
	"github.com/peterwade153/ivents/utils"
)

// UserSignUp controller for creating new users
func (a *App) UserSignUp(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Registered successfully"}

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

	usr, _ := user.GetUser(a.DB)
	if usr != nil {
		resp["status"] = "failed"
		resp["message"] = "User already registered, please login"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	user.Prepare() // here strip the text of white spaces

	err = user.Validate("") // default were all fields(email, lastname, firstname, password, profileimage) are validated
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userCreated, err := user.SaveUser(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["user"] = userCreated
	responses.JSON(w, http.StatusCreated, resp)
	return
}

// Login signs in users
func (a *App) Login(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "logged in"}

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

	user.Prepare() // here strip the text of white spaces

	err = user.Validate("login") // fields(email, password) are validated
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	usr, err := user.GetUser(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if usr == nil { // user is not registered
		resp["status"] = "failed"
		resp["message"] = "Login failed, please signup"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	err = models.CheckPasswordHash(user.Password, usr.Password)
	if err != nil {
		resp["status"] = "failed"
		resp["message"] = "Login failed, please try again"
		responses.JSON(w, http.StatusForbidden, resp)
		return
	}
	token, err := utils.EncodeAuthToken(usr.ID)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["token"] = token
	responses.JSON(w, http.StatusOK, resp)
	return
}

// GetAllUsers returns all users
func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
	return
}
