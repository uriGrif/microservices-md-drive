package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"permission-service/protogen/golang/permissions"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// without this, only the headers starting with "Grpc-Metadata-" will be forwarded to the gRPC server
func CustomHeaderMatcher(key string) (string, bool) {
	switch strings.ToLower(key) {
	case "x-authenticated-user":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file was not found")
	}

	grpcAddr := os.Getenv("GRPC_URL") + ":" + os.Getenv("GRPC_PORT")

	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to permissions service: %v", err)
	}
	defer conn.Close()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(CustomHeaderMatcher),
	)
	if err = permissions.RegisterPermissionServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the permissions server: %v", err)
	}

	// start listening to requests from the gateway server
	addr := "0.0.0.0:" + os.Getenv("HTTP_PORT")
	fmt.Println("API gateway server is running on " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
