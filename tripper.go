package banye

import (
	"net/http"
)

// Pass a client(client.NewClient) instance and function that will execute before and after
// tip: you can scaffold those funcs, but it wasn't intended
func (c *Client) UseTripper(before func(*http.Request) error, after func(*http.Request, *http.Response, error) error) {
	if c.Base.Transport == nil {
		c.Base.Transport = http.DefaultTransport
	}

	c.Base.Transport = &MiddlewareTripper{
		Next:   c.Base.Transport,
		Before: before,
		After:  after,
	}
}
