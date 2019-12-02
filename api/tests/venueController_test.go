package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateVenue(t *testing.T) {
	var dummydata = []byte(`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	var venuedata = []byte(`{"name":"kyakabale", "description":"low dust v", "location":"brazil", "capacity":200, "category":"outdoor"}`)
	newreq, _ := http.NewRequest("POST", "/api/venues", bytes.NewBuffer(venuedata))
	newreq.Header.Set("Authorization", res["token"])
	newresponse := executeRequest(newreq)
	checkResponseCode(t, http.StatusCreated, newresponse.Code)
}

func TestGetVenues(t *testing.T) {
	var dummydata = []byte(`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	newreq, _ := http.NewRequest("GET", "/api/venues", nil)
	newreq.Header.Set("Authorization", res["token"])
	newresponse := executeRequest(newreq)
	checkResponseCode(t, http.StatusOK, newresponse.Code)
}

func TestGetVenue(t *testing.T) {
	tearDown(a.DB) // clearing the db
	seedUsers(a.DB)

	var dummydata = []byte(`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	// create venue to edit
	var venuedata = []byte(`{"name":"kyakabale", "description":"low dust v", "location":"brazil", "capacity":200, "category":"outdoor"}`)
	newreq, _ := http.NewRequest("POST", "/api/venues", bytes.NewBuffer(venuedata))
	newreq.Header.Set("Authorization", res["token"])
	newresponse := executeRequest(newreq)
	checkResponseCode(t, http.StatusCreated, newresponse.Code)

	getreq, _ := http.NewRequest("GET", "/api/venues/1", nil)
	getreq.Header.Set("Authorization", res["token"])
	getresponse := executeRequest(getreq)
	checkResponseCode(t, http.StatusOK, getresponse.Code)
}

func TestGetNonExistentVenue(t *testing.T) {
	tearDown(a.DB) // clearing the db
	seedUsers(a.DB)

	var dummydata = []byte(`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	// we shall use a random venue id
	getreq, _ := http.NewRequest("GET", "/api/venues/7", nil)
	getreq.Header.Set("Authorization", res["token"])
	getresponse := executeRequest(getreq)
	checkResponseCode(t, http.StatusInternalServerError, getresponse.Code)
}

func TestEditVenue(t *testing.T) {
	tearDown(a.DB) // clearing the db
	seedUsers(a.DB)

	var dummydata = []byte(`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	// create venue to edit
	var venuedata = []byte(`{"name":"kyakabale", "description":"low dust v", "location":"brazil", "capacity":200, "category":"outdoor"}`)
	newreq, _ := http.NewRequest("POST", "/api/venues", bytes.NewBuffer(venuedata))
	newreq.Header.Set("Authorization", res["token"])
	newresponse := executeRequest(newreq)
	checkResponseCode(t, http.StatusCreated, newresponse.Code)

	var editdata = []byte(`{"name":"weconnect biz", "description":"low dust v", "location":"brazil"}`)
	editreq, _ := http.NewRequest("PUT", "/api/venues/1", bytes.NewBuffer(editdata))
	editreq.Header.Set("Authorization", res["token"])
	editresponse := executeRequest(editreq)
	checkResponseCode(t, http.StatusOK, editresponse.Code)
}

func TestDeleteVenue(t *testing.T) {
	tearDown(a.DB) // clearing the db
	seedUsers(a.DB)

	var dummydata = []byte(`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	// Create Venue to delete
	var venuedata = []byte(`{"name":"rocket", "description":"low dust v", "location":"brazil", "capacity":200, "category":"outdoor"}`)
	newreq, _ := http.NewRequest("POST", "/api/venues", bytes.NewBuffer(venuedata))
	newreq.Header.Set("Authorization", res["token"])
	newresponse := executeRequest(newreq)
	checkResponseCode(t, http.StatusCreated, newresponse.Code)

	delreq, _ := http.NewRequest("DELETE", "/api/venues/1", nil)
	delreq.Header.Set("Authorization", res["token"])
	delresponse := executeRequest(delreq)
	checkResponseCode(t, http.StatusOK, delresponse.Code)
}
