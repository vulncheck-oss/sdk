package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Benchmark float64  `json:"_benchmark"`
	Data      UserData `json:"data"`
	Type      string   `json:"type"`
	Error     bool     `json:"error"`
	Errors    []string `json:"errors"`
	Success   bool     `json:"success"`
}

// https://docs.vulncheck.com/api/logout
func (c *Client) Logout() (responseJSON *Response, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/logout", nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("errors: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

// Strings representation of the response
func (r Response) String() string {
	return fmt.Sprintf("Benchmark: %v\nResponse: %v\n", r.Benchmark, r.Success)
}
