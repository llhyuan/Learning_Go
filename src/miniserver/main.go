package main

import (
	"log"
	"miniserver/handlers"
	"net/http"
	"time"
)

func main() {
	handleFuncs := handlers.Handlers()

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handleFuncs,
	}

	log.Println("Starting the server.")

	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}

}
