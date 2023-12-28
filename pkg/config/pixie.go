package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PixieURL         string
	PxLFilePath      string
	PxL              string
	PixieStreamSleep int
	MaxErrorCount    int
}

// NewConfig creates a new Config struct with default configuration.
// It attempts to override these defaults with environment variables if they are set.
//
// Returns:
//   - A pointer to an Config struct which contains configuration settings.
//   - An error.
//
// Usage:
//
//	config, err := NewConfig()
//	if err != nil {
//		// handle error
//	}
func NewConfig() (*Config, error) {
	config := &Config{
		PixieURL:         "127.0.0.1:12345",     // Default URL
		PxLFilePath:      "./config/config.pxl", // Default script path
		PxL:              "",                    // PxL script
		PixieStreamSleep: 10,                    // Default sleep time in seconds
		MaxErrorCount:    3,                     // Default maximum error count
	}

	// Override defaults if environment variables are set
	if url := os.Getenv("PIXIE_URL"); url != "" {
		config.PixieURL = url
	}
	if path := os.Getenv("PXL_FILE_PATH"); path != "" {
		config.PxLFilePath = path
	}
	if sleep := os.Getenv("PIXIE_STREAM_SLEEP"); sleep != "" {
		val, err := strconv.Atoi(sleep)
		if err != nil {
			return nil, fmt.Errorf("error: %w", err)
		} else {
			config.PixieStreamSleep = val
		}
	}
	if maxErr := os.Getenv("PIXIE_ERROR_MAX"); maxErr != "" {
		val, err := strconv.Atoi(maxErr)
		if err != nil {
			return nil, fmt.Errorf("error: %w", err)
		} else {
			config.MaxErrorCount = val
		}
	}

	// Read PxL script from file
	content, err := os.ReadFile(config.PxLFilePath)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}
	config.PxL = string(content)

	return config, nil
}
