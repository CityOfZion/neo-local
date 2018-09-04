package services

import (
	"net/http"
	"time"

	"github.com/CityOfZion/neo-local/cli/logger"
)

const (
	endpoint = "http://localhost:4000"
)

var (
	sleepInterval = 3 * time.Second
)

// IsNeoScanStarted pings the NeoScan service to check if it has
// started.
func IsNeoScanStarted() error {
	spinner := logger.NewSpinner("Waiting for NeoScan to start")
	spinner.Start()

	for {
		resp, err := http.Get(endpoint)
		if err != nil || resp.StatusCode != http.StatusOK {
			time.Sleep(sleepInterval)
			continue
		}

		spinner.Stop()
		break
	}

	return nil
}
