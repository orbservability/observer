package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"px.dev/pxapi"
	"px.dev/pxapi/errdefs"
	"px.dev/pxapi/types"
)

const MaxErrorCount = 3

func main() {
	ctx := context.Background()

	// Create a Pixie client with local standalonePEM listening address
	client, err := pxapi.NewClient(
		ctx,
		pxapi.WithDirectAddr("127.0.0.1:12345"),
		pxapi.WithDirectCredsInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to create Pixie client: %v", err)
	}

	// Create a connection to the host.
	hostID := "localhost"
	vz, err := client.NewVizierClient(ctx, hostID)
	if err != nil {
		log.Fatalf("Failed to create Vizier client: %v", err)
	}

	// Create TableMuxer to accept results table.
	tm := &tableMux{}

	// Read PxL script from file
	content, err := os.ReadFile("./config/config.pxl")
	if err != nil {
		log.Fatalf("Failed to read PxL script: %v", err)
	}
	pxl := string(content)

	executionErrorCount := 0

	for {
		// Execute the PxL script and check for resultSet
		resultSet, err := vz.ExecuteScript(ctx, pxl, tm)
		if err != nil {
			executionErrorCount += 1
			if executionErrorCount > MaxErrorCount {
				log.Fatalf("Failed to execute PxL script: %v", err)
			} else {
				time.Sleep(time.Second)
				continue
			}
		}

		for {
			// Receive the PxL script results.
			err := resultSet.Stream()
			if err != nil {
				if err == io.EOF || err.Error() == "stream has already been closed" {
					// End of stream or stream closed, break to reopen stream
					break
				}
				if errdefs.IsCompilationError(err) {
					log.Fatalf("Compilation error: %v", err)
				}

				break
			}
		}
		resultSet.Close()
		time.Sleep(time.Second)
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
		recordMap[t.columnNames[i]] = d.String()
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
