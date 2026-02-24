package banye

import (
	"net/http"
	"time"
)

// Create new client instance with middleware
func NewClient(base *http.Client) *Client {
	if base == nil {
		base = &http.Client{Timeout: 10 * time.Second}
	}
	return &Client{Base: base}
}
