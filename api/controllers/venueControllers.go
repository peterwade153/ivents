package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/peterwade153/ivents/api/models"
	"github.com/peterwade153/ivents/api/responses"
)

// CreateVenue parses request, validates data and saves the new venue
func (a *App) CreateVenue(w http.ResponseWriter, r *http.Request) {
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

	if err = venue.Validate(); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if vne, _ := venue.GetVenue(a.DB); vne != nil {
		resp["status"] = "failed"
		resp["message"] = "Venue already registered, please choose another name"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	venue.UserID = uint(user)

	venueCreated, err := venue.Save(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["venue"] = venueCreated
	responses.JSON(w, http.StatusCreated, resp)
	return
}

func (a *App) GetVenues(w http.ResponseWriter, r *http.Request) {
	venues, err := models.GetVenues(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, venues)
	return
}

func (a *App) GetVenue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	venue, err := models.GetVenueById(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, venue)
	return
}

func (a *App) UpdateVenue(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Venue updated successfully"}

	vars := mux.Vars(r)

	user := r.Context().Value("userID").(float64)
	userID := uint(user)

	id, _ := strconv.Atoi(vars["id"])

	venue, err := models.GetVenueById(id, a.DB)

	if venue.UserID != userID {
		resp["status"] = "failed"
		resp["message"] = "Unauthorized venue update"
		responses.JSON(w, http.StatusUnauthorized, resp)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	venueUpdate := models.Venue{}
	if err = json.Unmarshal(body, &venueUpdate); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	venueUpdate.Prepare()

	_, err = venueUpdate.UpdateVenue(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, resp)
	return
}

func (a *App) DeleteVenue(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Venue deleted successfully"}

	vars := mux.Vars(r)

	user := r.Context().Value("userID").(float64)
	userID := uint(user)

	id, _ := strconv.Atoi(vars["id"])

	venue, err := models.GetVenueById(id, a.DB)

	if venue.UserID != userID {
		resp["status"] = "failed"
		resp["message"] = "Unauthorized venue delete"
		responses.JSON(w, http.StatusUnauthorized, resp)
		return
	}

	err = models.DeleteVenue(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, resp)
	return
}
