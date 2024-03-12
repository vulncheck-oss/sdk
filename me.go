package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserData struct {
	ID                   int         `json:"ID"`
	Email                string      `json:"Email"`
	Name                 string      `json:"Name"`
	Avatar               string      `json:"Avatar"`
	Payload              interface{} `json:"Payload"`
	Stripe               string      `json:"Stripe"`
	Terms                bool        `json:"Terms"`
	Roles                []string    `json:"Roles"`
	Settings             interface{} `json:"Settings"`
	Providers            interface{} `json:"Providers"`
	Teams                interface{} `json:"Teams"`
	URL                  string      `json:"URL"`
	Initials             string      `json:"Initials"`
	TrialDays            int         `json:"TrialDays"`
	Pivot                interface{} `json:"Pivot"`
	OrgRoles             interface{} `json:"OrgRoles"`
	CurrentToken         interface{} `json:"CurrentToken"`
	ActivatedAt          interface{} `json:"ActivatedAt"`
	OrgUsers             interface{} `json:"OrgUsers"`
	Orgs                 interface{} `json:"Orgs"`
	CurrentOrg           interface{} `json:"CurrentOrg"`
	OrgID                interface{} `json:"OrgID"`
	Org                  interface{} `json:"Org"`
	IsServiceUser        bool        `json:"IsServiceUser"`
	HasEmployeeRole      bool        `json:"HasEmployeeRole"`
	HasEmployeeAdminRole bool        `json:"HasEmployeeAdminRole"`
	HasOrgManagerRole    bool        `json:"HasOrgManagerRole"`
	HasTrial             bool        `json:"HasTrial"`
	HasInitial           bool        `json:"HasInitial"`
	HasVuln              bool        `json:"HasVuln"`
	HasAgent             bool        `json:"HasAgent"`
	HasSbom              bool        `json:"HasSbom"`
	OnlyCommunity        bool        `json:"OnlyCommunity"`
	CreatedAt            string      `json:"created_at"`
	UpdatedAt            string      `json:"updated_at"`
}

type UserResponse struct {
	Benchmark float64  `json:"_benchmark"`
	Data      UserData `json:"data"`
}

// https://docs.vulncheck.com/api/me
func (c *Client) GetMe() (responseJSON UserResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/me", nil)
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
func (r UserResponse) String() string {
	return fmt.Sprintf("Benchmark: %v\nData: %v\n", r.Benchmark, r.Data)
}
