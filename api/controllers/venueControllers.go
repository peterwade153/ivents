package controllers

import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/peterwade153/ivents/api/models"
	"github.com/peterwade153/ivents/api/responses"
)

// CreateVenue parses request, validates data and saves the new venue
func CreateVenue(w http.ResponseWriter, r *http.Request){
	var resp = map[string]interface{}{"status": "success", "message": "Venue successfully created"}

	user := r.Context().Value("userID").(float64)
	venue := &models.Venue{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &venue)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	venue.Prepare()

	if err = venue.Validate(""); err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if vne, _ := venue.GetVenue(); vne != nil{
		resp["status"] = "failed"
		resp["message"] = "Venue already registered, please choose another name"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	venue.UserID = uint(user)

	venueCreated, err := venue.Save()
	if err != nil{
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["venue"] = venueCreated
	responses.JSON(w, http.StatusCreated, resp)
	return
}
