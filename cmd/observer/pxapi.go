package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"px.dev/pxapi"
	"px.dev/pxapi/errdefs"
	"px.dev/pxapi/types"
)

// Satisfies the TableRecordHandler interface.
type tablePrinter struct {
	headerValues []string // A slice of strings to hold column names
}

func (t *tablePrinter) HandleInit(ctx context.Context, metadata types.TableMetadata) error {
	// Store column names in order
	for _, col := range metadata.ColInfo {
		t.headerValues = append(t.headerValues, col.Name)
	}
	return nil
}

func (t *tablePrinter) HandleRecord(ctx context.Context, r *types.Record) error {
	if len(r.Data) != len(t.headerValues) {
		return fmt.Errorf("%w: mismatch in header and data sizes", errdefs.ErrInvalidArgument)
	}

	recordMap := make(map[string]interface{})

	for i, d := range r.Data {
		var value interface{}

		switch v := d.(type) {
		case *types.BooleanValue:
			value = v.Value()
		case *types.Int64Value:
			value = v.Value()
		case types.Float64Value:
			value = v.Value()
		case *types.Time64NSValue:
			value = v.Value()
		case *types.UInt128Value:
			value = v.Value()
		default:
			// Fallback to string representation if type is unknown
			value = d.String()
		}

		recordMap[t.headerValues[i]] = value
	}

	jsonRecord, err := json.Marshal(recordMap)
	if err != nil {
		log.Printf("Error marshaling record to JSON: %s", err)
		return err
	}

	fmt.Println(string(jsonRecord))
	return nil
}

func (t *tablePrinter) HandleDone(ctx context.Context) error {
	return nil
}

// Satisfies the TableMuxer interface.
type tableMux struct{}

func (s *tableMux) AcceptTable(ctx context.Context, metadata types.TableMetadata) (pxapi.TableRecordHandler, error) {
	return &tablePrinter{}, nil
}
