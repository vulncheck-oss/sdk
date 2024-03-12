package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IndexesMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Href        string `json:"href"`
}

type IndexesResponse struct {
	Benchmark float64       `json:"_benchmark"`
	Data      []IndexesMeta `json:"data"`
}

// https://docs.vulncheck.com/api/indexes
func (c *Client) GetIndexes() (responseJSON *IndexesResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/index", nil)
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

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

// Strings representation of the response
func (r IndexesResponse) String() string {
	return fmt.Sprintf("Benchmark: %v\nData: %v\n", r.Benchmark, r.Data)
}

// Returns the data from the response
func (r IndexesResponse) GetData() []IndexesMeta {
	return r.Data
}
