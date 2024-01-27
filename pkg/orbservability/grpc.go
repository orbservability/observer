package orbservability

import (
	"context"

	pb "github.com/orbservability/schemas/v1"
	"github.com/orbservability/telemetry/pkg/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"orbservability/observer/pkg/config"
)

func CreateGrpcStream(ctx context.Context, cfg *config.Config) (*grpc.ClientConn, pb.EventGatewayService_StreamEventsClient, error) {
	conn, err := grpc.Dial(
		cfg.OrbservabilityURL,
		grpc.WithChainUnaryInterceptor(logs.UnaryClientInterceptor),
		grpc.WithChainStreamInterceptor(logs.StreamClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewEventGatewayServiceClient(conn)
	stream, err := client.StreamEvents(ctx)
	if err != nil {
		return nil, nil, err
	}

	return conn, stream, nil
}
