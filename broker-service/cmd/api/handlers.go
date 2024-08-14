package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	goUtils "ehutchllew/go-utils"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var utils = goUtils.Utils{}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := goUtils.JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = utils.WriteJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	default:
		utils.ErrorJSON(w, errors.New("Unknown Action"))
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		utils.ErrorJSON(w, errors.New("Invalid Credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		utils.ErrorJSON(w, errors.New("Error Calling Auth Service"))
		return
	}

	// create variable we'll read response.Body into
	var jsonFromService goUtils.JsonResponse

	// decode the json from the auth service
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		utils.ErrorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload goUtils.JsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromService.Data

	utils.WriteJSON(w, http.StatusAccepted, payload)
}
