package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateUser(t *testing.T) {
	var dummydata = []byte(`{"firstname":"james", "lastname":"low", "password":"gdshdshsd", "email":"jlow@yuj.com"}`)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(dummydata))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)
}

func TestCreateUserWithExistingEmail(t *testing.T) {
	var dummydata = []byte(`{"firstname":"james", "lastname":"low", "password":"gdshdshsd", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(dummydata))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestLoginUser(t *testing.T) {
	var dummydata = []byte(`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetUsers(t *testing.T) {
	var dummydata = []byte(`{"password":"password", "email":"steven@gmail.com"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(dummydata))
	response := executeRequest(req)

	var res map[string]string
	json.Unmarshal(response.Body.Bytes(), &res)

	newreq, _ := http.NewRequest("GET", "/api/users", nil)
	newreq.Header.Set("Authorization", res["token"])
	newresponse := executeRequest(newreq)
	checkResponseCode(t, http.StatusOK, newresponse.Code)
}
