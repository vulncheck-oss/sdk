package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/octoper/go-ray"
)

type CpeResponse struct {
	Benchmark float64 `json:"_benchmark"`
	Meta      struct {
		Cpe       string `json:"cpe"`
		CpeStruct struct {
			Part      string `json:"part"`
			Vendor    string `json:"vendor"`
			Product   string `json:"product"`
			Version   string `json:"version"`
			Update    string `json:"update"`
			Edition   string `json:"edition"`
			Language  string `json:"language"`
			SwEdition string `json:"sw_edition"`
			TargetSw  string `json:"target_sw"`
			TargetHw  string `json:"target_hw"`
			Other     string `json:"other"`
		} `json:"cpe_struct"`
		Timestamp      string  `json:"timestamp"`
		TotalDocuments float64 `json:"total_documents"`
	} `json:"_meta"`
	Data []string `json:"data"`
}

// https://docs.vulncheck.com/api/cpe
func (c *Client) GetCpe(cpe string) (responseJSON *CpeResponse, err error) {
	resp, err := c.Request("GET", "/v3/cpe?cpe="+cpe)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	ray.Ray(responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r CpeResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// Returns the data from the response
func (r CpeResponse) GetData() []string {
	return r.Data
}
