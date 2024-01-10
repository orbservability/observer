package main

import (
	"context"
	"io"
	"orbservability/observer/pkg/config"
	"time"

	"px.dev/pxapi"
	"px.dev/pxapi/errdefs"
)

func createPixieClient(ctx context.Context, cfg *config.Config) (*pxapi.Client, error) {
	return pxapi.NewClient(
		ctx,
		pxapi.WithDirectAddr(cfg.PixieURL),
		pxapi.WithDirectCredsInsecure(),
	)
}

func executeAndStream(ctx context.Context, client *pxapi.Client, cfg *config.Config, tm *tableMux) error {
	vz, err := client.NewVizierClient(ctx, cfg.VizierHost)
	if err != nil {
		return err
	}

	executionErrorCount := 0
	for {
		resultSet, err := vz.ExecuteScript(ctx, cfg.PxL, tm)
		if err != nil {
			executionErrorCount++
			if executionErrorCount > cfg.MaxErrorCount {
				return err
			}
			time.Sleep(time.Second * time.Duration(cfg.PixieStreamSleep))
			continue
		}
		defer resultSet.Close()

		if err := streamResults(resultSet); err != nil {
			return err
		}

		time.Sleep(time.Second * time.Duration(cfg.PixieStreamSleep))
	}
}

func streamResults(resultSet *pxapi.ScriptResults) error {
	for {
		err := resultSet.Stream()
		if err != nil {
			if err == io.EOF || err.Error() == "stream has already been closed" {
				return nil // End of stream or stream closed, return successfully
			}
			if errdefs.IsCompilationError(err) {
				return err // Unrecoverable error
			}

			return nil // Unknown error, return successfully for retries
		}
	}
}
