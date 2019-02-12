package webapi

type Client struct{
	token string
}

func New(token string) *Client {
	client := &Client{token:token}
	return client
}

func (c *Client) Sum (x int,y int) int {
	return x+y
}