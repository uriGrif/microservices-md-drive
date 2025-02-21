package main

import (
	"log"
	"net"
	"os"
	"permission-service/internal"
	"permission-service/protogen/golang/permissions"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file was not found")
	}

	grpcAddr := "0.0.0.0:" + os.Getenv("GRPC_PORT")

	// create a TCP listener on the specified port
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a gRPC server instance
	server := grpc.NewServer()

	db := internal.GetDbHandle()
	service := internal.NewPermissionService(db)

	permissions.RegisterPermissionServiceServer(server, &service)

	// start listening to requests
	log.Printf("Grpc server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
