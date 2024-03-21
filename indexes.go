package sdk

import (
	"encoding/json"
	"fmt"
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

// GetIndexes https://docs.vulncheck.com/api/indexes
func (c *Client) GetIndexes() (responseJSON *IndexesResponse, err error) {

	resp, err := c.Request("GET", "/v3/index")
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r IndexesResponse) String() string {
	return fmt.Sprintf("Benchmark: %v\nData: %v\n", r.Benchmark, r.Data)
}

// GetData - Returns the data from the response
func (r IndexesResponse) GetData() []IndexesMeta {
	return r.Data
}
