package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.readJSON(w, r, &requestPayload); err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against DB
	user, err := c.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		c.errorJSON(w, errors.New("Invalid Credentials"), http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		c.errorJSON(w, errors.New("Invalid Credentials"), http.StatusUnauthorized)
		return
	}

	payload := JSONResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	c.writeJSON(w, http.StatusOK, payload)
}
