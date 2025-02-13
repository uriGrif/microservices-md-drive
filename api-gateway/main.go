package main

import (
	"api-gateway/server"
	"log"

	"github.com/joho/godotenv"
	// "net/http/httputil"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init server
	server.Run()
}
