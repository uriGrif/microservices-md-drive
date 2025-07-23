package main

import (
	"files-service/internal"
	permissions "files-service/protogen/golang"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file was not found")
	}

	db := internal.GetDbHandle()
	client, err := grpc.NewClient(os.Getenv("PERMISSIONS_SERVICE_GRPC_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create permissions gRPC client: %v", err)
	}
	service := internal.NewFilesService(db, permissions.NewPermissionServiceClient(client))
	controller := internal.NewFilesController(&service)

	router := gin.Default()

	router.GET("files/by-user", controller.ListFilesByUser)
	router.GET("files/:id", controller.GetFile)
	router.POST("files/", controller.CreateFile)
	router.PUT("files/:id", controller.UpdateFile)
	router.DELETE("files/:id", controller.DeleteFile)

	router.Run()
}
