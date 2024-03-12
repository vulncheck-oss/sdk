package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type PurlMeta struct {
	Namespace  string      `json:"namespace"`
	Name       string      `json:"name"`
	Version    string      `json:"version"`
	Qualifiers interface{} `json:"qualifiers"`
	Subpath    string      `json:"subpath"`
	Type       string      `json:"type"`
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
		PurlStruct     PurlMeta `json:"purl_struct"`
		Timestamp      string   `json:"timestamp"`
		TotalDocuments float64  `json:"total_documents"`
	} `json:"_meta"`
	Data PurlData `json:"data"`
}

// https://docs.vulncheck.com/api/purl
func (c *Client) GetPurl(purl string) (responseJSON *PurlResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/purl", nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := url.Values{}
	query.Add("purl", purl)
	req.URL.RawQuery = query.Encode()

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
func (r PurlResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// Returns the data from the response
func (r PurlResponse) GetData() PurlData {
	return r.Data
}

// Returns the list of CVEs associated with the purl
func (r PurlResponse) Cves() []string {
	return r.Data.Cves
}

// Returns the list of vulnerabilities associated with the purl
func (r PurlResponse) Vulnerabilities() []PurlVulnerability {
	return r.Data.Vulnerabilities
}
