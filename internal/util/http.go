package util

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// GetURLContent fetches content from a given URL and returns an io.ReadCloser.
// It handles basic HTTP GET requests.
func GetURLContent(url string) (io.ReadCloser, error) {
	client := &http.Client{
		Timeout: 30 * time.Second, // Set a timeout
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Important: Add a User-Agent header to mimic a real browser,
	// as some sites might block requests without it.
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close() // Close the body if status is not OK
		return nil, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	return resp.Body, nil
}
