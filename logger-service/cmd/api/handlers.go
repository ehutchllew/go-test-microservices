package main

import (
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (c *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read json into var
	var requestPayload JSONPayload
	_ = c.readJSON(w, r, &requestPayload)

	// insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	if err := c.Models.LogEntry.Insert(event); err != nil {
		c.errorJSON(w, err)
		return
	}

	resp := JSONResponse{
		Error:   false,
		Message: "logged",
	}

	c.writeJSON(w, http.StatusOK, resp)
}
