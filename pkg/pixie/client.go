package pixie

import (
	"context"
	"orbservability/observer/pkg/config"

	"px.dev/pxapi"
)

func CreateClient(ctx context.Context, cfg *config.Config) (*pxapi.Client, error) {
	return pxapi.NewClient(
		ctx,
		pxapi.WithDirectAddr(cfg.PixieURL),
		pxapi.WithDirectCredsInsecure(),
	)
}
