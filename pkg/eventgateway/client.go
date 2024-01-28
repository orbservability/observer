package eventgateway

import (
	"context"

	pb "github.com/orbservability/schemas/v1"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	client pb.EventGatewayServiceClient
}

func (s *ServiceClient) RegisterClient(conn *grpc.ClientConn) {
	s.client = pb.NewEventGatewayServiceClient(conn)
}

func (s *ServiceClient) StreamEvents(ctx context.Context) (pb.EventGatewayService_StreamEventsClient, error) {
	return s.client.StreamEvents(ctx)
}
