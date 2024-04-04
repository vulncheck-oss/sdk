package sdk

import (
	"encoding/json"
	"fmt"
)

type SbomPagination struct {
	Timestamp      string `json:"timestamp"`
	Limit          int    `json:"limit"`
	TotalDocuments int    `json:"total_documents"`
	Sort           string `json:"sort"`
	Order          string `json:"order"`
	Page           int    `json:"page"`
	TotalPages     int    `json:"total_pages"`
	MaxPages       int    `json:"max_pages"`
	FirstItem      int    `json:"first_item"`
	LastItem       int    `json:"last_item"`
}

type SbomMeta struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Repository string `json:"repository"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type SbomListResponse struct {
	Benchmark float64        `json:"_benchmark"`
	Meta      SbomPagination `json:"_meta"`
	Data      []SbomMeta     `json:"data"`
}

type SbomScanMeta struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type SbomScansResponse struct {
	Benchmark float64        `json:"_benchmark"`
	Meta      SbomPagination `json:"_meta"`
	Data      []SbomScanMeta `json:"data"`
}

// https://docs.vulncheck.com/api/sbom
func (c *Client) SbomList() (responseJSON *SbomListResponse, err error) {
	resp, err := c.Request("GET", "/v3/sbom")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r SbomListResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// GetData Returns the data from the response
func (r SbomListResponse) GetData() []SbomMeta {
	return r.Data
}

func (c *Client) SbomScans(id string) (responseJSON *SbomScansResponse, err error) {
	resp, err := c.Request("GET", fmt.Sprintf("/v3/sbom/%s/scan", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r SbomScansResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// GetData Returns the data from the response
func (r SbomScansResponse) GetData() []SbomScanMeta {
	return r.Data
}
