package sdk

import (
	"io"
	"net/http"
)

// https://docs.vulncheck.com/api/pdns
func (c *Client) GetPdns() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/pdns/vulncheck-c2", nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return string(body), nil
}
