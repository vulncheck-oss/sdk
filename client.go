package sdk

import "net/http"

type Client struct {
	Env   string
	Token string
}

type MetaError struct {
	Error  bool     `json:"error"`
	Errors []string `json:"errors"`
}

var (
	TestURL  = "https://api.vulncheck.ai/v3"
	StageURL = "https://api.vulncheck.ai/v3"
	ProdURL  = "https://api.vulncheck.com/v3"
)

func Connect(token string, env ...string) *Client {
	if len(env) == 0 {
		return &Client{Token: token, Env: "prod"}
	}

	return &Client{Token: token, Env: env[0]}
}

func (c *Client) GetToken() string {
	return c.Token
}

func (c *Client) SetToken(token string) {
	c.Token = token
}

func (c *Client) SetEnv(env string) {
	c.Env = env
}

func (c *Client) GetEnv() string {
	return c.Env
}

func (c *Client) GetBaseURL() string {
	if c.Env == "stage" {
		return StageURL
	}
	return ProdURL
}

// Sets the Authorization header for the request
func (c *Client) SetAuthHeader(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
}
