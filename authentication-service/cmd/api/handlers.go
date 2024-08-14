package main

import (
	"errors"
	"fmt"
	"net/http"

	goUtils "ehutchllew/go-utils"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	utils := goUtils.Utils{}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		utils.ErrorJSON(w, errors.New("Invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := goUtils.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	utils.WriteJSON(w, http.StatusAccepted, payload)
}
