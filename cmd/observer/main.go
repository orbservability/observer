package main

import (
	"context"

	"github.com/orbservability/telemetry/pkg/logs"
	"github.com/rs/zerolog/log"

	"orbservability/observer/pkg/config"
	"orbservability/observer/pkg/orbservability"
	"orbservability/observer/pkg/pixie"
)

func main() {
	ctx := context.Background()

	// Initialize logger
	err := logs.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("Error initializing logger")
	}

	// Load Config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading config")
	}

	// Create a Pixie client
	pixieClient, err := pixie.CreateClient(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating Pixie client")
	}

	// Establish a connection to the gRPC server
	grpcConn, grpcStream, err := orbservability.CreateGrpcStream(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating gRPC stream")
	}
	defer grpcConn.Close()
	defer grpcStream.CloseAndRecv()

	// Execute PxL scripts and handle records
	tm := &pixie.TableMux{GrpcStream: grpcStream}
	if err := pixie.ExecuteAndStream(ctx, pixieClient, cfg, tm); err != nil {
		log.Fatal().Err(err).Msg("Error handling records")
	}
}
