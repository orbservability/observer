package pixie

import (
	"context"
	"fmt"
	"reflect"

	pb "github.com/orbservability/schemas/v1"
	"px.dev/pxapi/errdefs"
	"px.dev/pxapi/types"
)

// Satisfies the TableRecordHandler interface.
type TablePrinter struct {
	HeaderValues []string // A slice of strings to hold column names
	GrpcStream   pb.EventGatewayService_StreamEventsClient
}

func (t *TablePrinter) HandleInit(ctx context.Context, metadata types.TableMetadata) error {
	// Store column names in order
	for _, col := range metadata.ColInfo {
		t.HeaderValues = append(t.HeaderValues, col.Name)
	}
	return nil
}

func (t *TablePrinter) HandleRecord(ctx context.Context, r *types.Record) error {
	if len(r.Data) != len(t.HeaderValues) {
		return fmt.Errorf("%w: mismatch in header and data sizes", errdefs.ErrInvalidArgument)
	}

	msg := &pb.PixieEvent{}
	msgVal := reflect.ValueOf(msg).Elem()

	for i, d := range r.Data {
		fieldName := t.HeaderValues[i]
		field := msgVal.FieldByName(fieldName)
		if !field.IsValid() || !field.CanSet() {
			continue // Skip invalid or unsettable fields
		}

		// Assuming d is an interface{} containing the value
		fieldValue := reflect.ValueOf(d)

		if field.Type() != fieldValue.Type() {
			// Handle type conversion or skip the field
			continue
		}

		field.Set(fieldValue)
	}

	if err := t.GrpcStream.Send(msg); err != nil {
		return err
	}

	return nil
}

func (t *TablePrinter) HandleDone(ctx context.Context) error {
	return nil
}
