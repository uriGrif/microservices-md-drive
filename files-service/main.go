package main

import (
	"files-service/internal"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file was not found")
	}

	db := internal.GetDbHandle()
	service := internal.NewFilesService(db)
	controller := internal.NewFilesController(&service)

	router := gin.Default()

	router.GET("/by-user/:userId", controller.ListFilesByUser)
	router.GET("/:id", controller.GetFile)
	router.POST("/", controller.CreateFile)
	router.PUT("/:id", controller.UpdateFile)
	router.DELETE("/:id", controller.DeleteFile)

	router.Run()
}
