package main

import (
	"net/http"

	goUtils "ehutchllew/go-utils"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := goUtils.JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	utils := goUtils.Utils{}

	_ = utils.WriteJSON(w, http.StatusOK, payload)
}
