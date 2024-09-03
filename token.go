package sdk

import "encoding/json"

type TokenLocation struct {
	City        string
	Region      string
	Country     string
	TimeZone    string
	CountryName string
}

type TokenData struct {
	ID        string
	Token     string
	Source    string
	Label     string
	Icon      string
	Ip        string
	Agent     string
	IsCurrent bool
	IsIssused bool
	Location  TokenLocation
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TokenResult struct {
	Benchmark float64 `json:"_benchmark"`
	Meta      struct {
		Timestamp      string  `json:"timestamp"`
		Limit          int     `json:"limit"`
		TotalDocuments float64 `json:"total_documents"`
		Sort           string  `json:"sort"`
		Order          string  `json:"order"`
		Page           int     `json:"page"`
		TotalPages     int     `json:"total_pages"`
		MaxPages       int     `json:"max_pages"`
		FirstItem      int     `json:"first_item"`
		LastItem       int     `json:"last_item"`
	} `json:"_meta"`
	Data []TokenData
}

type TokenResponse struct {
	Benchmark float64   `json:"_benchmark"`
	Message   string    `json:"message"`
	Success   bool      `json:"success"`
	Type      string    `json:"type"`
	Data      TokenData `json:"data"`
}

func (c *Client) GetTokens() (responseJSON *TokenResult, err error) {
	resp, err := c.Request("GET", "/token")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

func (c *Client) CreateToken(label string) (responseJSON *TokenResponse, err error) {
	resp, err := c.Form("label", label).Form("icon", "i-vc-icon").Request("POST", "/token")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

func (c *Client) DeleteToken(ID string) (responseJSON *TokenResponse, err error) {
	resp, err := c.Request("DELETE", "/token/"+ID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)
	return responseJSON, nil
}

// GetData - Returns the data from the response
func (r TokenResult) GetData() []TokenData {
	return r.Data
}
