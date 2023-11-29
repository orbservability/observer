package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"px.dev/pxapi"
	"px.dev/pxapi/errdefs"
	"px.dev/pxapi/types"
)

func main() {
	// Create a Pixie client with local standalonePEM listening address
	ctx := context.Background()
	client, err := pxapi.NewClient(
		ctx,
		pxapi.WithDirectAddr("127.0.0.1:12345"),
		pxapi.WithDirectCredsInsecure(),
	)
	if err != nil {
		panic(err)
	}

	// Create a connection to the host.
	hostID := "localhost"
	vz, err := client.NewVizierClient(ctx, hostID)
	if err != nil {
		panic(err)
	}

	// Create TableMuxer to accept results table.
	tm := &tableMux{}

	// Read PxL script from file
	content, err := os.ReadFile("./config/config.pxl")
	if err != nil {
		panic(err)
	}
	pxl := string(content)

	// Execute the PxL script and get the resultSet
	resultSet, err := vz.ExecuteScript(ctx, pxl, tm)
	if err != nil {
		panic(err)
	}
	defer resultSet.Close()

	// Loop to receive the PxL script results.
	for {
		err := resultSet.Stream()
		if err != nil {
			if err == io.EOF {
				// End of stream
				break
			}
			if errdefs.IsCompilationError(err) {
				log.Printf("Error compiling stream: %s", err.Error())
				break
			}
			// Handle other kinds of runtime errors

			log.Printf("Error streaming: %+v", err)
		}
	}
}

// Satisfies the TableRecordHandler interface.
type tablePrinter struct {
	columnNames []string // A slice of strings to hold column names
}

func (t *tablePrinter) HandleInit(ctx context.Context, metadata types.TableMetadata) error {
	// Store column names in order
	for _, col := range metadata.ColInfo {
		t.columnNames = append(t.columnNames, col.Name)
	}
	return nil
}

func (t *tablePrinter) HandleRecord(ctx context.Context, r *types.Record) error {
	recordMap := make(map[string]interface{})

	for i, d := range r.Data {
		recordMap[t.columnNames[i]] = d
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
