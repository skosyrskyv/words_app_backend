package httpclient

import (
	"net/http"
	"time"
)

type Client struct {
	c *http.Client
}

func NewHTTPClient() *Client {
	cl := http.Client{
		Timeout: time.Second * 30,
	}

	return &Client{
		c: &cl,
	}
}
