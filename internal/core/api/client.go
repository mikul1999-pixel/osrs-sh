package api

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	http      *http.Client
	limiter   <-chan time.Time
	userAgent string
}

type Options struct {
	Timeout   time.Duration
	UserAgent string
	Rate      time.Duration
}

func NewClient(opts Options) *Client {
	// Conventions:
	// - clear user agent with contact
	// - rate limit to 1 req per second

	if opts.Timeout == 0 {
		opts.Timeout = 5 * time.Second
	}
	if opts.UserAgent == "" {
		opts.UserAgent = "osrs-sh/0.1 (github.com/mikul1999-pixel/osrs-sh)"
	}
	if opts.Rate == 0 {
		opts.Rate = 1 * time.Second
	}

	return &Client{
		http: &http.Client{
			Timeout: opts.Timeout,
		},
		limiter:   time.Tick(opts.Rate), // 1 request per second
		userAgent: opts.UserAgent,
	}
}

// do executes an http request
func (c *Client) do(req *http.Request) ([]byte, int, error) {
	<-c.limiter
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}
