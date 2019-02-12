package webapi

import "net/http"

// httpClient defines the minimal interface needed for an http.Client to be implemented.
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct{
	token string
	httpClient httpClient
}

const CF_URL = "https://g.codefresh.io/api/"
func New(token string) *Client {
	client := &Client{token:token,httpClient: &http.Client{}}
	return client
}

func (c *Client) Sum (x int,y int) int {
	return x+y
}