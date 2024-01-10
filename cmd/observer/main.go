package main

import (
	"context"
	"log"

	"orbservability/observer/pkg/config"
)

func main() {
	ctx := context.Background()

	// Load Config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	// Create a Pixie client
	pixieClient, err := createPixieClient(ctx, cfg)
	if err != nil {
		log.Fatal("Error creating Pixie client: ", err)
	}

	// Establish a connection to the gRPC server
	grpcConn, grpcStream, err := createGrpcStream(ctx, cfg)
	if err != nil {
		log.Fatal("Error creating gRPC stream: ", err)
	}
	defer grpcConn.Close()
	defer grpcStream.CloseAndRecv()

	// Execute PxL scripts and handle records
	tm := &tableMux{grpcStream: grpcStream}
	if err := executeAndStream(ctx, pixieClient, cfg, tm); err != nil {
		log.Fatal("Error handling records: ", err)
	}
}
