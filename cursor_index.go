package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type IndexMetaCursor struct {
	Timestamp      string                `json:"timestamp"`
	Index          string                `json:"index"`
	Limit          int                   `json:"limit"`
	TotalDocuments int                   `json:"total_documents"`
	Sort           string                `json:"sort"`
	Parameters     []IndexMetaParameters `json:"parameters"`
	Order          string                `json:"order"`
	Cursor         string                `json:"cursor"`
	NextCursor     string                `json:"next_cursor"`
	PrevCursor     string                `json:"prev_cursor"` // this may not exist
}

type IndexCursorResponse struct {
	Benchmark float64       `json:"_benchmark"`
	Meta      IndexMeta     `json:"_meta"`
	Data      []interface{} `json:"data"`
}

// https://docs.vulncheck.com/api/cursor
func (c *Client) GetCursorIndex(index string, cursor string, queryParameters ...IndexQueryParameters) (responseJSON *IndexCursorResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape(index)+"/cursor", nil)

	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	SetIndexQueryParameters(query, queryParameters...)

	if cursor != "" {
		query.Add("cursor", cursor)
	}
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
func (r IndexCursorResponse) String() string {
	return fmt.Sprintf("Benchmark: %f\nMeta: %v\nData: %v\n", r.Benchmark, r.Meta, r.Data)
}

// GetData - Returns the data from the response
func (r IndexCursorResponse) GetData() []interface{} {
	return r.Data
}
