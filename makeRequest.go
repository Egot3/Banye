package banye

import (
	"context"
	"io"
	"log"
	"net/http"
)

func (c *Client) MakeRequest(ctx context.Context, method, path string, auth map[string]string) (*Response, error) {
	req, _ := http.NewRequestWithContext(ctx, method, path, nil)

	req.SetBasicAuth(auth["username"], auth["password"])

	req.Header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Expires", "0")

	req.Header.Set("User-Agent", "banye/0.0.1") //lawful neutral

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Body parsing error: %v", err)
	}

	return &Response{
		Body:       bodyBytes,
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
	}, nil
}
