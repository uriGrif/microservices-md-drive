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
		log.Println(".env file was not found")
	}

	// init server
	server.Run()
}
