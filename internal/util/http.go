package util

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const (
	MAX_RETRIES         = 10
	INITIAL_RETRY_DELAY = 1 * time.Second // Starting delay for retries
)

// GetURLContent fetches content from a given URL and returns an io.ReadCloser.
// It handles HTTP GET requests with retries and exponential backoff.
func GetURLContent(url string) (io.ReadCloser, error) {
	client := &http.Client{
		Timeout: 30 * time.Second, // Set a timeout for the *entire* request attempt
	}

	var resp *http.Response
	var err error

	for attempt := 0; attempt < MAX_RETRIES; attempt++ {

		req, reqErr := http.NewRequest("GET", url, nil)
		if reqErr != nil {
			// This is typically a fatal error (e.g., malformed URL), no point in retrying.
			return nil, fmt.Errorf("failed to create request for %s: %w", url, reqErr)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

		resp, err = client.Do(req)

		// Check for success condition
		if err == nil && resp.StatusCode == http.StatusOK {
			return resp.Body, nil // Success! Return the body
		}

		// Handle errors or non-OK status codes
		fmt.Printf("Attempt %d/%d failed: %s\n",
			attempt+1, MAX_RETRIES, func() string {
				if resp != nil {
					return resp.Status
				}
				return "N/A"
			}())

		if attempt < MAX_RETRIES-1 {
			// incremental backoff with randomness
			randomJitter := time.Duration(rand.Intn(1000)) * time.Millisecond
			sleepDuration := INITIAL_RETRY_DELAY*time.Duration(1<<attempt) + randomJitter

			fmt.Printf("Retrying in %v...\n", sleepDuration)
			time.Sleep(sleepDuration)
		} else {
			// Last attempt failed
			if err != nil {
				return nil, fmt.Errorf("HTTP request for %s failed after %d attempts: %w", url, MAX_RETRIES, err)
			}
			if resp != nil {
				resp.Body.Close() // Close body for non-200 responses on last attempt
				return nil, fmt.Errorf("HTTP request for %s failed with status %s after %d attempts", url, resp.Status, MAX_RETRIES)
			}
			// Should not reach here, but as a fallback
			return nil, fmt.Errorf("HTTP request for %s failed after %d attempts without specific error or status", url, MAX_RETRIES)
		}
	}

	// This part should technically be unreachable if the loop logic is sound,
	// but included for safety.
	return nil, fmt.Errorf("unexpected error: GetURLContent loop exited without returning or erroring for %s", url)
}
