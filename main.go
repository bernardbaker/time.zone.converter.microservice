package main

import (
	"log"
	"net"
	"os"

	"github.com/bernardbaker/time.zone.converter.microservice/app"
	"github.com/bernardbaker/time.zone.converter.microservice/infrastructure"
	"github.com/bernardbaker/time.zone.converter.microservice/proto"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	service := app.NewConverterService()

	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Server listening at %v", listener.Addr().String())

	server := grpc.NewServer()

	proto.RegisterTimeZoneConverterServer(server, infrastructure.NewTimeZoneConverterServer(service))

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("Starting gRPC server on port %s", port)
}
