package main

import (
	"context"

	"github.com/orbservability/io/pkg/client"
	_ "github.com/orbservability/telemetry/pkg/logs"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"orbservability/observer/pkg/config"
	"orbservability/observer/pkg/eventgateway"
	"orbservability/observer/pkg/pixie"
)

func main() {
	ctx := context.Background()

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

	// Initialize gRPC client
	eventGateway := &eventgateway.ServiceClient{}
	grpcConn, err := client.DialGRPC(cfg.OrbservabilityURL, eventGateway, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating gRPC connection")
	}
	defer grpcConn.Close()
	grpcStream, err := eventGateway.StreamEvents(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating gRPC stream")
	}
	defer grpcStream.CloseAndRecv()

	// Execute PxL scripts and handle records
	tm := &pixie.TableMux{GrpcStream: grpcStream}
	if err := pixie.ExecuteAndStream(ctx, pixieClient, cfg, tm); err != nil {
		log.Fatal().Err(err).Msg("Error handling records")
	}
}
