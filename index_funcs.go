package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/vulncheck-oss/sdk/pkg/client"
	"net/http"
	"net/url"
)

type GetIndexA10Response struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryA10 `json:"data"`
}

func (c *Client) GetIndexA10(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexA10Response, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("a10"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAbbResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryABBAdvisory `json:"data"`
}

func (c *Client) GetIndexAbb(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAbbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("abb"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAbbottResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryAbbott `json:"data"`
}

func (c *Client) GetIndexAbbott(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAbbottResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("abbott"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAbsoluteResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAbsolute `json:"data"`
}

func (c *Client) GetIndexAbsolute(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAbsoluteResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("absolute"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAcronisResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAcronis `json:"data"`
}

func (c *Client) GetIndexAcronis(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAcronisResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("acronis"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAdobeResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAdobeAdvisory `json:"data"`
}

func (c *Client) GetIndexAdobe(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAdobeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("adobe"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAlephResearchResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAlephResearch `json:"data"`
}

func (c *Client) GetIndexAlephResearch(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAlephResearchResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("aleph-research"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAlmaResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryAlmaLinuxUpdate `json:"data"`
}

func (c *Client) GetIndexAlma(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAlmaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("alma"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAlpineResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryAlpineLinuxSecDB `json:"data"`
}

func (c *Client) GetIndexAlpine(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAlpineResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("alpine"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAmazonResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryUpdate `json:"data"`
}

func (c *Client) GetIndexAmazon(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAmazonResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("amazon"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAmdResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAMD `json:"data"`
}

func (c *Client) GetIndexAmd(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAmdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("amd"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAmiResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAMI `json:"data"`
}

func (c *Client) GetIndexAmi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAmiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ami"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAndroidResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryAndroidAdvisory `json:"data"`
}

func (c *Client) GetIndexAndroid(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAndroidResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("android"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheActivemqResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryApacheActiveMQ `json:"data"`
}

func (c *Client) GetIndexApacheActivemq(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheActivemqResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-activemq"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheArchivaResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheArchiva `json:"data"`
}

func (c *Client) GetIndexApacheArchiva(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheArchivaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-archiva"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheArrowResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheArrow `json:"data"`
}

func (c *Client) GetIndexApacheArrow(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheArrowResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-arrow"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheCamelResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheCamel `json:"data"`
}

func (c *Client) GetIndexApacheCamel(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheCamelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-camel"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheCommonsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheCommons `json:"data"`
}

func (c *Client) GetIndexApacheCommons(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheCommonsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-commons"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheCouchdbResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheCouchDB `json:"data"`
}

func (c *Client) GetIndexApacheCouchdb(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheCouchdbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-couchdb"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheFlinkResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheFlink `json:"data"`
}

func (c *Client) GetIndexApacheFlink(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheFlinkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-flink"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheGuacamoleResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryApacheGuacamole `json:"data"`
}

func (c *Client) GetIndexApacheGuacamole(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheGuacamoleResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-guacamole"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheHadoopResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApacheHadoop `json:"data"`
}

func (c *Client) GetIndexApacheHadoop(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheHadoopResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-hadoop"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheHttpResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryApacheHTTP `json:"data"`
}

func (c *Client) GetIndexApacheHttp(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheHttpResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-http"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheJspwikiResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheJSPWiki `json:"data"`
}

func (c *Client) GetIndexApacheJspwiki(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheJspwikiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-jspwiki"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheKafkaResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheKafka `json:"data"`
}

func (c *Client) GetIndexApacheKafka(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheKafkaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-kafka"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheLoggingservicesResponse struct {
	Benchmark float64                                `json:"_benchmark"`
	Meta      IndexMeta                              `json:"_meta"`
	Data      []client.AdvisoryApacheLoggingServices `json:"data"`
}

func (c *Client) GetIndexApacheLoggingservices(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheLoggingservicesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-loggingservices"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheNifiResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryApacheNiFi `json:"data"`
}

func (c *Client) GetIndexApacheNifi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheNifiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-nifi"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheOfbizResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheOFBiz `json:"data"`
}

func (c *Client) GetIndexApacheOfbiz(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheOfbizResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-ofbiz"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheOpenmeetingsResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryApacheOpenMeetings `json:"data"`
}

func (c *Client) GetIndexApacheOpenmeetings(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheOpenmeetingsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-openmeetings"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheOpenofficeResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryApacheOpenOffice `json:"data"`
}

func (c *Client) GetIndexApacheOpenoffice(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheOpenofficeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-openoffice"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApachePulsarResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApachePulsar `json:"data"`
}

func (c *Client) GetIndexApachePulsar(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApachePulsarResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-pulsar"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheShiroResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheShiro `json:"data"`
}

func (c *Client) GetIndexApacheShiro(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheShiroResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-shiro"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheSparkResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheSpark `json:"data"`
}

func (c *Client) GetIndexApacheSpark(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheSparkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-spark"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheStrutsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApacheStruts `json:"data"`
}

func (c *Client) GetIndexApacheStruts(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheStrutsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-struts"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheSubversionResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryApacheSubversion `json:"data"`
}

func (c *Client) GetIndexApacheSubversion(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheSubversionResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-subversion"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheSupersetResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryApacheSuperset `json:"data"`
}

func (c *Client) GetIndexApacheSuperset(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheSupersetResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-superset"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheTomcatResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApacheTomcat `json:"data"`
}

func (c *Client) GetIndexApacheTomcat(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheTomcatResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-tomcat"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexApacheZookeeperResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryApacheZooKeeper `json:"data"`
}

func (c *Client) GetIndexApacheZookeeper(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexApacheZookeeperResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apache-zookeeper"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAppcheckResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAppCheck `json:"data"`
}

func (c *Client) GetIndexAppcheck(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAppcheckResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("appcheck"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAppgateResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAppgate `json:"data"`
}

func (c *Client) GetIndexAppgate(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAppgateResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("appgate"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAppleResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAppleAdvisory `json:"data"`
}

func (c *Client) GetIndexApple(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAppleResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("apple"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexArchResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryArchIssue `json:"data"`
}

func (c *Client) GetIndexArch(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexArchResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("arch"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAristaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryArista `json:"data"`
}

func (c *Client) GetIndexArista(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAristaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("arista"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexArubaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryAruba `json:"data"`
}

func (c *Client) GetIndexAruba(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexArubaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("aruba"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAsrgResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryASRG `json:"data"`
}

func (c *Client) GetIndexAsrg(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAsrgResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("asrg"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAssetnoteResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryAssetNote `json:"data"`
}

func (c *Client) GetIndexAssetnote(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAssetnoteResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("assetnote"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAsteriskResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAsterisk `json:"data"`
}

func (c *Client) GetIndexAsterisk(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAsteriskResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("asterisk"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAsusResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryAsus `json:"data"`
}

func (c *Client) GetIndexAsus(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAsusResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("asus"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAtlassianResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryAtlassianAdvisory `json:"data"`
}

func (c *Client) GetIndexAtlassian(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAtlassianResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("atlassian"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAuscertResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAusCert `json:"data"`
}

func (c *Client) GetIndexAuscert(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAuscertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("auscert"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAutodeskResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAutodesk `json:"data"`
}

func (c *Client) GetIndexAutodesk(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAutodeskResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("autodesk"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAvayaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryAvaya `json:"data"`
}

func (c *Client) GetIndexAvaya(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAvayaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("avaya"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAvevaResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAVEVAAdvisory `json:"data"`
}

func (c *Client) GetIndexAveva(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAvevaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("aveva"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAvigilonResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAvigilon `json:"data"`
}

func (c *Client) GetIndexAvigilon(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAvigilonResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("avigilon"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAwsResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAWS `json:"data"`
}

func (c *Client) GetIndexAws(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAwsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("aws"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexAxisResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryAxis `json:"data"`
}

func (c *Client) GetIndexAxis(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexAxisResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("axis"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBandrResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryBandr `json:"data"`
}

func (c *Client) GetIndexBandr(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBandrResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bandr"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBaxterResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryBaxterAdvisory `json:"data"`
}

func (c *Client) GetIndexBaxter(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBaxterResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("baxter"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBbraunResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryBBraunAdvisory `json:"data"`
}

func (c *Client) GetIndexBbraun(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBbraunResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bbraun"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBdResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryBectonDickinsonAdvisory `json:"data"`
}

func (c *Client) GetIndexBd(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bd"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBduResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryBDUAdvisory `json:"data"`
}

func (c *Client) GetIndexBdu(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBduResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bdu"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBeckhoffResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryBeckhoffAdvisory `json:"data"`
}

func (c *Client) GetIndexBeckhoff(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBeckhoffResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("beckhoff"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBeldenResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryBeldenAdvisory `json:"data"`
}

func (c *Client) GetIndexBelden(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBeldenResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("belden"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBeyondTrustResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryBeyondTrust `json:"data"`
}

func (c *Client) GetIndexBeyondTrust(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBeyondTrustResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("beyond-trust"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBinarlyResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryBinarly `json:"data"`
}

func (c *Client) GetIndexBinarly(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBinarlyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("binarly"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBitdefenderResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryBitDefender `json:"data"`
}

func (c *Client) GetIndexBitdefender(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBitdefenderResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bitdefender"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBlackberryResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryBlackBerry `json:"data"`
}

func (c *Client) GetIndexBlackberry(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBlackberryResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("blackberry"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBlsResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryBLS `json:"data"`
}

func (c *Client) GetIndexBls(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBlsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bls"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBoschResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryBoschAdvisory `json:"data"`
}

func (c *Client) GetIndexBosch(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBoschResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("bosch"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBostonScientificResponse struct {
	Benchmark float64                                   `json:"_benchmark"`
	Meta      IndexMeta                                 `json:"_meta"`
	Data      []client.AdvisoryBostonScientificAdvisory `json:"data"`
}

func (c *Client) GetIndexBostonScientific(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBostonScientificResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("boston-scientific"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexBotnetsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryBotnet `json:"data"`
}

func (c *Client) GetIndexBotnets(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexBotnetsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("botnets"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCaCyberCentreResponse struct {
	Benchmark float64                                `json:"_benchmark"`
	Meta      IndexMeta                              `json:"_meta"`
	Data      []client.AdvisoryCACyberCentreAdvisory `json:"data"`
}

func (c *Client) GetIndexCaCyberCentre(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCaCyberCentreResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ca-cyber-centre"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCanvasResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryCanvasExploit `json:"data"`
}

func (c *Client) GetIndexCanvas(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCanvasResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("canvas"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCarestreamResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryCarestreamAdvisory `json:"data"`
}

func (c *Client) GetIndexCarestream(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCarestreamResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("carestream"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCargoResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexCargo(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCargoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cargo"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCarrierResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryCarrier `json:"data"`
}

func (c *Client) GetIndexCarrier(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCarrierResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("carrier"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCblMarinerResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryCBLMariner `json:"data"`
}

func (c *Client) GetIndexCblMariner(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCblMarinerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cbl-mariner"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCentosResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryCESA `json:"data"`
}

func (c *Client) GetIndexCentos(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCentosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("centos"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCertBeResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryCertBE `json:"data"`
}

func (c *Client) GetIndexCertBe(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCertBeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cert-be"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCertUaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryCertUA `json:"data"`
}

func (c *Client) GetIndexCertUa(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCertUaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cert-ua"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCerteuResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCERTEUAdvisory `json:"data"`
}

func (c *Client) GetIndexCerteu(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCerteuResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("certeu"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCertfrResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCertFRAdvisory `json:"data"`
}

func (c *Client) GetIndexCertfr(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCertfrResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("certfr"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexChainguardResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryChainGuard `json:"data"`
}

func (c *Client) GetIndexChainguard(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexChainguardResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("chainguard"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCheckpointResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryCheckPoint `json:"data"`
}

func (c *Client) GetIndexCheckpoint(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCheckpointResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("checkpoint"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexChromeResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryChrome `json:"data"`
}

func (c *Client) GetIndexChrome(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexChromeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("chrome"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCisaAlertsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryCISAAlert `json:"data"`
}

func (c *Client) GetIndexCisaAlerts(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCisaAlertsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cisa-alerts"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCisaKevResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryKEVCatalogVulnerability `json:"data"`
}

func (c *Client) GetIndexCisaKev(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCisaKevResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cisa-kev"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCiscoResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryCiscoAdvisory `json:"data"`
}

func (c *Client) GetIndexCisco(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCiscoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cisco"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCiscoTalosResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryTalosAdvisory `json:"data"`
}

func (c *Client) GetIndexCiscoTalos(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCiscoTalosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cisco-talos"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCitrixResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCitrixAdvisory `json:"data"`
}

func (c *Client) GetIndexCitrix(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCitrixResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("citrix"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexClarotyResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryClarotyVulnerability `json:"data"`
}

func (c *Client) GetIndexClaroty(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexClarotyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("claroty"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCloudbeesResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryCloudBees `json:"data"`
}

func (c *Client) GetIndexCloudbees(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCloudbeesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cloudbees"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCloudvulndbResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryCloudVulnDBAdvisory `json:"data"`
}

func (c *Client) GetIndexCloudvulndb(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCloudvulndbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cloudvulndb"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCnnvdResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCNNVDEntryJSON `json:"data"`
}

func (c *Client) GetIndexCnnvd(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCnnvdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cnnvd"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCnvdBulletinsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryCNVDBulletin `json:"data"`
}

func (c *Client) GetIndexCnvdBulletins(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCnvdBulletinsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cnvd-bulletins"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCnvdFlawsResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryCNVDFlaw `json:"data"`
}

func (c *Client) GetIndexCnvdFlaws(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCnvdFlawsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cnvd-flaws"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCocoapodsResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexCocoapods(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCocoapodsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cocoapods"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCodesysResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryCodesysAdvisory `json:"data"`
}

func (c *Client) GetIndexCodesys(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCodesysResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("codesys"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCompassSecurityResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryCompassSecurity `json:"data"`
}

func (c *Client) GetIndexCompassSecurity(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCompassSecurityResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("compass-security"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexComposerResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexComposer(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexComposerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("composer"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexConanResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexConan(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexConanResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("conan"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCrestronResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryCrestron `json:"data"`
}

func (c *Client) GetIndexCrestron(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCrestronResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("crestron"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCurlResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryCurl `json:"data"`
}

func (c *Client) GetIndexCurl(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCurlResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("curl"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexCweResponse struct {
	Benchmark float64         `json:"_benchmark"`
	Meta      IndexMeta       `json:"_meta"`
	Data      []client.ApiCWE `json:"data"`
}

func (c *Client) GetIndexCwe(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexCweResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("cwe"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexDahuaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryDahua `json:"data"`
}

func (c *Client) GetIndexDahua(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexDahuaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dahua"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexDassaultResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryDassault `json:"data"`
}

func (c *Client) GetIndexDassault(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexDassaultResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dassault"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexDebianResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryVulnerableDebianPackage `json:"data"`
}

func (c *Client) GetIndexDebian(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexDebianResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("debian"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexDebianDsaResponse struct {
	Benchmark float64                                 `json:"_benchmark"`
	Meta      IndexMeta                               `json:"_meta"`
	Data      []client.AdvisoryDebianSecurityAdvisory `json:"data"`
}

func (c *Client) GetIndexDebianDsa(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexDebianDsaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("debian-dsa"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexDellResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryDell `json:"data"`
}

func (c *Client) GetIndexDell(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexDellResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dell"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexDeltaResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryDeltaAdvisory `json:"data"`
}

func (c *Client) GetIndexDelta(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexDeltaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("delta"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexDotcmsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryDotCMS `json:"data"`
}

func (c *Client) GetIndexDotcms(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexDotcmsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dotcms"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexDragosResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryDragosAdvisory `json:"data"`
}

func (c *Client) GetIndexDragos(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexDragosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("dragos"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexDraytekResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryDraytek `json:"data"`
}

func (c *Client) GetIndexDraytek(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexDraytekResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("draytek"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexEatonResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryEatonAdvisory `json:"data"`
}

func (c *Client) GetIndexEaton(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexEatonResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("eaton"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexElasticResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryElastic `json:"data"`
}

func (c *Client) GetIndexElastic(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexElasticResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("elastic"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexEmersonResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryEmersonAdvisory `json:"data"`
}

func (c *Client) GetIndexEmerson(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexEmersonResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("emerson"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexEolResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryEOLReleaseData `json:"data"`
}

func (c *Client) GetIndexEol(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexEolResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("eol"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexEpssResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.ApiEPSSData `json:"data"`
}

func (c *Client) GetIndexEpss(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexEpssResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("epss"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexExodusIntelResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryExodusIntel `json:"data"`
}

func (c *Client) GetIndexExodusIntel(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexExodusIntelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("exodus-intel"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexExploitChainsResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.ApiExploitChain `json:"data"`
}

func (c *Client) GetIndexExploitChains(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexExploitChainsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("exploit-chains"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexExploitsResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.ApiExploitV3Result `json:"data"`
}

func (c *Client) GetIndexExploits(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexExploitsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("exploits"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexFSecureResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryFSecure `json:"data"`
}

func (c *Client) GetIndexFSecure(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexFSecureResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("f-secure"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexFastlyResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryFastly `json:"data"`
}

func (c *Client) GetIndexFastly(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexFastlyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("fastly"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexFedoraResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryUpdate `json:"data"`
}

func (c *Client) GetIndexFedora(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexFedoraResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("fedora"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexFilecloudResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryFileCloud `json:"data"`
}

func (c *Client) GetIndexFilecloud(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexFilecloudResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("filecloud"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexForgerockResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryForgeRock `json:"data"`
}

func (c *Client) GetIndexForgerock(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexForgerockResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("forgerock"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexFortinetResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryFortinetAdvisory `json:"data"`
}

func (c *Client) GetIndexFortinet(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexFortinetResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("fortinet"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexFreebsdResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAdvisory `json:"data"`
}

func (c *Client) GetIndexFreebsd(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexFreebsdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("freebsd"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGallagherResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryGallagher `json:"data"`
}

func (c *Client) GetIndexGallagher(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGallagherResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gallagher"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGcpResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryGCP `json:"data"`
}

func (c *Client) GetIndexGcp(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGcpResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gcp"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGeGasResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryGEGas `json:"data"`
}

func (c *Client) GetIndexGeGas(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGeGasResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ge-gas"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGeHealthcareResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryGEHealthcareAdvisory `json:"data"`
}

func (c *Client) GetIndexGeHealthcare(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGeHealthcareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ge-healthcare"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGemResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexGem(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGemResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gem"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGenetecResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryGenetec `json:"data"`
}

func (c *Client) GetIndexGenetec(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGenetecResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("genetec"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGigabyteResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryGigabyte `json:"data"`
}

func (c *Client) GetIndexGigabyte(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGigabyteResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gigabyte"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGiteeExploitsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryGiteeExploit `json:"data"`
}

func (c *Client) GetIndexGiteeExploits(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGiteeExploitsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gitee-exploits"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGithubExploitsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryGitHubExploit `json:"data"`
}

func (c *Client) GetIndexGithubExploits(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGithubExploitsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("github-exploits"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGitlabAdvisoriesCommunityResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryGitlabAdvisory `json:"data"`
}

func (c *Client) GetIndexGitlabAdvisoriesCommunity(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGitlabAdvisoriesCommunityResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gitlab-advisories-community"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGitlabExploitsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryGitLabExploit `json:"data"`
}

func (c *Client) GetIndexGitlabExploits(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGitlabExploitsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gitlab-exploits"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGnutlsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryGnuTLS `json:"data"`
}

func (c *Client) GetIndexGnutls(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGnutlsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("gnutls"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGolangResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexGolang(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGolangResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("golang"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGoogle0dayItwResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryITWExploit `json:"data"`
}

func (c *Client) GetIndexGoogle0dayItw(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGoogle0dayItwResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("google0day-itw"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGrafanaResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryGrafana `json:"data"`
}

func (c *Client) GetIndexGrafana(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGrafanaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("grafana"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexGreynoiseMetadataResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryGreyNoiseDetection `json:"data"`
}

func (c *Client) GetIndexGreynoiseMetadata(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexGreynoiseMetadataResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("greynoise-metadata"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHackageResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexHackage(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHackageResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hackage"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHarmonyosResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHarmonyOS `json:"data"`
}

func (c *Client) GetIndexHarmonyos(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHarmonyosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("harmonyos"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHashicorpResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHashiCorp `json:"data"`
}

func (c *Client) GetIndexHashicorp(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHashicorpResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hashicorp"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHaskellSadbResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryHaskellSADBAdvisory `json:"data"`
}

func (c *Client) GetIndexHaskellSadb(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHaskellSadbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("haskell-sadb"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHclResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryHCL `json:"data"`
}

func (c *Client) GetIndexHcl(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHclResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hcl"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHexResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexHex(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHexResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hex"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHikvisionResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHIKVision `json:"data"`
}

func (c *Client) GetIndexHikvision(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHikvisionResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hikvision"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHillromResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryHillromAdvisory `json:"data"`
}

func (c *Client) GetIndexHillrom(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHillromResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hillrom"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHitachiResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryHitachi `json:"data"`
}

func (c *Client) GetIndexHitachi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHitachiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hitachi"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHitachiEnergyResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryHitachiEnergy `json:"data"`
}

func (c *Client) GetIndexHitachiEnergy(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHitachiEnergyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hitachi-energy"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHkcertResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryHKCert `json:"data"`
}

func (c *Client) GetIndexHkcert(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHkcertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hkcert"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHoneywellResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHoneywell `json:"data"`
}

func (c *Client) GetIndexHoneywell(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHoneywellResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("honeywell"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHpResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryHP `json:"data"`
}

func (c *Client) GetIndexHp(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHpResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("hp"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHuaweiEulerosResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryHuaweiEulerOS `json:"data"`
}

func (c *Client) GetIndexHuaweiEuleros(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHuaweiEulerosResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("huawei-euleros"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHuaweiIpsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHuaweiIPS `json:"data"`
}

func (c *Client) GetIndexHuaweiIps(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHuaweiIpsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("huawei-ips"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexHuaweiPsirtResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryHuawei `json:"data"`
}

func (c *Client) GetIndexHuaweiPsirt(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexHuaweiPsirtResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("huawei-psirt"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIavaResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryIAVA `json:"data"`
}

func (c *Client) GetIndexIava(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIavaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("iava"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIbmResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryIBM `json:"data"`
}

func (c *Client) GetIndexIbm(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIbmResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ibm"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIdemiaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryIdemia `json:"data"`
}

func (c *Client) GetIndexIdemia(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIdemiaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("idemia"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIlAlertsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryIsraeliAlert `json:"data"`
}

func (c *Client) GetIndexIlAlerts(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIlAlertsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("il-alerts"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIlVulnerabilitiesResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryIsraeliVulnerability `json:"data"`
}

func (c *Client) GetIndexIlVulnerabilities(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIlVulnerabilitiesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("il-vulnerabilities"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIncibeResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryIncibeAdvisory `json:"data"`
}

func (c *Client) GetIndexIncibe(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIncibeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("incibe"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexInitialAccessResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.ApiInitialAccess `json:"data"`
}

func (c *Client) GetIndexInitialAccess(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexInitialAccessResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("initial-access"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexInitialAccessGitResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.ApiInitialAccess `json:"data"`
}

func (c *Client) GetIndexInitialAccessGit(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexInitialAccessGitResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("initial-access-git"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIntelResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryIntel `json:"data"`
}

func (c *Client) GetIndexIntel(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIntelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("intel"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIpintel10dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

func (c *Client) GetIndexIpintel10d(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIpintel10dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel10d"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIpintel30dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

func (c *Client) GetIndexIpintel30d(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIpintel30dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel30d"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIpintel3dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

func (c *Client) GetIndexIpintel3d(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIpintel3dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel3d"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIpintel90dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

func (c *Client) GetIndexIpintel90d(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIpintel90dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel90d"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIstioResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryIstio `json:"data"`
}

func (c *Client) GetIndexIstio(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIstioResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("istio"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexIvantiResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryIvanti `json:"data"`
}

func (c *Client) GetIndexIvanti(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexIvantiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ivanti"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexJenkinsResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryJenkins `json:"data"`
}

func (c *Client) GetIndexJenkins(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexJenkinsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jenkins"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexJetbrainsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryJetBrains `json:"data"`
}

func (c *Client) GetIndexJetbrains(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexJetbrainsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jetbrains"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexJfrogResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryJFrog `json:"data"`
}

func (c *Client) GetIndexJfrog(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexJfrogResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jfrog"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexJnjResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryJNJAdvisory `json:"data"`
}

func (c *Client) GetIndexJnj(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexJnjResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jnj"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexKasperskyIcsCertResponse struct {
	Benchmark float64                                   `json:"_benchmark"`
	Meta      IndexMeta                                 `json:"_meta"`
	Data      []client.AdvisoryKasperskyICSCERTAdvisory `json:"data"`
}

func (c *Client) GetIndexKasperskyIcsCert(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexKasperskyIcsCertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("kaspersky-ics-cert"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexKrcertSecurityNoticesResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryKRCertAdvisory `json:"data"`
}

func (c *Client) GetIndexKrcertSecurityNotices(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexKrcertSecurityNoticesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("krcert-security-notices"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexKrcertVulnerabilitiesResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryKRCertAdvisory `json:"data"`
}

func (c *Client) GetIndexKrcertVulnerabilities(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexKrcertVulnerabilitiesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("krcert-vulnerabilities"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexKubernetesResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryK8S `json:"data"`
}

func (c *Client) GetIndexKubernetes(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexKubernetesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("kubernetes"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexLenovoResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryLenovo `json:"data"`
}

func (c *Client) GetIndexLenovo(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexLenovoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("lenovo"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexLexmarkResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryLexmarkAdvisory `json:"data"`
}

func (c *Client) GetIndexLexmark(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexLexmarkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("lexmark"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexLgResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryLG `json:"data"`
}

func (c *Client) GetIndexLg(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexLgResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("lg"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexLibreOfficeResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryLibreOffice `json:"data"`
}

func (c *Client) GetIndexLibreOffice(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexLibreOfficeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("libre-office"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMFilesResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMFiles `json:"data"`
}

func (c *Client) GetIndexMFiles(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMFilesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("m-files"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMacertResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMACert `json:"data"`
}

func (c *Client) GetIndexMacert(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMacertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("macert"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexManageengineResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryManageEngineAdvisory `json:"data"`
}

func (c *Client) GetIndexManageengine(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexManageengineResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("manageengine"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMavenResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexMaven(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMavenResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("maven"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMbedTlsResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryMbedTLS `json:"data"`
}

func (c *Client) GetIndexMbedTls(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMbedTlsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mbed-tls"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMediatekResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMediatek `json:"data"`
}

func (c *Client) GetIndexMediatek(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMediatekResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mediatek"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMedtronicResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryMedtronicAdvisory `json:"data"`
}

func (c *Client) GetIndexMedtronic(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMedtronicResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("medtronic"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMendixResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMendix `json:"data"`
}

func (c *Client) GetIndexMendix(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMendixResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mendix"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMetasploitResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryMetasploitExploit `json:"data"`
}

func (c *Client) GetIndexMetasploit(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMetasploitResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("metasploit"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMicrosoftCvrfResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryMicrosoftCVRF `json:"data"`
}

func (c *Client) GetIndexMicrosoftCvrf(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMicrosoftCvrfResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("microsoft-cvrf"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMikrotikResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMikrotik `json:"data"`
}

func (c *Client) GetIndexMikrotik(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMikrotikResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mikrotik"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMindrayResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryMindray `json:"data"`
}

func (c *Client) GetIndexMindray(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMindrayResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mindray"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMispThreatActorsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryMispValue `json:"data"`
}

func (c *Client) GetIndexMispThreatActors(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMispThreatActorsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("misp-threat-actors"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMitelResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryMitel `json:"data"`
}

func (c *Client) GetIndexMitel(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMitelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mitel"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMitreAttackCveResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.ApiMitreAttackToCVE `json:"data"`
}

func (c *Client) GetIndexMitreAttackCve(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMitreAttackCveResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mitre-attack-cve"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMitsubishiElectricResponse struct {
	Benchmark float64                                     `json:"_benchmark"`
	Meta      IndexMeta                                   `json:"_meta"`
	Data      []client.AdvisoryMitsubishiElectricAdvisory `json:"data"`
}

func (c *Client) GetIndexMitsubishiElectric(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMitsubishiElectricResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mitsubishi-electric"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMongodbResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryMongoDB `json:"data"`
}

func (c *Client) GetIndexMongodb(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMongodbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mongodb"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMoxaResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryMoxaAdvisory `json:"data"`
}

func (c *Client) GetIndexMoxa(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMoxaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("moxa"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexMozillaResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryMozillaAdvisory `json:"data"`
}

func (c *Client) GetIndexMozilla(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexMozillaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mozilla"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNaverResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryNaver `json:"data"`
}

func (c *Client) GetIndexNaver(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNaverResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("naver"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNecResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryNEC `json:"data"`
}

func (c *Client) GetIndexNec(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNecResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nec"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNetappResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryNetApp `json:"data"`
}

func (c *Client) GetIndexNetapp(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNetappResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("netapp"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNetgateResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryNetgate `json:"data"`
}

func (c *Client) GetIndexNetgate(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNetgateResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("netgate"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNetgearResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryNetgear `json:"data"`
}

func (c *Client) GetIndexNetgear(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNetgearResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("netgear"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNetskopeResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryNetskope `json:"data"`
}

func (c *Client) GetIndexNetskope(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNetskopeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("netskope"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNginxResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryNginxAdvisory `json:"data"`
}

func (c *Client) GetIndexNginx(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNginxResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nginx"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNhsResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryNHS `json:"data"`
}

func (c *Client) GetIndexNhs(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNhsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nhs"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNiResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryNI `json:"data"`
}

func (c *Client) GetIndexNi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ni"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNistNvdResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.ApiCveItems `json:"data"`
}

func (c *Client) GetIndexNistNvd(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNistNvdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nist-nvd"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNistNvd2CpematchResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.ApiNVD20CPEMatch `json:"data"`
}

func (c *Client) GetIndexNistNvd2Cpematch(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNistNvd2CpematchResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nist-nvd2-cpematch"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNodejsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryNodeJS `json:"data"`
}

func (c *Client) GetIndexNodejs(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNodejsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nodejs"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNokiaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryNokia `json:"data"`
}

func (c *Client) GetIndexNokia(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNokiaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nokia"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNozomiResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryNozomi `json:"data"`
}

func (c *Client) GetIndexNozomi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNozomiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nozomi"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNpmResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexNpm(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNpmResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("npm"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNugetResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexNuget(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNugetResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nuget"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNvidiaResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisorySecurityBulletin `json:"data"`
}

func (c *Client) GetIndexNvidia(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNvidiaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nvidia"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexNzAdvisoriesResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryNZAdvisory `json:"data"`
}

func (c *Client) GetIndexNzAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexNzAdvisoriesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nz-advisories"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOctopusDeployResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryOctopusDeploy `json:"data"`
}

func (c *Client) GetIndexOctopusDeploy(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOctopusDeployResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("octopus-deploy"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOktaResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryOkta `json:"data"`
}

func (c *Client) GetIndexOkta(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOktaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("okta"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOmronResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryOmron `json:"data"`
}

func (c *Client) GetIndexOmron(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOmronResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("omron"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOneEResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryOneE `json:"data"`
}

func (c *Client) GetIndexOneE(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOneEResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("one-e"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOpamResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexOpam(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOpamResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("opam"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOpenCvdbResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryOpenCVDB `json:"data"`
}

func (c *Client) GetIndexOpenCvdb(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOpenCvdbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("open-cvdb"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOpenbsdResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryOpenBSD `json:"data"`
}

func (c *Client) GetIndexOpenbsd(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOpenbsdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openbsd"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOpensshResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryOpenSSH `json:"data"`
}

func (c *Client) GetIndexOpenssh(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOpensshResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openssh"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOpensslSecadvResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryOpenSSLSecAdv `json:"data"`
}

func (c *Client) GetIndexOpensslSecadv(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOpensslSecadvResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openssl-secadv"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOpenstackResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryOpenStack `json:"data"`
}

func (c *Client) GetIndexOpenstack(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOpenstackResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openstack"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOpenwrtResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryWRT `json:"data"`
}

func (c *Client) GetIndexOpenwrt(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOpenwrtResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("openwrt"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOracleResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMetaData `json:"data"`
}

func (c *Client) GetIndexOracle(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOracleResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("oracle"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOracleCpuResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryOracleCPU `json:"data"`
}

func (c *Client) GetIndexOracleCpu(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOracleCpuResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("oracle-cpu"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOracleCpuCsafResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryOracleCPUCSAF `json:"data"`
}

func (c *Client) GetIndexOracleCpuCsaf(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOracleCpuCsafResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("oracle-cpu-csaf"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOsvResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryOSV `json:"data"`
}

func (c *Client) GetIndexOsv(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOsvResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("osv"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOtrsResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryOTRS `json:"data"`
}

func (c *Client) GetIndexOtrs(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOtrsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("otrs"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexOwncloudResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryOwnCloud `json:"data"`
}

func (c *Client) GetIndexOwncloud(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexOwncloudResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("owncloud"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPalantirResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryPalantir `json:"data"`
}

func (c *Client) GetIndexPalantir(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPalantirResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("palantir"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPaloAltoResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryPaloAltoAdvisory `json:"data"`
}

func (c *Client) GetIndexPaloAlto(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPaloAltoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("palo-alto"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPanasonicResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryPanasonic `json:"data"`
}

func (c *Client) GetIndexPanasonic(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPanasonicResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("panasonic"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPapercutResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryPaperCut `json:"data"`
}

func (c *Client) GetIndexPapercut(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPapercutResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("papercut"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPegaResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryPega `json:"data"`
}

func (c *Client) GetIndexPega(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPegaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pega"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPhilipsResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryPhilipsAdvisory `json:"data"`
}

func (c *Client) GetIndexPhilips(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPhilipsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("philips"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPhoenixContactResponse struct {
	Benchmark float64                                 `json:"_benchmark"`
	Meta      IndexMeta                               `json:"_meta"`
	Data      []client.AdvisoryPhoenixContactAdvisory `json:"data"`
}

func (c *Client) GetIndexPhoenixContact(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPhoenixContactResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("phoenix-contact"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPostgressqlResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryPostgresSQL `json:"data"`
}

func (c *Client) GetIndexPostgressql(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPostgressqlResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("postgressql"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexProgressResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryProgress `json:"data"`
}

func (c *Client) GetIndexProgress(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexProgressResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("progress"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexProofpointResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryProofpoint `json:"data"`
}

func (c *Client) GetIndexProofpoint(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexProofpointResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("proofpoint"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPubResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexPub(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPubResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pub"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPureStorageResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryPureStorage `json:"data"`
}

func (c *Client) GetIndexPureStorage(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPureStorageResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pure-storage"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPypaAdvisoriesResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryPyPAAdvisory `json:"data"`
}

func (c *Client) GetIndexPypaAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPypaAdvisoriesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pypa-advisories"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexPypiResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexPypi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexPypiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("pypi"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexQnapResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryQNAPAdvisory `json:"data"`
}

func (c *Client) GetIndexQnap(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexQnapResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("qnap"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexQualcommResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryQualcomm `json:"data"`
}

func (c *Client) GetIndexQualcomm(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexQualcommResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("qualcomm"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexQualysResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryQualys `json:"data"`
}

func (c *Client) GetIndexQualys(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexQualysResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("qualys"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexRansomwareResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryRansomwareExploit `json:"data"`
}

func (c *Client) GetIndexRansomware(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexRansomwareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ransomware"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexRedhatResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryRedhatCVE `json:"data"`
}

func (c *Client) GetIndexRedhat(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexRedhatResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("redhat"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexRenesasResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryRenesas `json:"data"`
}

func (c *Client) GetIndexRenesas(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexRenesasResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("renesas"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexReviveResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryRevive `json:"data"`
}

func (c *Client) GetIndexRevive(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexReviveResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("revive"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexRockwellResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryRockwell `json:"data"`
}

func (c *Client) GetIndexRockwell(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexRockwellResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("rockwell"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexRockyResponse struct {
	Benchmark float64            `json:"_benchmark"`
	Meta      IndexMeta          `json:"_meta"`
	Data      []client.ApiUpdate `json:"data"`
}

func (c *Client) GetIndexRocky(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexRockyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("rocky"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexRuckusResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryRuckus `json:"data"`
}

func (c *Client) GetIndexRuckus(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexRuckusResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ruckus"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexRustsecAdvisoriesResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryRustsecAdvisory `json:"data"`
}

func (c *Client) GetIndexRustsecAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexRustsecAdvisoriesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("rustsec-advisories"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSacertResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySAAdvisory `json:"data"`
}

func (c *Client) GetIndexSacert(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSacertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sacert"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSaintResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisorySaintExploit `json:"data"`
}

func (c *Client) GetIndexSaint(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSaintResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("saint"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSalesforceResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySalesForce `json:"data"`
}

func (c *Client) GetIndexSalesforce(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSalesforceResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("salesforce"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSambaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisorySamba `json:"data"`
}

func (c *Client) GetIndexSamba(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSambaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("samba"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSapResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisorySAP `json:"data"`
}

func (c *Client) GetIndexSap(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSapResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sap"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSchneiderElectricResponse struct {
	Benchmark float64                                    `json:"_benchmark"`
	Meta      IndexMeta                                  `json:"_meta"`
	Data      []client.AdvisorySchneiderElectricAdvisory `json:"data"`
}

func (c *Client) GetIndexSchneiderElectric(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSchneiderElectricResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("schneider-electric"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSecConsultResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySECConsult `json:"data"`
}

func (c *Client) GetIndexSecConsult(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSecConsultResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sec-consult"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSelResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisorySel `json:"data"`
}

func (c *Client) GetIndexSel(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sel"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSentineloneResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisorySentinelOne `json:"data"`
}

func (c *Client) GetIndexSentinelone(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSentineloneResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sentinelone"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexServicenowResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryServiceNow `json:"data"`
}

func (c *Client) GetIndexServicenow(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexServicenowResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("servicenow"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexShadowserverExploitedResponse struct {
	Benchmark float64                                             `json:"_benchmark"`
	Meta      IndexMeta                                           `json:"_meta"`
	Data      []client.AdvisoryShadowServerExploitedVulnerability `json:"data"`
}

func (c *Client) GetIndexShadowserverExploited(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexShadowserverExploitedResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("shadowserver-exploited"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSickResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisorySick `json:"data"`
}

func (c *Client) GetIndexSick(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSickResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sick"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSiemensResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisorySiemensAdvisory `json:"data"`
}

func (c *Client) GetIndexSiemens(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSiemensResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("siemens"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSierraWirelessResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisorySierraWireless `json:"data"`
}

func (c *Client) GetIndexSierraWireless(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSierraWirelessResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sierra-wireless"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSingcertResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisorySingCert `json:"data"`
}

func (c *Client) GetIndexSingcert(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSingcertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("singcert"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSlackwareResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisorySlackware `json:"data"`
}

func (c *Client) GetIndexSlackware(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSlackwareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("slackware"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSolarwindsResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisorySolarWindsAdvisory `json:"data"`
}

func (c *Client) GetIndexSolarwinds(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSolarwindsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("solarwinds"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSolrResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisorySolr `json:"data"`
}

func (c *Client) GetIndexSolr(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSolrResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("solr"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSonicwallResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisorySonicWallAdvisory `json:"data"`
}

func (c *Client) GetIndexSonicwall(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSonicwallResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sonicwall"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSpacelabsHealthcareResponse struct {
	Benchmark float64                                      `json:"_benchmark"`
	Meta      IndexMeta                                    `json:"_meta"`
	Data      []client.AdvisorySpacelabsHealthcareAdvisory `json:"data"`
}

func (c *Client) GetIndexSpacelabsHealthcare(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSpacelabsHealthcareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("spacelabs-healthcare"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSpringResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisorySpring `json:"data"`
}

func (c *Client) GetIndexSpring(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSpringResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("spring"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSsdResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisorySSDAdvisory `json:"data"`
}

func (c *Client) GetIndexSsd(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSsdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ssd"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexStormshieldResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryStormshield `json:"data"`
}

func (c *Client) GetIndexStormshield(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexStormshieldResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("stormshield"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexStrykerResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryStrykerAdvisory `json:"data"`
}

func (c *Client) GetIndexStryker(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexStrykerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("stryker"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSudoResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisorySudo `json:"data"`
}

func (c *Client) GetIndexSudo(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSudoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("sudo"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSuseResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryCvrf `json:"data"`
}

func (c *Client) GetIndexSuse(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSuseResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("suse"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSwiftResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexSwift(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSwiftResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("swift"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSwisslogHealthcareResponse struct {
	Benchmark float64                                     `json:"_benchmark"`
	Meta      IndexMeta                                   `json:"_meta"`
	Data      []client.AdvisorySwisslogHealthcareAdvisory `json:"data"`
}

func (c *Client) GetIndexSwisslogHealthcare(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSwisslogHealthcareResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("swisslog-healthcare"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSymfonyResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisorySymfony `json:"data"`
}

func (c *Client) GetIndexSymfony(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSymfonyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("symfony"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSyncrosoftResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySyncroSoft `json:"data"`
}

func (c *Client) GetIndexSyncrosoft(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSyncrosoftResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("syncrosoft"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexSynologyResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisorySynology `json:"data"`
}

func (c *Client) GetIndexSynology(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexSynologyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("synology"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexTeamviewerResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryTeamViewer `json:"data"`
}

func (c *Client) GetIndexTeamviewer(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexTeamviewerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("teamviewer"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexTencentResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryTencent `json:"data"`
}

func (c *Client) GetIndexTencent(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexTencentResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("tencent"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexThalesResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryThales `json:"data"`
}

func (c *Client) GetIndexThales(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexThalesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("thales"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexThemissinglinkResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryTheMissingLink `json:"data"`
}

func (c *Client) GetIndexThemissinglink(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexThemissinglinkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("themissinglink"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexThreatActorsResponse struct {
	Benchmark float64                                         `json:"_benchmark"`
	Meta      IndexMeta                                       `json:"_meta"`
	Data      []client.AdvisoryThreatActorWithExternalObjects `json:"data"`
}

func (c *Client) GetIndexThreatActors(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexThreatActorsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("threat-actors"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexTiResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryTI `json:"data"`
}

func (c *Client) GetIndexTi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexTiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ti"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexTibcoResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryTibco `json:"data"`
}

func (c *Client) GetIndexTibco(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexTibcoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("tibco"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexTpLinkResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryTPLink `json:"data"`
}

func (c *Client) GetIndexTpLink(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexTpLinkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("tp-link"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexTrendmicroResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryTrendMicro `json:"data"`
}

func (c *Client) GetIndexTrendmicro(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexTrendmicroResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("trendmicro"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexTrustwaveResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryTrustwave `json:"data"`
}

func (c *Client) GetIndexTrustwave(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexTrustwaveResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("trustwave"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexTwcertResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryTWCertAdvisory `json:"data"`
}

func (c *Client) GetIndexTwcert(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexTwcertResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("twcert"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexUbiquitiResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryUbiquiti `json:"data"`
}

func (c *Client) GetIndexUbiquiti(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexUbiquitiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ubiquiti"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexUbuntuResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryUbuntuCVE `json:"data"`
}

func (c *Client) GetIndexUbuntu(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexUbuntuResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ubuntu"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexUnifyResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryUnify `json:"data"`
}

func (c *Client) GetIndexUnify(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexUnifyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("unify"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexUnisocResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryUnisoc `json:"data"`
}

func (c *Client) GetIndexUnisoc(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexUnisocResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("unisoc"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexUsdResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryUSD `json:"data"`
}

func (c *Client) GetIndexUsd(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexUsdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("usd"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexUsomResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryUSOMAdvisory `json:"data"`
}

func (c *Client) GetIndexUsom(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexUsomResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("usom"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVandykeResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryVanDyke `json:"data"`
}

func (c *Client) GetIndexVandyke(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVandykeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vandyke"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVapidlabsResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryVapidLabsAdvisory `json:"data"`
}

func (c *Client) GetIndexVapidlabs(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVapidlabsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vapidlabs"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVdeResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryVDEAdvisory `json:"data"`
}

func (c *Client) GetIndexVde(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVdeResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vde"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVeeamResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryVeeam `json:"data"`
}

func (c *Client) GetIndexVeeam(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVeeamResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("veeam"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVoidsecResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryVoidSec `json:"data"`
}

func (c *Client) GetIndexVoidsec(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVoidsecResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("voidsec"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVulncheckKevResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryVulnCheckKEV `json:"data"`
}

func (c *Client) GetIndexVulncheckKev(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVulncheckKevResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck-kev"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVulncheckNvdResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.ApiCveItemsExtended `json:"data"`
}

func (c *Client) GetIndexVulncheckNvd(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVulncheckNvdResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck-nvd"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVulncheckNvd2Response struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.ApiNVD20CVEExtended `json:"data"`
}

func (c *Client) GetIndexVulncheckNvd2(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVulncheckNvd2Response, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck-nvd2"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVulnerabilityAliasesResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.ApiVulnerabilityAlias `json:"data"`
}

func (c *Client) GetIndexVulnerabilityAliases(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVulnerabilityAliasesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulnerability-aliases"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexVyaireResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryVYAIREAdvisory `json:"data"`
}

func (c *Client) GetIndexVyaire(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexVyaireResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vyaire"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexWatchguardResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryWatchGuard `json:"data"`
}

func (c *Client) GetIndexWatchguard(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexWatchguardResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("watchguard"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexWhatsappResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryWhatsApp `json:"data"`
}

func (c *Client) GetIndexWhatsapp(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexWhatsappResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("whatsapp"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexWibuResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryWibu `json:"data"`
}

func (c *Client) GetIndexWibu(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexWibuResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wibu"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexWiresharkResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryWireshark `json:"data"`
}

func (c *Client) GetIndexWireshark(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexWiresharkResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wireshark"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexWithSecureResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryWithSecure `json:"data"`
}

func (c *Client) GetIndexWithSecure(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexWithSecureResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("with-secure"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexWolfiResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryWolfi `json:"data"`
}

func (c *Client) GetIndexWolfi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexWolfiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wolfi"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexWolfsslResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryWolfSSL `json:"data"`
}

func (c *Client) GetIndexWolfssl(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexWolfsslResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wolfssl"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexWordfenceResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryWordfence `json:"data"`
}

func (c *Client) GetIndexWordfence(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexWordfenceResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("wordfence"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexXenResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryXen `json:"data"`
}

func (c *Client) GetIndexXen(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexXenResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("xen"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexXeroxResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryXerox `json:"data"`
}

func (c *Client) GetIndexXerox(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexXeroxResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("xerox"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexXiaomiResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryXiaomi `json:"data"`
}

func (c *Client) GetIndexXiaomi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexXiaomiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("xiaomi"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexXylemResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryXylem `json:"data"`
}

func (c *Client) GetIndexXylem(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexXylemResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("xylem"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexYokogawaResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryYokogawaAdvisory `json:"data"`
}

func (c *Client) GetIndexYokogawa(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexYokogawaResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("yokogawa"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexYubicoResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryYubico `json:"data"`
}

func (c *Client) GetIndexYubico(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexYubicoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("yubico"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexZdiResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryZeroDayAdvisory `json:"data"`
}

func (c *Client) GetIndexZdi(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexZdiResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zdi"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexZeroDayResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryZeroDayCZ `json:"data"`
}

func (c *Client) GetIndexZeroDay(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexZeroDayResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zero-day"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexZeroscienceResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryZeroScienceAdvisory `json:"data"`
}

func (c *Client) GetIndexZeroscience(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexZeroscienceResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zeroscience"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexZimbraResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryZimbra `json:"data"`
}

func (c *Client) GetIndexZimbra(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexZimbraResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zimbra"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexZoomResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryZoom `json:"data"`
}

func (c *Client) GetIndexZoom(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexZoomResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zoom"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexZscalerResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryZscaler `json:"data"`
}

func (c *Client) GetIndexZscaler(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexZscalerResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zscaler"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexZusoResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryZuso `json:"data"`
}

func (c *Client) GetIndexZuso(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexZusoResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zuso"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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

type GetIndexZyxelResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryZyxel `json:"data"`
}

func (c *Client) GetIndexZyxel(queryParameters ...IndexQueryParameters) (responseJSON *GetIndexZyxelResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("zyxel"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
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
