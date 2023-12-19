package main

import (
	"context"
	"io"
	"log"
	"time"

	"px.dev/pxapi"
	"px.dev/pxapi/errdefs"

	"orbservability/observer/pkg/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error with Config:", err)
	}

	// Create a Pixie client with local standalonePEM listening address
	client, err := pxapi.NewClient(
		ctx,
		pxapi.WithDirectAddr(cfg.PixieURL),
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

	executionErrorCount := 0

	for {
		// Execute the PxL script and check for resultSet
		resultSet, err := vz.ExecuteScript(ctx, cfg.PxL, tm)
		if err != nil {
			executionErrorCount += 1
			if executionErrorCount > cfg.MaxErrorCount {
				log.Fatalf("Failed to execute PxL script: %v", err)
			} else {
				time.Sleep(time.Second * time.Duration(cfg.PixieStreamSleep))
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
		time.Sleep(time.Second * time.Duration(cfg.PixieStreamSleep))
	}
}
