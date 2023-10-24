package main

import (
	"net/http"

	utils "ehutchllew/go-utils"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = utils.WriteJSON(w, http.StatusOK, payload)
}
