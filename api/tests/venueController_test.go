package main

import (
	"encoding/json"
	"bytes"
	"testing"
	"net/http"
)

func TestCreateVenue(t *testing.T){
	var dummydata = []byte (`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	var venuedata = []byte (`{"name":"kyakabale", "description":"low dust v", "location":"brazil", "capacity":200, "category":"outdoor"}`)
	newreq, _ := http.NewRequest("POST", "/api/venues", bytes.NewBuffer(venuedata))
	newreq.Header.Set("Authorization", res["token"])
	newresponse := executeRequest(newreq)
	checkResponseCode(t, http.StatusCreated, newresponse.Code)
}

func TestGetVenues(t *testing.T){
	var dummydata = []byte (`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	newreq, _ := http.NewRequest("GET", "/api/venues", nil)
	newreq.Header.Set("Authorization", res["token"])
	newresponse := executeRequest(newreq)
	checkResponseCode(t, http.StatusOK, newresponse.Code)
}

func TestEditVenue(t *testing.T){
	var dummydata = []byte (`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	var venuedata = []byte (`{"name":"weconnect biz", "description":"low dust v", "location":"brazil", "capacity":200, "category":"outdoor"}`)
	newreq, _ := http.NewRequest("PUT", "/api/venues/1", bytes.NewBuffer(venuedata))
	newreq.Header.Set("Authorization", res["token"])
	newresponse := executeRequest(newreq)
	checkResponseCode(t, http.StatusOK, newresponse.Code)
}