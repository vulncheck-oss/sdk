package sdk

import "net/http"

type Client struct {
	Url   string
	Token string
}

type MetaError struct {
	Error  bool     `json:"error"`
	Errors []string `json:"errors"`
}

func Connect(url string, token string) *Client {

	return &Client{Url: url, Token: token}
}

func (c *Client) GetToken() string {
	return c.Token
}

func (c *Client) SetToken(token string) {
	c.Token = token
}

func (c *Client) SetUrl(env string) {
	c.Url = env
}

func (c *Client) GetUrl() string {
	return c.Url
}

// Sets the Authorization header for the request
func (c *Client) SetAuthHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
}
