package main

import (
	"net/http"
)

func (c *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JSONResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = c.writeJSON(w, http.StatusOK, payload)
}
