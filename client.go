package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Url        string
	Token      string
	HttpClient *http.Client
}

type MetaError struct {
	Error  bool     `json:"error"`
	Errors []string `json:"errors"`
}

var ErrorUnauthorized = fmt.Errorf("unauthorized")

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

// SetAuthHeader Sets the Authorization header for the request
func (c *Client) SetAuthHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
}

func (c *Client) Request(method string, url string) (*http.Response, error) {

	if c.HttpClient == nil {
		c.HttpClient = &http.Client{}

	}
	req, err := http.NewRequest(method, c.GetUrl()+url, nil)
	if err != nil {
		return nil, err
	}

	c.SetAuthHeader(req)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		return nil, ErrorUnauthorized
	}

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)
		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	return resp, nil

}
