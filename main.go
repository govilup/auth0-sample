package main

import (
	"auth0-sample/platform/authenticator"
	"auth0-sample/platform/router"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the end vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Printf("Server listening on http://localhost:8081/")
	if err := http.ListenAndServe("0.0.0.0:8081", rtr); err != nil {
		log.Fatalf("There was an error witht he http server: %v", err)
	}
}
