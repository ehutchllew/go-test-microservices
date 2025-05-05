package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const PORT = "8999"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting auth service")

	// TODO: connect to DB

	// Setup config
	c := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: c.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
