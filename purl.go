package sdk

import (
	"encoding/json"
	"fmt"
)

type PurlMeta struct {
	Namespace  string   `json:"namespace"`
	Name       string   `json:"name"`
	Version    string   `json:"version"`
	Qualifiers []string `json:"qualifiers"`
	Subpath    string   `json:"subpath"`
	Type       string   `json:"type"`
}

type PurlVulnerability struct {
	Detection    string `json:"detection"`
	FixedVersion string `json:"fixed_version"`
}

type PurlData struct {
	Cves            []string            `json:"cves"`
	Vulnerabilities []PurlVulnerability `json:"vulnerabilities"`
}

type PurlResponse struct {
	Benchmark float64 `json:"_benchmark"`
	Meta      struct {
		PurlMeta       PurlMeta `json:"purl_struct"`
		Timestamp      string   `json:"timestamp"`
		TotalDocuments float64  `json:"total_documents"`
	} `json:"_meta"`
	Data PurlData `json:"data"`
}

// GetPurl https://docs.vulncheck.com/api/purl
func (c *Client) GetPurl(purl string) (responseJSON *PurlResponse, err error) {
	resp, err := c.Query("purl", purl).Request("GET", "/v3/purl")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

// Strings representation of the response
func (r PurlResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// GetData Returns the data from the response
func (r PurlResponse) GetData() PurlData {
	return r.Data
}

// GetPurlMeta Returns the PurlMeta from the Metadata
func (r PurlResponse) GetPurlMeta() PurlMeta {
	return r.Meta.PurlMeta
}

// GetCves Cves Returns the list of CVEs associated with the purl
func (r PurlResponse) GetCves() []string {
	return r.Data.Cves
}

// GetVulnerabilities Vulnerabilities Returns the list of vulnerabilities associated with the purl
func (r PurlResponse) GetVulnerabilities() []PurlVulnerability {
	return r.Data.Vulnerabilities
}
