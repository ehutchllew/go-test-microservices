package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitzero"`
}

func (c *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = c.writeJSON(w, http.StatusOK, payload)
}

func (c *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	if err := c.readJSON(w, r, &requestPayload); err != nil {
		c.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		c.authenticate(w, &requestPayload.Auth)
	default:
		c.errorJSON(w, errors.New("Unknown Action"), http.StatusBadRequest)
	}
}

func (c *Config) authenticate(w http.ResponseWriter, ap *AuthPayload) {
	jsonData, _ := json.MarshalIndent(ap, "", "\t")

	req, err := http.NewRequest("POST", "http://authentication-service:8999/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		c.errorJSON(w, err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("res: %+v\n", res)
	if res.StatusCode == http.StatusUnauthorized {
		c.errorJSON(w, errors.New("Invalid Credentials"))
		return
	} else if res.StatusCode != http.StatusOK {
		c.errorJSON(w, errors.New("Error calling auth service"))
		return
	}

	var jsonFromService JSONResponse
	err = json.NewDecoder(res.Body).Decode(&jsonFromService)
	if err != nil {
		c.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		c.errorJSON(w, errors.New(jsonFromService.Message), http.StatusUnauthorized)
		return
	}

	var payload JSONResponse
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromService.Data

	c.writeJSON(w, http.StatusOK, payload)
}
