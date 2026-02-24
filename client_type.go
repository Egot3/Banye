package banye

import (
	"fmt"
	"net/http"
)

type Client struct {
	Base       *http.Client
	middleware []Middleware
}

type Middleware func(*http.Request) (*http.Request, error)

type MiddlewareTripper struct {
	Next   http.RoundTripper
	Before func(*http.Request) error
	After  func(*http.Request, *http.Response, error) error
}

func (m *MiddlewareTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.Before != nil {
		if err := m.Before(req); err != nil {
			return nil, err
		}
	}

	resp, err := m.Next.RoundTrip(req)

	if m.After != nil {
		err = m.After(req, resp, err)
	}
	return resp, err
}

// Append a after middleware to client
func (c *Client) Use(mw ...Middleware) {
	c.middleware = append(c.middleware, mw...)
}

// net/http client.Do but with middleware
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	var err error
	for _, mw := range c.middleware {
		req, err = mw(req)
		if err != nil {
			return nil, fmt.Errorf("Middleware err: %v", err)
		}
	}
	return c.Base.Do(req) //pseudo-recursion moment
}
