package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BackupFile struct {
	Filename      string `json:"filename"`
	Sha256        string `json:"sha256"`
	DateAdded     string `json:"date_added"`
	URL           string `json:"url"`
	URLTtlMinutes int    `json:"url_ttl_minutes"`
	URLExpires    string `json:"url_expires"`
}

type BackupResponse struct {
	Benchmark float64 `json:"_benchmark"`
	Meta      struct {
		Timestamp string `json:"timestamp"`
		Index     string `json:"index"`
	} `json:"_meta"`
	Data []BackupFile `json:"data"`
}

// https://docs.vulncheck.com/api/backup
func (c *Client) GetIndexBackup(index string) (responseJSON BackupResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/backup/"+index, nil)
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

		return responseJSON, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

// Strings representation of the response
func (r BackupResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// Returns the data from the response
func (r BackupResponse) GetData() []BackupFile {
	return r.Data
}

// Returns the list of filenames associated with the backup
func (r BackupResponse) Filenames() []string {
	var filenames []string
	for _, v := range r.Data {
		filenames = append(filenames, v.Filename)
	}
	return filenames
}

// Returns the list of URLs associated with the backup
func (r BackupResponse) Urls() []string {
	var urls []string
	for _, v := range r.Data {
		urls = append(urls, v.URL)
	}
	return urls
}
