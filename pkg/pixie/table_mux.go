package pixie

import (
	"context"

	pb "orbservability/observer/pkg/gen/pb/v1"

	"px.dev/pxapi"
	"px.dev/pxapi/types"
)

// Satisfies the TableMuxer interface.
type TableMux struct {
	GrpcStream pb.EventGatewayService_StreamEventsClient
}

func (s *TableMux) AcceptTable(ctx context.Context, metadata types.TableMetadata) (pxapi.TableRecordHandler, error) {
	return &TablePrinter{
		GrpcStream: s.GrpcStream,
	}, nil
}
