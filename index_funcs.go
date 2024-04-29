package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/vulncheck-oss/sdk/pkg/client"
	"net/http"
	"net/url"
)

// THIS FILE IS GENERATED. DO NOT EDIT

type IndexA10Response struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryA10 `json:"data"`
}

func (c *Client) GetIndexA10(queryParameters ...IndexQueryParameters) (responseJSON *IndexA10Response, err error) {

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

type IndexAbbResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryABBAdvisory `json:"data"`
}

func (c *Client) GetIndexAbb(queryParameters ...IndexQueryParameters) (responseJSON *IndexAbbResponse, err error) {

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

type IndexAbbottResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryAbbott `json:"data"`
}

func (c *Client) GetIndexAbbott(queryParameters ...IndexQueryParameters) (responseJSON *IndexAbbottResponse, err error) {

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

type IndexAbsoluteResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAbsolute `json:"data"`
}

func (c *Client) GetIndexAbsolute(queryParameters ...IndexQueryParameters) (responseJSON *IndexAbsoluteResponse, err error) {

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

type IndexAcronisResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAcronis `json:"data"`
}

func (c *Client) GetIndexAcronis(queryParameters ...IndexQueryParameters) (responseJSON *IndexAcronisResponse, err error) {

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

type IndexAdobeResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAdobeAdvisory `json:"data"`
}

func (c *Client) GetIndexAdobe(queryParameters ...IndexQueryParameters) (responseJSON *IndexAdobeResponse, err error) {

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

type IndexAlephResearchResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAlephResearch `json:"data"`
}

func (c *Client) GetIndexAlephResearch(queryParameters ...IndexQueryParameters) (responseJSON *IndexAlephResearchResponse, err error) {

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

type IndexAlmaResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryAlmaLinuxUpdate `json:"data"`
}

func (c *Client) GetIndexAlma(queryParameters ...IndexQueryParameters) (responseJSON *IndexAlmaResponse, err error) {

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

type IndexAlpineResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryAlpineLinuxSecDB `json:"data"`
}

func (c *Client) GetIndexAlpine(queryParameters ...IndexQueryParameters) (responseJSON *IndexAlpineResponse, err error) {

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

type IndexAmazonResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryUpdate `json:"data"`
}

func (c *Client) GetIndexAmazon(queryParameters ...IndexQueryParameters) (responseJSON *IndexAmazonResponse, err error) {

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

type IndexAmdResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAMD `json:"data"`
}

func (c *Client) GetIndexAmd(queryParameters ...IndexQueryParameters) (responseJSON *IndexAmdResponse, err error) {

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

type IndexAmiResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAMI `json:"data"`
}

func (c *Client) GetIndexAmi(queryParameters ...IndexQueryParameters) (responseJSON *IndexAmiResponse, err error) {

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

type IndexAnchoreNvdOverrideResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryAnchoreNVDOverride `json:"data"`
}

func (c *Client) GetIndexAnchoreNvdOverride(queryParameters ...IndexQueryParameters) (responseJSON *IndexAnchoreNvdOverrideResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("anchore-nvd-override"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAndroidResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryAndroidAdvisory `json:"data"`
}

func (c *Client) GetIndexAndroid(queryParameters ...IndexQueryParameters) (responseJSON *IndexAndroidResponse, err error) {

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

type IndexApacheActivemqResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryApacheActiveMQ `json:"data"`
}

func (c *Client) GetIndexApacheActivemq(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheActivemqResponse, err error) {

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

type IndexApacheArchivaResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheArchiva `json:"data"`
}

func (c *Client) GetIndexApacheArchiva(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheArchivaResponse, err error) {

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

type IndexApacheArrowResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheArrow `json:"data"`
}

func (c *Client) GetIndexApacheArrow(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheArrowResponse, err error) {

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

type IndexApacheCamelResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheCamel `json:"data"`
}

func (c *Client) GetIndexApacheCamel(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheCamelResponse, err error) {

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

type IndexApacheCommonsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheCommons `json:"data"`
}

func (c *Client) GetIndexApacheCommons(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheCommonsResponse, err error) {

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

type IndexApacheCouchdbResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheCouchDB `json:"data"`
}

func (c *Client) GetIndexApacheCouchdb(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheCouchdbResponse, err error) {

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

type IndexApacheFlinkResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheFlink `json:"data"`
}

func (c *Client) GetIndexApacheFlink(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheFlinkResponse, err error) {

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

type IndexApacheGuacamoleResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryApacheGuacamole `json:"data"`
}

func (c *Client) GetIndexApacheGuacamole(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheGuacamoleResponse, err error) {

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

type IndexApacheHadoopResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApacheHadoop `json:"data"`
}

func (c *Client) GetIndexApacheHadoop(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheHadoopResponse, err error) {

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

type IndexApacheHttpResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryApacheHTTP `json:"data"`
}

func (c *Client) GetIndexApacheHttp(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheHttpResponse, err error) {

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

type IndexApacheJspwikiResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryApacheJSPWiki `json:"data"`
}

func (c *Client) GetIndexApacheJspwiki(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheJspwikiResponse, err error) {

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

type IndexApacheKafkaResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheKafka `json:"data"`
}

func (c *Client) GetIndexApacheKafka(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheKafkaResponse, err error) {

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

type IndexApacheLoggingservicesResponse struct {
	Benchmark float64                                `json:"_benchmark"`
	Meta      IndexMeta                              `json:"_meta"`
	Data      []client.AdvisoryApacheLoggingServices `json:"data"`
}

func (c *Client) GetIndexApacheLoggingservices(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheLoggingservicesResponse, err error) {

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

type IndexApacheNifiResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryApacheNiFi `json:"data"`
}

func (c *Client) GetIndexApacheNifi(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheNifiResponse, err error) {

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

type IndexApacheOfbizResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheOFBiz `json:"data"`
}

func (c *Client) GetIndexApacheOfbiz(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheOfbizResponse, err error) {

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

type IndexApacheOpenmeetingsResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryApacheOpenMeetings `json:"data"`
}

func (c *Client) GetIndexApacheOpenmeetings(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheOpenmeetingsResponse, err error) {

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

type IndexApacheOpenofficeResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryApacheOpenOffice `json:"data"`
}

func (c *Client) GetIndexApacheOpenoffice(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheOpenofficeResponse, err error) {

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

type IndexApachePulsarResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApachePulsar `json:"data"`
}

func (c *Client) GetIndexApachePulsar(queryParameters ...IndexQueryParameters) (responseJSON *IndexApachePulsarResponse, err error) {

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

type IndexApacheShiroResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheShiro `json:"data"`
}

func (c *Client) GetIndexApacheShiro(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheShiroResponse, err error) {

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

type IndexApacheSparkResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryApacheSpark `json:"data"`
}

func (c *Client) GetIndexApacheSpark(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheSparkResponse, err error) {

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

type IndexApacheStrutsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApacheStruts `json:"data"`
}

func (c *Client) GetIndexApacheStruts(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheStrutsResponse, err error) {

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

type IndexApacheSubversionResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryApacheSubversion `json:"data"`
}

func (c *Client) GetIndexApacheSubversion(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheSubversionResponse, err error) {

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

type IndexApacheSupersetResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryApacheSuperset `json:"data"`
}

func (c *Client) GetIndexApacheSuperset(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheSupersetResponse, err error) {

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

type IndexApacheTomcatResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryApacheTomcat `json:"data"`
}

func (c *Client) GetIndexApacheTomcat(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheTomcatResponse, err error) {

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

type IndexApacheZookeeperResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryApacheZooKeeper `json:"data"`
}

func (c *Client) GetIndexApacheZookeeper(queryParameters ...IndexQueryParameters) (responseJSON *IndexApacheZookeeperResponse, err error) {

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

type IndexAppcheckResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAppCheck `json:"data"`
}

func (c *Client) GetIndexAppcheck(queryParameters ...IndexQueryParameters) (responseJSON *IndexAppcheckResponse, err error) {

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

type IndexAppgateResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAppgate `json:"data"`
}

func (c *Client) GetIndexAppgate(queryParameters ...IndexQueryParameters) (responseJSON *IndexAppgateResponse, err error) {

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

type IndexAppleResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAppleAdvisory `json:"data"`
}

func (c *Client) GetIndexApple(queryParameters ...IndexQueryParameters) (responseJSON *IndexAppleResponse, err error) {

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

type IndexArchResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryArchIssue `json:"data"`
}

func (c *Client) GetIndexArch(queryParameters ...IndexQueryParameters) (responseJSON *IndexArchResponse, err error) {

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

type IndexAristaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryArista `json:"data"`
}

func (c *Client) GetIndexArista(queryParameters ...IndexQueryParameters) (responseJSON *IndexAristaResponse, err error) {

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

type IndexArubaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryAruba `json:"data"`
}

func (c *Client) GetIndexAruba(queryParameters ...IndexQueryParameters) (responseJSON *IndexArubaResponse, err error) {

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

type IndexAsrgResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryASRG `json:"data"`
}

func (c *Client) GetIndexAsrg(queryParameters ...IndexQueryParameters) (responseJSON *IndexAsrgResponse, err error) {

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

type IndexAssetnoteResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryAssetNote `json:"data"`
}

func (c *Client) GetIndexAssetnote(queryParameters ...IndexQueryParameters) (responseJSON *IndexAssetnoteResponse, err error) {

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

type IndexAsteriskResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAsterisk `json:"data"`
}

func (c *Client) GetIndexAsterisk(queryParameters ...IndexQueryParameters) (responseJSON *IndexAsteriskResponse, err error) {

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

type IndexAsusResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryAsus `json:"data"`
}

func (c *Client) GetIndexAsus(queryParameters ...IndexQueryParameters) (responseJSON *IndexAsusResponse, err error) {

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

type IndexAtlassianResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryAtlassianAdvisory `json:"data"`
}

func (c *Client) GetIndexAtlassian(queryParameters ...IndexQueryParameters) (responseJSON *IndexAtlassianResponse, err error) {

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

type IndexAtlassianVulnsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAtlassianVuln `json:"data"`
}

func (c *Client) GetIndexAtlassianVulns(queryParameters ...IndexQueryParameters) (responseJSON *IndexAtlassianVulnsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("atlassian-vulns"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAtredisResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAtredis `json:"data"`
}

func (c *Client) GetIndexAtredis(queryParameters ...IndexQueryParameters) (responseJSON *IndexAtredisResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("atredis"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexAuscertResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryAusCert `json:"data"`
}

func (c *Client) GetIndexAuscert(queryParameters ...IndexQueryParameters) (responseJSON *IndexAuscertResponse, err error) {

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

type IndexAutodeskResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAutodesk `json:"data"`
}

func (c *Client) GetIndexAutodesk(queryParameters ...IndexQueryParameters) (responseJSON *IndexAutodeskResponse, err error) {

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

type IndexAvayaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryAvaya `json:"data"`
}

func (c *Client) GetIndexAvaya(queryParameters ...IndexQueryParameters) (responseJSON *IndexAvayaResponse, err error) {

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

type IndexAvevaResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryAVEVAAdvisory `json:"data"`
}

func (c *Client) GetIndexAveva(queryParameters ...IndexQueryParameters) (responseJSON *IndexAvevaResponse, err error) {

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

type IndexAvigilonResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAvigilon `json:"data"`
}

func (c *Client) GetIndexAvigilon(queryParameters ...IndexQueryParameters) (responseJSON *IndexAvigilonResponse, err error) {

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

type IndexAwsResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryAWS `json:"data"`
}

func (c *Client) GetIndexAws(queryParameters ...IndexQueryParameters) (responseJSON *IndexAwsResponse, err error) {

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

type IndexAxisResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryAxis `json:"data"`
}

func (c *Client) GetIndexAxis(queryParameters ...IndexQueryParameters) (responseJSON *IndexAxisResponse, err error) {

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

type IndexBandrResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryBandr `json:"data"`
}

func (c *Client) GetIndexBandr(queryParameters ...IndexQueryParameters) (responseJSON *IndexBandrResponse, err error) {

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

type IndexBaxterResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryBaxterAdvisory `json:"data"`
}

func (c *Client) GetIndexBaxter(queryParameters ...IndexQueryParameters) (responseJSON *IndexBaxterResponse, err error) {

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

type IndexBbraunResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryBBraunAdvisory `json:"data"`
}

func (c *Client) GetIndexBbraun(queryParameters ...IndexQueryParameters) (responseJSON *IndexBbraunResponse, err error) {

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

type IndexBdResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryBectonDickinsonAdvisory `json:"data"`
}

func (c *Client) GetIndexBd(queryParameters ...IndexQueryParameters) (responseJSON *IndexBdResponse, err error) {

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

type IndexBduResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryBDUAdvisory `json:"data"`
}

func (c *Client) GetIndexBdu(queryParameters ...IndexQueryParameters) (responseJSON *IndexBduResponse, err error) {

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

type IndexBeckhoffResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryBeckhoffAdvisory `json:"data"`
}

func (c *Client) GetIndexBeckhoff(queryParameters ...IndexQueryParameters) (responseJSON *IndexBeckhoffResponse, err error) {

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

type IndexBeldenResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryBeldenAdvisory `json:"data"`
}

func (c *Client) GetIndexBelden(queryParameters ...IndexQueryParameters) (responseJSON *IndexBeldenResponse, err error) {

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

type IndexBeyondTrustResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryBeyondTrust `json:"data"`
}

func (c *Client) GetIndexBeyondTrust(queryParameters ...IndexQueryParameters) (responseJSON *IndexBeyondTrustResponse, err error) {

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

type IndexBinarlyResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryBinarly `json:"data"`
}

func (c *Client) GetIndexBinarly(queryParameters ...IndexQueryParameters) (responseJSON *IndexBinarlyResponse, err error) {

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

type IndexBitdefenderResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryBitDefender `json:"data"`
}

func (c *Client) GetIndexBitdefender(queryParameters ...IndexQueryParameters) (responseJSON *IndexBitdefenderResponse, err error) {

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

type IndexBlackberryResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryBlackBerry `json:"data"`
}

func (c *Client) GetIndexBlackberry(queryParameters ...IndexQueryParameters) (responseJSON *IndexBlackberryResponse, err error) {

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

type IndexBlsResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryBLS `json:"data"`
}

func (c *Client) GetIndexBls(queryParameters ...IndexQueryParameters) (responseJSON *IndexBlsResponse, err error) {

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

type IndexBoschResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryBoschAdvisory `json:"data"`
}

func (c *Client) GetIndexBosch(queryParameters ...IndexQueryParameters) (responseJSON *IndexBoschResponse, err error) {

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

type IndexBostonScientificResponse struct {
	Benchmark float64                                   `json:"_benchmark"`
	Meta      IndexMeta                                 `json:"_meta"`
	Data      []client.AdvisoryBostonScientificAdvisory `json:"data"`
}

func (c *Client) GetIndexBostonScientific(queryParameters ...IndexQueryParameters) (responseJSON *IndexBostonScientificResponse, err error) {

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

type IndexBotnetsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryBotnet `json:"data"`
}

func (c *Client) GetIndexBotnets(queryParameters ...IndexQueryParameters) (responseJSON *IndexBotnetsResponse, err error) {

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

type IndexCaCyberCentreResponse struct {
	Benchmark float64                                `json:"_benchmark"`
	Meta      IndexMeta                              `json:"_meta"`
	Data      []client.AdvisoryCACyberCentreAdvisory `json:"data"`
}

func (c *Client) GetIndexCaCyberCentre(queryParameters ...IndexQueryParameters) (responseJSON *IndexCaCyberCentreResponse, err error) {

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

type IndexCanvasResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryCanvasExploit `json:"data"`
}

func (c *Client) GetIndexCanvas(queryParameters ...IndexQueryParameters) (responseJSON *IndexCanvasResponse, err error) {

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

type IndexCarestreamResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryCarestreamAdvisory `json:"data"`
}

func (c *Client) GetIndexCarestream(queryParameters ...IndexQueryParameters) (responseJSON *IndexCarestreamResponse, err error) {

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

type IndexCargoResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexCargo(queryParameters ...IndexQueryParameters) (responseJSON *IndexCargoResponse, err error) {

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

type IndexCarrierResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryCarrier `json:"data"`
}

func (c *Client) GetIndexCarrier(queryParameters ...IndexQueryParameters) (responseJSON *IndexCarrierResponse, err error) {

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

type IndexCblMarinerResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryCBLMariner `json:"data"`
}

func (c *Client) GetIndexCblMariner(queryParameters ...IndexQueryParameters) (responseJSON *IndexCblMarinerResponse, err error) {

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

type IndexCentosResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryCESA `json:"data"`
}

func (c *Client) GetIndexCentos(queryParameters ...IndexQueryParameters) (responseJSON *IndexCentosResponse, err error) {

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

type IndexCertBeResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryCertBE `json:"data"`
}

func (c *Client) GetIndexCertBe(queryParameters ...IndexQueryParameters) (responseJSON *IndexCertBeResponse, err error) {

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

type IndexCertUaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryCertUA `json:"data"`
}

func (c *Client) GetIndexCertUa(queryParameters ...IndexQueryParameters) (responseJSON *IndexCertUaResponse, err error) {

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

type IndexCerteuResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCERTEUAdvisory `json:"data"`
}

func (c *Client) GetIndexCerteu(queryParameters ...IndexQueryParameters) (responseJSON *IndexCerteuResponse, err error) {

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

type IndexCertfrResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCertFRAdvisory `json:"data"`
}

func (c *Client) GetIndexCertfr(queryParameters ...IndexQueryParameters) (responseJSON *IndexCertfrResponse, err error) {

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

type IndexChainguardResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryChainGuard `json:"data"`
}

func (c *Client) GetIndexChainguard(queryParameters ...IndexQueryParameters) (responseJSON *IndexChainguardResponse, err error) {

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

type IndexCheckpointResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryCheckPoint `json:"data"`
}

func (c *Client) GetIndexCheckpoint(queryParameters ...IndexQueryParameters) (responseJSON *IndexCheckpointResponse, err error) {

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

type IndexChromeResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryChrome `json:"data"`
}

func (c *Client) GetIndexChrome(queryParameters ...IndexQueryParameters) (responseJSON *IndexChromeResponse, err error) {

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

type IndexCisaAlertsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryCISAAlert `json:"data"`
}

func (c *Client) GetIndexCisaAlerts(queryParameters ...IndexQueryParameters) (responseJSON *IndexCisaAlertsResponse, err error) {

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

type IndexCisaKevResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryKEVCatalogVulnerability `json:"data"`
}

func (c *Client) GetIndexCisaKev(queryParameters ...IndexQueryParameters) (responseJSON *IndexCisaKevResponse, err error) {

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

type IndexCiscoResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryCiscoAdvisory `json:"data"`
}

func (c *Client) GetIndexCisco(queryParameters ...IndexQueryParameters) (responseJSON *IndexCiscoResponse, err error) {

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

type IndexCiscoTalosResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryTalosAdvisory `json:"data"`
}

func (c *Client) GetIndexCiscoTalos(queryParameters ...IndexQueryParameters) (responseJSON *IndexCiscoTalosResponse, err error) {

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

type IndexCitrixResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCitrixAdvisory `json:"data"`
}

func (c *Client) GetIndexCitrix(queryParameters ...IndexQueryParameters) (responseJSON *IndexCitrixResponse, err error) {

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

type IndexClarotyResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryClarotyVulnerability `json:"data"`
}

func (c *Client) GetIndexClaroty(queryParameters ...IndexQueryParameters) (responseJSON *IndexClarotyResponse, err error) {

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

type IndexCloudbeesResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryCloudBees `json:"data"`
}

func (c *Client) GetIndexCloudbees(queryParameters ...IndexQueryParameters) (responseJSON *IndexCloudbeesResponse, err error) {

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

type IndexCloudvulndbResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryCloudVulnDBAdvisory `json:"data"`
}

func (c *Client) GetIndexCloudvulndb(queryParameters ...IndexQueryParameters) (responseJSON *IndexCloudvulndbResponse, err error) {

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

type IndexCnnvdResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryCNNVDEntryJSON `json:"data"`
}

func (c *Client) GetIndexCnnvd(queryParameters ...IndexQueryParameters) (responseJSON *IndexCnnvdResponse, err error) {

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

type IndexCnvdBulletinsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryCNVDBulletin `json:"data"`
}

func (c *Client) GetIndexCnvdBulletins(queryParameters ...IndexQueryParameters) (responseJSON *IndexCnvdBulletinsResponse, err error) {

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

type IndexCnvdFlawsResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryCNVDFlaw `json:"data"`
}

func (c *Client) GetIndexCnvdFlaws(queryParameters ...IndexQueryParameters) (responseJSON *IndexCnvdFlawsResponse, err error) {

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

type IndexCocoapodsResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexCocoapods(queryParameters ...IndexQueryParameters) (responseJSON *IndexCocoapodsResponse, err error) {

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

type IndexCodesysResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryCodesysAdvisory `json:"data"`
}

func (c *Client) GetIndexCodesys(queryParameters ...IndexQueryParameters) (responseJSON *IndexCodesysResponse, err error) {

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

type IndexCompassSecurityResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryCompassSecurity `json:"data"`
}

func (c *Client) GetIndexCompassSecurity(queryParameters ...IndexQueryParameters) (responseJSON *IndexCompassSecurityResponse, err error) {

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

type IndexComposerResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexComposer(queryParameters ...IndexQueryParameters) (responseJSON *IndexComposerResponse, err error) {

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

type IndexConanResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexConan(queryParameters ...IndexQueryParameters) (responseJSON *IndexConanResponse, err error) {

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

type IndexCrestronResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryCrestron `json:"data"`
}

func (c *Client) GetIndexCrestron(queryParameters ...IndexQueryParameters) (responseJSON *IndexCrestronResponse, err error) {

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

type IndexCurlResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryCurl `json:"data"`
}

func (c *Client) GetIndexCurl(queryParameters ...IndexQueryParameters) (responseJSON *IndexCurlResponse, err error) {

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

type IndexCweResponse struct {
	Benchmark float64         `json:"_benchmark"`
	Meta      IndexMeta       `json:"_meta"`
	Data      []client.ApiCWE `json:"data"`
}

func (c *Client) GetIndexCwe(queryParameters ...IndexQueryParameters) (responseJSON *IndexCweResponse, err error) {

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

type IndexDahuaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryDahua `json:"data"`
}

func (c *Client) GetIndexDahua(queryParameters ...IndexQueryParameters) (responseJSON *IndexDahuaResponse, err error) {

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

type IndexDassaultResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryDassault `json:"data"`
}

func (c *Client) GetIndexDassault(queryParameters ...IndexQueryParameters) (responseJSON *IndexDassaultResponse, err error) {

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

type IndexDebianResponse struct {
	Benchmark float64                                  `json:"_benchmark"`
	Meta      IndexMeta                                `json:"_meta"`
	Data      []client.AdvisoryVulnerableDebianPackage `json:"data"`
}

func (c *Client) GetIndexDebian(queryParameters ...IndexQueryParameters) (responseJSON *IndexDebianResponse, err error) {

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

type IndexDebianDsaResponse struct {
	Benchmark float64                                 `json:"_benchmark"`
	Meta      IndexMeta                               `json:"_meta"`
	Data      []client.AdvisoryDebianSecurityAdvisory `json:"data"`
}

func (c *Client) GetIndexDebianDsa(queryParameters ...IndexQueryParameters) (responseJSON *IndexDebianDsaResponse, err error) {

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

type IndexDellResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryDell `json:"data"`
}

func (c *Client) GetIndexDell(queryParameters ...IndexQueryParameters) (responseJSON *IndexDellResponse, err error) {

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

type IndexDeltaResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryDeltaAdvisory `json:"data"`
}

func (c *Client) GetIndexDelta(queryParameters ...IndexQueryParameters) (responseJSON *IndexDeltaResponse, err error) {

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

type IndexDotcmsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryDotCMS `json:"data"`
}

func (c *Client) GetIndexDotcms(queryParameters ...IndexQueryParameters) (responseJSON *IndexDotcmsResponse, err error) {

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

type IndexDragosResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryDragosAdvisory `json:"data"`
}

func (c *Client) GetIndexDragos(queryParameters ...IndexQueryParameters) (responseJSON *IndexDragosResponse, err error) {

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

type IndexDraytekResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryDraytek `json:"data"`
}

func (c *Client) GetIndexDraytek(queryParameters ...IndexQueryParameters) (responseJSON *IndexDraytekResponse, err error) {

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

type IndexEatonResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryEatonAdvisory `json:"data"`
}

func (c *Client) GetIndexEaton(queryParameters ...IndexQueryParameters) (responseJSON *IndexEatonResponse, err error) {

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

type IndexElasticResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryElastic `json:"data"`
}

func (c *Client) GetIndexElastic(queryParameters ...IndexQueryParameters) (responseJSON *IndexElasticResponse, err error) {

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

type IndexElspecResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryElspec `json:"data"`
}

func (c *Client) GetIndexElspec(queryParameters ...IndexQueryParameters) (responseJSON *IndexElspecResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("elspec"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexEmersonResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryEmersonAdvisory `json:"data"`
}

func (c *Client) GetIndexEmerson(queryParameters ...IndexQueryParameters) (responseJSON *IndexEmersonResponse, err error) {

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

type IndexEolResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryEOLReleaseData `json:"data"`
}

func (c *Client) GetIndexEol(queryParameters ...IndexQueryParameters) (responseJSON *IndexEolResponse, err error) {

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

type IndexEpssResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.ApiEPSSData `json:"data"`
}

func (c *Client) GetIndexEpss(queryParameters ...IndexQueryParameters) (responseJSON *IndexEpssResponse, err error) {

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

type IndexExodusIntelResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryExodusIntel `json:"data"`
}

func (c *Client) GetIndexExodusIntel(queryParameters ...IndexQueryParameters) (responseJSON *IndexExodusIntelResponse, err error) {

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

type IndexExploitChainsResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.ApiExploitChain `json:"data"`
}

func (c *Client) GetIndexExploitChains(queryParameters ...IndexQueryParameters) (responseJSON *IndexExploitChainsResponse, err error) {

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

type IndexExploitdbResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryExploitDBExploitv2 `json:"data"`
}

func (c *Client) GetIndexExploitdb(queryParameters ...IndexQueryParameters) (responseJSON *IndexExploitdbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("exploitdb"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexExploitsResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.ApiExploitV3Result `json:"data"`
}

func (c *Client) GetIndexExploits(queryParameters ...IndexQueryParameters) (responseJSON *IndexExploitsResponse, err error) {

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

type IndexFSecureResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryFSecure `json:"data"`
}

func (c *Client) GetIndexFSecure(queryParameters ...IndexQueryParameters) (responseJSON *IndexFSecureResponse, err error) {

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

type IndexFastlyResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryFastly `json:"data"`
}

func (c *Client) GetIndexFastly(queryParameters ...IndexQueryParameters) (responseJSON *IndexFastlyResponse, err error) {

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

type IndexFedoraResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryUpdate `json:"data"`
}

func (c *Client) GetIndexFedora(queryParameters ...IndexQueryParameters) (responseJSON *IndexFedoraResponse, err error) {

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

type IndexFilecloudResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryFileCloud `json:"data"`
}

func (c *Client) GetIndexFilecloud(queryParameters ...IndexQueryParameters) (responseJSON *IndexFilecloudResponse, err error) {

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

type IndexForgerockResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryForgeRock `json:"data"`
}

func (c *Client) GetIndexForgerock(queryParameters ...IndexQueryParameters) (responseJSON *IndexForgerockResponse, err error) {

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

type IndexFortinetResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryFortinetAdvisory `json:"data"`
}

func (c *Client) GetIndexFortinet(queryParameters ...IndexQueryParameters) (responseJSON *IndexFortinetResponse, err error) {

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

type IndexFreebsdResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryAdvisory `json:"data"`
}

func (c *Client) GetIndexFreebsd(queryParameters ...IndexQueryParameters) (responseJSON *IndexFreebsdResponse, err error) {

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

type IndexGallagherResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryGallagher `json:"data"`
}

func (c *Client) GetIndexGallagher(queryParameters ...IndexQueryParameters) (responseJSON *IndexGallagherResponse, err error) {

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

type IndexGcpResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryGCP `json:"data"`
}

func (c *Client) GetIndexGcp(queryParameters ...IndexQueryParameters) (responseJSON *IndexGcpResponse, err error) {

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

type IndexGeGasResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryGEGas `json:"data"`
}

func (c *Client) GetIndexGeGas(queryParameters ...IndexQueryParameters) (responseJSON *IndexGeGasResponse, err error) {

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

type IndexGeHealthcareResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryGEHealthcareAdvisory `json:"data"`
}

func (c *Client) GetIndexGeHealthcare(queryParameters ...IndexQueryParameters) (responseJSON *IndexGeHealthcareResponse, err error) {

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

type IndexGemResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexGem(queryParameters ...IndexQueryParameters) (responseJSON *IndexGemResponse, err error) {

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

type IndexGenetecResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryGenetec `json:"data"`
}

func (c *Client) GetIndexGenetec(queryParameters ...IndexQueryParameters) (responseJSON *IndexGenetecResponse, err error) {

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

type IndexGigabyteResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryGigabyte `json:"data"`
}

func (c *Client) GetIndexGigabyte(queryParameters ...IndexQueryParameters) (responseJSON *IndexGigabyteResponse, err error) {

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

type IndexGiteeExploitsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryGiteeExploit `json:"data"`
}

func (c *Client) GetIndexGiteeExploits(queryParameters ...IndexQueryParameters) (responseJSON *IndexGiteeExploitsResponse, err error) {

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

type IndexGithubExploitsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryGitHubExploit `json:"data"`
}

func (c *Client) GetIndexGithubExploits(queryParameters ...IndexQueryParameters) (responseJSON *IndexGithubExploitsResponse, err error) {

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

type IndexGithubSecurityAdvisoriesResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryGHAdvisoryJSONLean `json:"data"`
}

func (c *Client) GetIndexGithubSecurityAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *IndexGithubSecurityAdvisoriesResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("github-security-advisories"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGitlabAdvisoriesCommunityResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryGitlabAdvisory `json:"data"`
}

func (c *Client) GetIndexGitlabAdvisoriesCommunity(queryParameters ...IndexQueryParameters) (responseJSON *IndexGitlabAdvisoriesCommunityResponse, err error) {

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

type IndexGitlabExploitsResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryGitLabExploit `json:"data"`
}

func (c *Client) GetIndexGitlabExploits(queryParameters ...IndexQueryParameters) (responseJSON *IndexGitlabExploitsResponse, err error) {

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

type IndexGnutlsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryGnuTLS `json:"data"`
}

func (c *Client) GetIndexGnutls(queryParameters ...IndexQueryParameters) (responseJSON *IndexGnutlsResponse, err error) {

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

type IndexGolangResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexGolang(queryParameters ...IndexQueryParameters) (responseJSON *IndexGolangResponse, err error) {

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

type IndexGoogle0dayItwResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryITWExploit `json:"data"`
}

func (c *Client) GetIndexGoogle0dayItw(queryParameters ...IndexQueryParameters) (responseJSON *IndexGoogle0dayItwResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("google-0day-itw"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGoogleContainerOptimizedOsResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryContainerOS `json:"data"`
}

func (c *Client) GetIndexGoogleContainerOptimizedOs(queryParameters ...IndexQueryParameters) (responseJSON *IndexGoogleContainerOptimizedOsResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("google-container-optimized-os"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexGrafanaResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryGrafana `json:"data"`
}

func (c *Client) GetIndexGrafana(queryParameters ...IndexQueryParameters) (responseJSON *IndexGrafanaResponse, err error) {

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

type IndexGreynoiseMetadataResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisoryGreyNoiseDetection `json:"data"`
}

func (c *Client) GetIndexGreynoiseMetadata(queryParameters ...IndexQueryParameters) (responseJSON *IndexGreynoiseMetadataResponse, err error) {

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

type IndexHackageResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexHackage(queryParameters ...IndexQueryParameters) (responseJSON *IndexHackageResponse, err error) {

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

type IndexHarmonyosResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHarmonyOS `json:"data"`
}

func (c *Client) GetIndexHarmonyos(queryParameters ...IndexQueryParameters) (responseJSON *IndexHarmonyosResponse, err error) {

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

type IndexHashicorpResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHashiCorp `json:"data"`
}

func (c *Client) GetIndexHashicorp(queryParameters ...IndexQueryParameters) (responseJSON *IndexHashicorpResponse, err error) {

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

type IndexHaskellSadbResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryHaskellSADBAdvisory `json:"data"`
}

func (c *Client) GetIndexHaskellSadb(queryParameters ...IndexQueryParameters) (responseJSON *IndexHaskellSadbResponse, err error) {

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

type IndexHclResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryHCL `json:"data"`
}

func (c *Client) GetIndexHcl(queryParameters ...IndexQueryParameters) (responseJSON *IndexHclResponse, err error) {

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

type IndexHexResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexHex(queryParameters ...IndexQueryParameters) (responseJSON *IndexHexResponse, err error) {

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

type IndexHikvisionResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHIKVision `json:"data"`
}

func (c *Client) GetIndexHikvision(queryParameters ...IndexQueryParameters) (responseJSON *IndexHikvisionResponse, err error) {

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

type IndexHillromResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryHillromAdvisory `json:"data"`
}

func (c *Client) GetIndexHillrom(queryParameters ...IndexQueryParameters) (responseJSON *IndexHillromResponse, err error) {

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

type IndexHitachiResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryHitachi `json:"data"`
}

func (c *Client) GetIndexHitachi(queryParameters ...IndexQueryParameters) (responseJSON *IndexHitachiResponse, err error) {

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

type IndexHitachiEnergyResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryHitachiEnergy `json:"data"`
}

func (c *Client) GetIndexHitachiEnergy(queryParameters ...IndexQueryParameters) (responseJSON *IndexHitachiEnergyResponse, err error) {

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

type IndexHkcertResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryHKCert `json:"data"`
}

func (c *Client) GetIndexHkcert(queryParameters ...IndexQueryParameters) (responseJSON *IndexHkcertResponse, err error) {

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

type IndexHoneywellResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHoneywell `json:"data"`
}

func (c *Client) GetIndexHoneywell(queryParameters ...IndexQueryParameters) (responseJSON *IndexHoneywellResponse, err error) {

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

type IndexHpResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryHP `json:"data"`
}

func (c *Client) GetIndexHp(queryParameters ...IndexQueryParameters) (responseJSON *IndexHpResponse, err error) {

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

type IndexHuaweiEulerosResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryHuaweiEulerOS `json:"data"`
}

func (c *Client) GetIndexHuaweiEuleros(queryParameters ...IndexQueryParameters) (responseJSON *IndexHuaweiEulerosResponse, err error) {

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

type IndexHuaweiIpsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryHuaweiIPS `json:"data"`
}

func (c *Client) GetIndexHuaweiIps(queryParameters ...IndexQueryParameters) (responseJSON *IndexHuaweiIpsResponse, err error) {

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

type IndexHuaweiPsirtResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryHuawei `json:"data"`
}

func (c *Client) GetIndexHuaweiPsirt(queryParameters ...IndexQueryParameters) (responseJSON *IndexHuaweiPsirtResponse, err error) {

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

type IndexIavaResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryIAVA `json:"data"`
}

func (c *Client) GetIndexIava(queryParameters ...IndexQueryParameters) (responseJSON *IndexIavaResponse, err error) {

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

type IndexIbmResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryIBM `json:"data"`
}

func (c *Client) GetIndexIbm(queryParameters ...IndexQueryParameters) (responseJSON *IndexIbmResponse, err error) {

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

type IndexIdemiaResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryIdemia `json:"data"`
}

func (c *Client) GetIndexIdemia(queryParameters ...IndexQueryParameters) (responseJSON *IndexIdemiaResponse, err error) {

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

type IndexIlAlertsResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryIsraeliAlert `json:"data"`
}

func (c *Client) GetIndexIlAlerts(queryParameters ...IndexQueryParameters) (responseJSON *IndexIlAlertsResponse, err error) {

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

type IndexIlVulnerabilitiesResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryIsraeliVulnerability `json:"data"`
}

func (c *Client) GetIndexIlVulnerabilities(queryParameters ...IndexQueryParameters) (responseJSON *IndexIlVulnerabilitiesResponse, err error) {

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

type IndexIncibeResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryIncibeAdvisory `json:"data"`
}

func (c *Client) GetIndexIncibe(queryParameters ...IndexQueryParameters) (responseJSON *IndexIncibeResponse, err error) {

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

type IndexInitialAccessResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.ApiInitialAccess `json:"data"`
}

func (c *Client) GetIndexInitialAccess(queryParameters ...IndexQueryParameters) (responseJSON *IndexInitialAccessResponse, err error) {

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

type IndexInitialAccessGitResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.ApiInitialAccess `json:"data"`
}

func (c *Client) GetIndexInitialAccessGit(queryParameters ...IndexQueryParameters) (responseJSON *IndexInitialAccessGitResponse, err error) {

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

type IndexIntelResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryIntel `json:"data"`
}

func (c *Client) GetIndexIntel(queryParameters ...IndexQueryParameters) (responseJSON *IndexIntelResponse, err error) {

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

type IndexIpintel10dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

func (c *Client) GetIndexIpintel10d(queryParameters ...IndexQueryParameters) (responseJSON *IndexIpintel10dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel-10d"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIpintel30dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

func (c *Client) GetIndexIpintel30d(queryParameters ...IndexQueryParameters) (responseJSON *IndexIpintel30dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel-30d"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIpintel3dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

func (c *Client) GetIndexIpintel3d(queryParameters ...IndexQueryParameters) (responseJSON *IndexIpintel3dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel-3d"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIpintel90dResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryIpIntelRecord `json:"data"`
}

func (c *Client) GetIndexIpintel90d(queryParameters ...IndexQueryParameters) (responseJSON *IndexIpintel90dResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ipintel-90d"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexIstioResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryIstio `json:"data"`
}

func (c *Client) GetIndexIstio(queryParameters ...IndexQueryParameters) (responseJSON *IndexIstioResponse, err error) {

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

type IndexIvantiResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryIvanti `json:"data"`
}

func (c *Client) GetIndexIvanti(queryParameters ...IndexQueryParameters) (responseJSON *IndexIvantiResponse, err error) {

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

type IndexIvantiRssResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryIvantiRSS `json:"data"`
}

func (c *Client) GetIndexIvantiRss(queryParameters ...IndexQueryParameters) (responseJSON *IndexIvantiRssResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("ivanti-rss"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexJenkinsResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryJenkins `json:"data"`
}

func (c *Client) GetIndexJenkins(queryParameters ...IndexQueryParameters) (responseJSON *IndexJenkinsResponse, err error) {

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

type IndexJetbrainsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryJetBrains `json:"data"`
}

func (c *Client) GetIndexJetbrains(queryParameters ...IndexQueryParameters) (responseJSON *IndexJetbrainsResponse, err error) {

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

type IndexJfrogResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryJFrog `json:"data"`
}

func (c *Client) GetIndexJfrog(queryParameters ...IndexQueryParameters) (responseJSON *IndexJfrogResponse, err error) {

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

type IndexJnjResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryJNJAdvisory `json:"data"`
}

func (c *Client) GetIndexJnj(queryParameters ...IndexQueryParameters) (responseJSON *IndexJnjResponse, err error) {

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

type IndexJvndbResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryJVNAdvisoryItem `json:"data"`
}

func (c *Client) GetIndexJvndb(queryParameters ...IndexQueryParameters) (responseJSON *IndexJvndbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("jvndb"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexKasperskyIcsCertResponse struct {
	Benchmark float64                                   `json:"_benchmark"`
	Meta      IndexMeta                                 `json:"_meta"`
	Data      []client.AdvisoryKasperskyICSCERTAdvisory `json:"data"`
}

func (c *Client) GetIndexKasperskyIcsCert(queryParameters ...IndexQueryParameters) (responseJSON *IndexKasperskyIcsCertResponse, err error) {

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

type IndexKrcertSecurityNoticesResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryKRCertAdvisory `json:"data"`
}

func (c *Client) GetIndexKrcertSecurityNotices(queryParameters ...IndexQueryParameters) (responseJSON *IndexKrcertSecurityNoticesResponse, err error) {

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

type IndexKrcertVulnerabilitiesResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryKRCertAdvisory `json:"data"`
}

func (c *Client) GetIndexKrcertVulnerabilities(queryParameters ...IndexQueryParameters) (responseJSON *IndexKrcertVulnerabilitiesResponse, err error) {

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

type IndexKubernetesResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryK8S `json:"data"`
}

func (c *Client) GetIndexKubernetes(queryParameters ...IndexQueryParameters) (responseJSON *IndexKubernetesResponse, err error) {

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

type IndexLenovoResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryLenovo `json:"data"`
}

func (c *Client) GetIndexLenovo(queryParameters ...IndexQueryParameters) (responseJSON *IndexLenovoResponse, err error) {

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

type IndexLexmarkResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryLexmarkAdvisory `json:"data"`
}

func (c *Client) GetIndexLexmark(queryParameters ...IndexQueryParameters) (responseJSON *IndexLexmarkResponse, err error) {

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

type IndexLgResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryLG `json:"data"`
}

func (c *Client) GetIndexLg(queryParameters ...IndexQueryParameters) (responseJSON *IndexLgResponse, err error) {

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

type IndexLibreOfficeResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryLibreOffice `json:"data"`
}

func (c *Client) GetIndexLibreOffice(queryParameters ...IndexQueryParameters) (responseJSON *IndexLibreOfficeResponse, err error) {

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

type IndexLinuxResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryLinux `json:"data"`
}

func (c *Client) GetIndexLinux(queryParameters ...IndexQueryParameters) (responseJSON *IndexLinuxResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("linux"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMFilesResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMFiles `json:"data"`
}

func (c *Client) GetIndexMFiles(queryParameters ...IndexQueryParameters) (responseJSON *IndexMFilesResponse, err error) {

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

type IndexMacertResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMACert `json:"data"`
}

func (c *Client) GetIndexMacert(queryParameters ...IndexQueryParameters) (responseJSON *IndexMacertResponse, err error) {

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

type IndexManageengineResponse struct {
	Benchmark float64                               `json:"_benchmark"`
	Meta      IndexMeta                             `json:"_meta"`
	Data      []client.AdvisoryManageEngineAdvisory `json:"data"`
}

func (c *Client) GetIndexManageengine(queryParameters ...IndexQueryParameters) (responseJSON *IndexManageengineResponse, err error) {

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

type IndexMavenResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexMaven(queryParameters ...IndexQueryParameters) (responseJSON *IndexMavenResponse, err error) {

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

type IndexMbedTlsResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryMbedTLS `json:"data"`
}

func (c *Client) GetIndexMbedTls(queryParameters ...IndexQueryParameters) (responseJSON *IndexMbedTlsResponse, err error) {

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

type IndexMediatekResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMediatek `json:"data"`
}

func (c *Client) GetIndexMediatek(queryParameters ...IndexQueryParameters) (responseJSON *IndexMediatekResponse, err error) {

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

type IndexMedtronicResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryMedtronicAdvisory `json:"data"`
}

func (c *Client) GetIndexMedtronic(queryParameters ...IndexQueryParameters) (responseJSON *IndexMedtronicResponse, err error) {

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

type IndexMendixResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryMendix `json:"data"`
}

func (c *Client) GetIndexMendix(queryParameters ...IndexQueryParameters) (responseJSON *IndexMendixResponse, err error) {

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

type IndexMetasploitResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryMetasploitExploit `json:"data"`
}

func (c *Client) GetIndexMetasploit(queryParameters ...IndexQueryParameters) (responseJSON *IndexMetasploitResponse, err error) {

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

type IndexMicrosoftCvrfResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryMicrosoftCVRF `json:"data"`
}

func (c *Client) GetIndexMicrosoftCvrf(queryParameters ...IndexQueryParameters) (responseJSON *IndexMicrosoftCvrfResponse, err error) {

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

type IndexMicrosoftKbResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryMicrosoftKb `json:"data"`
}

func (c *Client) GetIndexMicrosoftKb(queryParameters ...IndexQueryParameters) (responseJSON *IndexMicrosoftKbResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("microsoft-kb"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMikrotikResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMikrotik `json:"data"`
}

func (c *Client) GetIndexMikrotik(queryParameters ...IndexQueryParameters) (responseJSON *IndexMikrotikResponse, err error) {

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

type IndexMindrayResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryMindray `json:"data"`
}

func (c *Client) GetIndexMindray(queryParameters ...IndexQueryParameters) (responseJSON *IndexMindrayResponse, err error) {

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

type IndexMispThreatActorsResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryMispValue `json:"data"`
}

func (c *Client) GetIndexMispThreatActors(queryParameters ...IndexQueryParameters) (responseJSON *IndexMispThreatActorsResponse, err error) {

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

type IndexMitelResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryMitel `json:"data"`
}

func (c *Client) GetIndexMitel(queryParameters ...IndexQueryParameters) (responseJSON *IndexMitelResponse, err error) {

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

type IndexMitreAttackCveResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.ApiMitreAttackToCVE `json:"data"`
}

func (c *Client) GetIndexMitreAttackCve(queryParameters ...IndexQueryParameters) (responseJSON *IndexMitreAttackCveResponse, err error) {

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

type IndexMitreCveResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMitreCVE `json:"data"`
}

func (c *Client) GetIndexMitreCve(queryParameters ...IndexQueryParameters) (responseJSON *IndexMitreCveResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("mitre-cve"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexMitsubishiElectricResponse struct {
	Benchmark float64                                     `json:"_benchmark"`
	Meta      IndexMeta                                   `json:"_meta"`
	Data      []client.AdvisoryMitsubishiElectricAdvisory `json:"data"`
}

func (c *Client) GetIndexMitsubishiElectric(queryParameters ...IndexQueryParameters) (responseJSON *IndexMitsubishiElectricResponse, err error) {

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

type IndexMongodbResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryMongoDB `json:"data"`
}

func (c *Client) GetIndexMongodb(queryParameters ...IndexQueryParameters) (responseJSON *IndexMongodbResponse, err error) {

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

type IndexMoxaResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryMoxaAdvisory `json:"data"`
}

func (c *Client) GetIndexMoxa(queryParameters ...IndexQueryParameters) (responseJSON *IndexMoxaResponse, err error) {

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

type IndexMozillaResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryMozillaAdvisory `json:"data"`
}

func (c *Client) GetIndexMozilla(queryParameters ...IndexQueryParameters) (responseJSON *IndexMozillaResponse, err error) {

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

type IndexNaverResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryNaver `json:"data"`
}

func (c *Client) GetIndexNaver(queryParameters ...IndexQueryParameters) (responseJSON *IndexNaverResponse, err error) {

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

type IndexNecResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryNEC `json:"data"`
}

func (c *Client) GetIndexNec(queryParameters ...IndexQueryParameters) (responseJSON *IndexNecResponse, err error) {

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

type IndexNetappResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryNetApp `json:"data"`
}

func (c *Client) GetIndexNetapp(queryParameters ...IndexQueryParameters) (responseJSON *IndexNetappResponse, err error) {

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

type IndexNetgateResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryNetgate `json:"data"`
}

func (c *Client) GetIndexNetgate(queryParameters ...IndexQueryParameters) (responseJSON *IndexNetgateResponse, err error) {

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

type IndexNetgearResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryNetgear `json:"data"`
}

func (c *Client) GetIndexNetgear(queryParameters ...IndexQueryParameters) (responseJSON *IndexNetgearResponse, err error) {

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

type IndexNetskopeResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryNetskope `json:"data"`
}

func (c *Client) GetIndexNetskope(queryParameters ...IndexQueryParameters) (responseJSON *IndexNetskopeResponse, err error) {

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

type IndexNginxResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryNginxAdvisory `json:"data"`
}

func (c *Client) GetIndexNginx(queryParameters ...IndexQueryParameters) (responseJSON *IndexNginxResponse, err error) {

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

type IndexNhsResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryNHS `json:"data"`
}

func (c *Client) GetIndexNhs(queryParameters ...IndexQueryParameters) (responseJSON *IndexNhsResponse, err error) {

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

type IndexNiResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryNI `json:"data"`
}

func (c *Client) GetIndexNi(queryParameters ...IndexQueryParameters) (responseJSON *IndexNiResponse, err error) {

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

type IndexNistNvdResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.ApiCveItems `json:"data"`
}

func (c *Client) GetIndexNistNvd(queryParameters ...IndexQueryParameters) (responseJSON *IndexNistNvdResponse, err error) {

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

type IndexNistNvd2Response struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.ApiNVD20CVE `json:"data"`
}

func (c *Client) GetIndexNistNvd2(queryParameters ...IndexQueryParameters) (responseJSON *IndexNistNvd2Response, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("nist-nvd2"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNistNvd2CpematchResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.ApiNVD20CPEMatch `json:"data"`
}

func (c *Client) GetIndexNistNvd2Cpematch(queryParameters ...IndexQueryParameters) (responseJSON *IndexNistNvd2CpematchResponse, err error) {

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

type IndexNodeSecurityResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryNodeSecurity `json:"data"`
}

func (c *Client) GetIndexNodeSecurity(queryParameters ...IndexQueryParameters) (responseJSON *IndexNodeSecurityResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("node-security"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexNodejsResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryNodeJS `json:"data"`
}

func (c *Client) GetIndexNodejs(queryParameters ...IndexQueryParameters) (responseJSON *IndexNodejsResponse, err error) {

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

type IndexNokiaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryNokia `json:"data"`
}

func (c *Client) GetIndexNokia(queryParameters ...IndexQueryParameters) (responseJSON *IndexNokiaResponse, err error) {

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

type IndexNozomiResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryNozomi `json:"data"`
}

func (c *Client) GetIndexNozomi(queryParameters ...IndexQueryParameters) (responseJSON *IndexNozomiResponse, err error) {

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

type IndexNpmResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexNpm(queryParameters ...IndexQueryParameters) (responseJSON *IndexNpmResponse, err error) {

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

type IndexNugetResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexNuget(queryParameters ...IndexQueryParameters) (responseJSON *IndexNugetResponse, err error) {

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

type IndexNvidiaResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisorySecurityBulletin `json:"data"`
}

func (c *Client) GetIndexNvidia(queryParameters ...IndexQueryParameters) (responseJSON *IndexNvidiaResponse, err error) {

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

type IndexNzAdvisoriesResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryNZAdvisory `json:"data"`
}

func (c *Client) GetIndexNzAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *IndexNzAdvisoriesResponse, err error) {

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

type IndexOctopusDeployResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryOctopusDeploy `json:"data"`
}

func (c *Client) GetIndexOctopusDeploy(queryParameters ...IndexQueryParameters) (responseJSON *IndexOctopusDeployResponse, err error) {

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

type IndexOktaResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryOkta `json:"data"`
}

func (c *Client) GetIndexOkta(queryParameters ...IndexQueryParameters) (responseJSON *IndexOktaResponse, err error) {

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

type IndexOmronResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryOmron `json:"data"`
}

func (c *Client) GetIndexOmron(queryParameters ...IndexQueryParameters) (responseJSON *IndexOmronResponse, err error) {

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

type IndexOneEResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryOneE `json:"data"`
}

func (c *Client) GetIndexOneE(queryParameters ...IndexQueryParameters) (responseJSON *IndexOneEResponse, err error) {

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

type IndexOpamResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexOpam(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpamResponse, err error) {

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

type IndexOpenCvdbResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryOpenCVDB `json:"data"`
}

func (c *Client) GetIndexOpenCvdb(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpenCvdbResponse, err error) {

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

type IndexOpenbsdResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryOpenBSD `json:"data"`
}

func (c *Client) GetIndexOpenbsd(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpenbsdResponse, err error) {

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

type IndexOpensshResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryOpenSSH `json:"data"`
}

func (c *Client) GetIndexOpenssh(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpensshResponse, err error) {

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

type IndexOpensslSecadvResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryOpenSSLSecAdv `json:"data"`
}

func (c *Client) GetIndexOpensslSecadv(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpensslSecadvResponse, err error) {

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

type IndexOpenstackResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryOpenStack `json:"data"`
}

func (c *Client) GetIndexOpenstack(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpenstackResponse, err error) {

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

type IndexOpenwrtResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryWRT `json:"data"`
}

func (c *Client) GetIndexOpenwrt(queryParameters ...IndexQueryParameters) (responseJSON *IndexOpenwrtResponse, err error) {

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

type IndexOracleResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryMetaData `json:"data"`
}

func (c *Client) GetIndexOracle(queryParameters ...IndexQueryParameters) (responseJSON *IndexOracleResponse, err error) {

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

type IndexOracleCpuResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryOracleCPU `json:"data"`
}

func (c *Client) GetIndexOracleCpu(queryParameters ...IndexQueryParameters) (responseJSON *IndexOracleCpuResponse, err error) {

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

type IndexOracleCpuCsafResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.AdvisoryOracleCPUCSAF `json:"data"`
}

func (c *Client) GetIndexOracleCpuCsaf(queryParameters ...IndexQueryParameters) (responseJSON *IndexOracleCpuCsafResponse, err error) {

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

type IndexOsvResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryOSV `json:"data"`
}

func (c *Client) GetIndexOsv(queryParameters ...IndexQueryParameters) (responseJSON *IndexOsvResponse, err error) {

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

type IndexOtrsResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryOTRS `json:"data"`
}

func (c *Client) GetIndexOtrs(queryParameters ...IndexQueryParameters) (responseJSON *IndexOtrsResponse, err error) {

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

type IndexOwncloudResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryOwnCloud `json:"data"`
}

func (c *Client) GetIndexOwncloud(queryParameters ...IndexQueryParameters) (responseJSON *IndexOwncloudResponse, err error) {

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

type IndexPalantirResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryPalantir `json:"data"`
}

func (c *Client) GetIndexPalantir(queryParameters ...IndexQueryParameters) (responseJSON *IndexPalantirResponse, err error) {

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

type IndexPaloAltoResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryPaloAltoAdvisory `json:"data"`
}

func (c *Client) GetIndexPaloAlto(queryParameters ...IndexQueryParameters) (responseJSON *IndexPaloAltoResponse, err error) {

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

type IndexPanasonicResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryPanasonic `json:"data"`
}

func (c *Client) GetIndexPanasonic(queryParameters ...IndexQueryParameters) (responseJSON *IndexPanasonicResponse, err error) {

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

type IndexPapercutResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryPaperCut `json:"data"`
}

func (c *Client) GetIndexPapercut(queryParameters ...IndexQueryParameters) (responseJSON *IndexPapercutResponse, err error) {

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

type IndexPegaResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryPega `json:"data"`
}

func (c *Client) GetIndexPega(queryParameters ...IndexQueryParameters) (responseJSON *IndexPegaResponse, err error) {

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

type IndexPhilipsResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryPhilipsAdvisory `json:"data"`
}

func (c *Client) GetIndexPhilips(queryParameters ...IndexQueryParameters) (responseJSON *IndexPhilipsResponse, err error) {

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

type IndexPhoenixContactResponse struct {
	Benchmark float64                                 `json:"_benchmark"`
	Meta      IndexMeta                               `json:"_meta"`
	Data      []client.AdvisoryPhoenixContactAdvisory `json:"data"`
}

func (c *Client) GetIndexPhoenixContact(queryParameters ...IndexQueryParameters) (responseJSON *IndexPhoenixContactResponse, err error) {

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

type IndexPostgressqlResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryPostgresSQL `json:"data"`
}

func (c *Client) GetIndexPostgressql(queryParameters ...IndexQueryParameters) (responseJSON *IndexPostgressqlResponse, err error) {

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

type IndexProgressResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryProgress `json:"data"`
}

func (c *Client) GetIndexProgress(queryParameters ...IndexQueryParameters) (responseJSON *IndexProgressResponse, err error) {

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

type IndexProofpointResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryProofpoint `json:"data"`
}

func (c *Client) GetIndexProofpoint(queryParameters ...IndexQueryParameters) (responseJSON *IndexProofpointResponse, err error) {

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

type IndexPubResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexPub(queryParameters ...IndexQueryParameters) (responseJSON *IndexPubResponse, err error) {

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

type IndexPureStorageResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryPureStorage `json:"data"`
}

func (c *Client) GetIndexPureStorage(queryParameters ...IndexQueryParameters) (responseJSON *IndexPureStorageResponse, err error) {

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

type IndexPypaAdvisoriesResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryPyPAAdvisory `json:"data"`
}

func (c *Client) GetIndexPypaAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *IndexPypaAdvisoriesResponse, err error) {

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

type IndexPypiResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexPypi(queryParameters ...IndexQueryParameters) (responseJSON *IndexPypiResponse, err error) {

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

type IndexQnapResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryQNAPAdvisory `json:"data"`
}

func (c *Client) GetIndexQnap(queryParameters ...IndexQueryParameters) (responseJSON *IndexQnapResponse, err error) {

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

type IndexQualcommResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryQualcomm `json:"data"`
}

func (c *Client) GetIndexQualcomm(queryParameters ...IndexQueryParameters) (responseJSON *IndexQualcommResponse, err error) {

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

type IndexQualysResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryQualys `json:"data"`
}

func (c *Client) GetIndexQualys(queryParameters ...IndexQueryParameters) (responseJSON *IndexQualysResponse, err error) {

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

type IndexRansomwareResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryRansomwareExploit `json:"data"`
}

func (c *Client) GetIndexRansomware(queryParameters ...IndexQueryParameters) (responseJSON *IndexRansomwareResponse, err error) {

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

type IndexRedhatResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryRedhatCVE `json:"data"`
}

func (c *Client) GetIndexRedhat(queryParameters ...IndexQueryParameters) (responseJSON *IndexRedhatResponse, err error) {

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

type IndexRenesasResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryRenesas `json:"data"`
}

func (c *Client) GetIndexRenesas(queryParameters ...IndexQueryParameters) (responseJSON *IndexRenesasResponse, err error) {

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

type IndexReviveResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryRevive `json:"data"`
}

func (c *Client) GetIndexRevive(queryParameters ...IndexQueryParameters) (responseJSON *IndexReviveResponse, err error) {

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

type IndexRockwellResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryRockwell `json:"data"`
}

func (c *Client) GetIndexRockwell(queryParameters ...IndexQueryParameters) (responseJSON *IndexRockwellResponse, err error) {

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

type IndexRockyResponse struct {
	Benchmark float64            `json:"_benchmark"`
	Meta      IndexMeta          `json:"_meta"`
	Data      []client.ApiUpdate `json:"data"`
}

func (c *Client) GetIndexRocky(queryParameters ...IndexQueryParameters) (responseJSON *IndexRockyResponse, err error) {

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

type IndexRuckusResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryRuckus `json:"data"`
}

func (c *Client) GetIndexRuckus(queryParameters ...IndexQueryParameters) (responseJSON *IndexRuckusResponse, err error) {

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

type IndexRustsecAdvisoriesResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryRustsecAdvisory `json:"data"`
}

func (c *Client) GetIndexRustsecAdvisories(queryParameters ...IndexQueryParameters) (responseJSON *IndexRustsecAdvisoriesResponse, err error) {

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

type IndexSacertResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySAAdvisory `json:"data"`
}

func (c *Client) GetIndexSacert(queryParameters ...IndexQueryParameters) (responseJSON *IndexSacertResponse, err error) {

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

type IndexSaintResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisorySaintExploit `json:"data"`
}

func (c *Client) GetIndexSaint(queryParameters ...IndexQueryParameters) (responseJSON *IndexSaintResponse, err error) {

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

type IndexSalesforceResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySalesForce `json:"data"`
}

func (c *Client) GetIndexSalesforce(queryParameters ...IndexQueryParameters) (responseJSON *IndexSalesforceResponse, err error) {

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

type IndexSambaResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisorySamba `json:"data"`
}

func (c *Client) GetIndexSamba(queryParameters ...IndexQueryParameters) (responseJSON *IndexSambaResponse, err error) {

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

type IndexSapResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisorySAP `json:"data"`
}

func (c *Client) GetIndexSap(queryParameters ...IndexQueryParameters) (responseJSON *IndexSapResponse, err error) {

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

type IndexSchneiderElectricResponse struct {
	Benchmark float64                                    `json:"_benchmark"`
	Meta      IndexMeta                                  `json:"_meta"`
	Data      []client.AdvisorySchneiderElectricAdvisory `json:"data"`
}

func (c *Client) GetIndexSchneiderElectric(queryParameters ...IndexQueryParameters) (responseJSON *IndexSchneiderElectricResponse, err error) {

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

type IndexSecConsultResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySECConsult `json:"data"`
}

func (c *Client) GetIndexSecConsult(queryParameters ...IndexQueryParameters) (responseJSON *IndexSecConsultResponse, err error) {

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

type IndexSelResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisorySel `json:"data"`
}

func (c *Client) GetIndexSel(queryParameters ...IndexQueryParameters) (responseJSON *IndexSelResponse, err error) {

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

type IndexSentineloneResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisorySentinelOne `json:"data"`
}

func (c *Client) GetIndexSentinelone(queryParameters ...IndexQueryParameters) (responseJSON *IndexSentineloneResponse, err error) {

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

type IndexServicenowResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryServiceNow `json:"data"`
}

func (c *Client) GetIndexServicenow(queryParameters ...IndexQueryParameters) (responseJSON *IndexServicenowResponse, err error) {

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

type IndexShadowserverExploitedResponse struct {
	Benchmark float64                                             `json:"_benchmark"`
	Meta      IndexMeta                                           `json:"_meta"`
	Data      []client.AdvisoryShadowServerExploitedVulnerability `json:"data"`
}

func (c *Client) GetIndexShadowserverExploited(queryParameters ...IndexQueryParameters) (responseJSON *IndexShadowserverExploitedResponse, err error) {

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

type IndexSickResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisorySick `json:"data"`
}

func (c *Client) GetIndexSick(queryParameters ...IndexQueryParameters) (responseJSON *IndexSickResponse, err error) {

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

type IndexSiemensResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisorySiemensAdvisory `json:"data"`
}

func (c *Client) GetIndexSiemens(queryParameters ...IndexQueryParameters) (responseJSON *IndexSiemensResponse, err error) {

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

type IndexSierraWirelessResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisorySierraWireless `json:"data"`
}

func (c *Client) GetIndexSierraWireless(queryParameters ...IndexQueryParameters) (responseJSON *IndexSierraWirelessResponse, err error) {

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

type IndexSingcertResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisorySingCert `json:"data"`
}

func (c *Client) GetIndexSingcert(queryParameters ...IndexQueryParameters) (responseJSON *IndexSingcertResponse, err error) {

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

type IndexSlackwareResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisorySlackware `json:"data"`
}

func (c *Client) GetIndexSlackware(queryParameters ...IndexQueryParameters) (responseJSON *IndexSlackwareResponse, err error) {

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

type IndexSolarwindsResponse struct {
	Benchmark float64                             `json:"_benchmark"`
	Meta      IndexMeta                           `json:"_meta"`
	Data      []client.AdvisorySolarWindsAdvisory `json:"data"`
}

func (c *Client) GetIndexSolarwinds(queryParameters ...IndexQueryParameters) (responseJSON *IndexSolarwindsResponse, err error) {

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

type IndexSolrResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisorySolr `json:"data"`
}

func (c *Client) GetIndexSolr(queryParameters ...IndexQueryParameters) (responseJSON *IndexSolrResponse, err error) {

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

type IndexSonicwallResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisorySonicWallAdvisory `json:"data"`
}

func (c *Client) GetIndexSonicwall(queryParameters ...IndexQueryParameters) (responseJSON *IndexSonicwallResponse, err error) {

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

type IndexSpacelabsHealthcareResponse struct {
	Benchmark float64                                      `json:"_benchmark"`
	Meta      IndexMeta                                    `json:"_meta"`
	Data      []client.AdvisorySpacelabsHealthcareAdvisory `json:"data"`
}

func (c *Client) GetIndexSpacelabsHealthcare(queryParameters ...IndexQueryParameters) (responseJSON *IndexSpacelabsHealthcareResponse, err error) {

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

type IndexSpringResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisorySpring `json:"data"`
}

func (c *Client) GetIndexSpring(queryParameters ...IndexQueryParameters) (responseJSON *IndexSpringResponse, err error) {

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

type IndexSsdResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisorySSDAdvisory `json:"data"`
}

func (c *Client) GetIndexSsd(queryParameters ...IndexQueryParameters) (responseJSON *IndexSsdResponse, err error) {

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

type IndexStormshieldResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryStormshield `json:"data"`
}

func (c *Client) GetIndexStormshield(queryParameters ...IndexQueryParameters) (responseJSON *IndexStormshieldResponse, err error) {

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

type IndexStrykerResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryStrykerAdvisory `json:"data"`
}

func (c *Client) GetIndexStryker(queryParameters ...IndexQueryParameters) (responseJSON *IndexStrykerResponse, err error) {

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

type IndexSudoResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisorySudo `json:"data"`
}

func (c *Client) GetIndexSudo(queryParameters ...IndexQueryParameters) (responseJSON *IndexSudoResponse, err error) {

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

type IndexSuseResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryCvrf `json:"data"`
}

func (c *Client) GetIndexSuse(queryParameters ...IndexQueryParameters) (responseJSON *IndexSuseResponse, err error) {

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

type IndexSwiftResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.ApiOSSPackage `json:"data"`
}

func (c *Client) GetIndexSwift(queryParameters ...IndexQueryParameters) (responseJSON *IndexSwiftResponse, err error) {

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

type IndexSwisslogHealthcareResponse struct {
	Benchmark float64                                     `json:"_benchmark"`
	Meta      IndexMeta                                   `json:"_meta"`
	Data      []client.AdvisorySwisslogHealthcareAdvisory `json:"data"`
}

func (c *Client) GetIndexSwisslogHealthcare(queryParameters ...IndexQueryParameters) (responseJSON *IndexSwisslogHealthcareResponse, err error) {

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

type IndexSymfonyResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisorySymfony `json:"data"`
}

func (c *Client) GetIndexSymfony(queryParameters ...IndexQueryParameters) (responseJSON *IndexSymfonyResponse, err error) {

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

type IndexSyncrosoftResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisorySyncroSoft `json:"data"`
}

func (c *Client) GetIndexSyncrosoft(queryParameters ...IndexQueryParameters) (responseJSON *IndexSyncrosoftResponse, err error) {

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

type IndexSynologyResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisorySynology `json:"data"`
}

func (c *Client) GetIndexSynology(queryParameters ...IndexQueryParameters) (responseJSON *IndexSynologyResponse, err error) {

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

type IndexTeamviewerResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryTeamViewer `json:"data"`
}

func (c *Client) GetIndexTeamviewer(queryParameters ...IndexQueryParameters) (responseJSON *IndexTeamviewerResponse, err error) {

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

type IndexTencentResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryTencent `json:"data"`
}

func (c *Client) GetIndexTencent(queryParameters ...IndexQueryParameters) (responseJSON *IndexTencentResponse, err error) {

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

type IndexThalesResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryThales `json:"data"`
}

func (c *Client) GetIndexThales(queryParameters ...IndexQueryParameters) (responseJSON *IndexThalesResponse, err error) {

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

type IndexThemissinglinkResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryTheMissingLink `json:"data"`
}

func (c *Client) GetIndexThemissinglink(queryParameters ...IndexQueryParameters) (responseJSON *IndexThemissinglinkResponse, err error) {

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

type IndexThreatActorsResponse struct {
	Benchmark float64                                         `json:"_benchmark"`
	Meta      IndexMeta                                       `json:"_meta"`
	Data      []client.AdvisoryThreatActorWithExternalObjects `json:"data"`
}

func (c *Client) GetIndexThreatActors(queryParameters ...IndexQueryParameters) (responseJSON *IndexThreatActorsResponse, err error) {

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

type IndexTiResponse struct {
	Benchmark float64             `json:"_benchmark"`
	Meta      IndexMeta           `json:"_meta"`
	Data      []client.AdvisoryTI `json:"data"`
}

func (c *Client) GetIndexTi(queryParameters ...IndexQueryParameters) (responseJSON *IndexTiResponse, err error) {

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

type IndexTibcoResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryTibco `json:"data"`
}

func (c *Client) GetIndexTibco(queryParameters ...IndexQueryParameters) (responseJSON *IndexTibcoResponse, err error) {

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

type IndexTpLinkResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryTPLink `json:"data"`
}

func (c *Client) GetIndexTpLink(queryParameters ...IndexQueryParameters) (responseJSON *IndexTpLinkResponse, err error) {

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

type IndexTraneTechnologyResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryTraneTechnology `json:"data"`
}

func (c *Client) GetIndexTraneTechnology(queryParameters ...IndexQueryParameters) (responseJSON *IndexTraneTechnologyResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("trane-technology"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexTrendmicroResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryTrendMicro `json:"data"`
}

func (c *Client) GetIndexTrendmicro(queryParameters ...IndexQueryParameters) (responseJSON *IndexTrendmicroResponse, err error) {

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

type IndexTrustwaveResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryTrustwave `json:"data"`
}

func (c *Client) GetIndexTrustwave(queryParameters ...IndexQueryParameters) (responseJSON *IndexTrustwaveResponse, err error) {

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

type IndexTwcertResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryTWCertAdvisory `json:"data"`
}

func (c *Client) GetIndexTwcert(queryParameters ...IndexQueryParameters) (responseJSON *IndexTwcertResponse, err error) {

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

type IndexUbiquitiResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryUbiquiti `json:"data"`
}

func (c *Client) GetIndexUbiquiti(queryParameters ...IndexQueryParameters) (responseJSON *IndexUbiquitiResponse, err error) {

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

type IndexUbuntuResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryUbuntuCVE `json:"data"`
}

func (c *Client) GetIndexUbuntu(queryParameters ...IndexQueryParameters) (responseJSON *IndexUbuntuResponse, err error) {

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

type IndexUnifyResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryUnify `json:"data"`
}

func (c *Client) GetIndexUnify(queryParameters ...IndexQueryParameters) (responseJSON *IndexUnifyResponse, err error) {

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

type IndexUnisocResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryUnisoc `json:"data"`
}

func (c *Client) GetIndexUnisoc(queryParameters ...IndexQueryParameters) (responseJSON *IndexUnisocResponse, err error) {

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

type IndexUsdResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryUSD `json:"data"`
}

func (c *Client) GetIndexUsd(queryParameters ...IndexQueryParameters) (responseJSON *IndexUsdResponse, err error) {

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

type IndexUsomResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryUSOMAdvisory `json:"data"`
}

func (c *Client) GetIndexUsom(queryParameters ...IndexQueryParameters) (responseJSON *IndexUsomResponse, err error) {

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

type IndexVandykeResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryVanDyke `json:"data"`
}

func (c *Client) GetIndexVandyke(queryParameters ...IndexQueryParameters) (responseJSON *IndexVandykeResponse, err error) {

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

type IndexVapidlabsResponse struct {
	Benchmark float64                            `json:"_benchmark"`
	Meta      IndexMeta                          `json:"_meta"`
	Data      []client.AdvisoryVapidLabsAdvisory `json:"data"`
}

func (c *Client) GetIndexVapidlabs(queryParameters ...IndexQueryParameters) (responseJSON *IndexVapidlabsResponse, err error) {

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

type IndexVdeResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.AdvisoryVDEAdvisory `json:"data"`
}

func (c *Client) GetIndexVde(queryParameters ...IndexQueryParameters) (responseJSON *IndexVdeResponse, err error) {

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

type IndexVeeamResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryVeeam `json:"data"`
}

func (c *Client) GetIndexVeeam(queryParameters ...IndexQueryParameters) (responseJSON *IndexVeeamResponse, err error) {

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

type IndexVoidsecResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryVoidSec `json:"data"`
}

func (c *Client) GetIndexVoidsec(queryParameters ...IndexQueryParameters) (responseJSON *IndexVoidsecResponse, err error) {

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

type IndexVulncheckConfigResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryVulnCheckConfig `json:"data"`
}

func (c *Client) GetIndexVulncheckConfig(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckConfigResponse, err error) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("vulncheck-config"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
}

type IndexVulncheckKevResponse struct {
	Benchmark float64                       `json:"_benchmark"`
	Meta      IndexMeta                     `json:"_meta"`
	Data      []client.AdvisoryVulnCheckKEV `json:"data"`
}

func (c *Client) GetIndexVulncheckKev(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckKevResponse, err error) {

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

type IndexVulncheckNvdResponse struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.ApiCveItemsExtended `json:"data"`
}

func (c *Client) GetIndexVulncheckNvd(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckNvdResponse, err error) {

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

type IndexVulncheckNvd2Response struct {
	Benchmark float64                      `json:"_benchmark"`
	Meta      IndexMeta                    `json:"_meta"`
	Data      []client.ApiNVD20CVEExtended `json:"data"`
}

func (c *Client) GetIndexVulncheckNvd2(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulncheckNvd2Response, err error) {

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

type IndexVulnerabilityAliasesResponse struct {
	Benchmark float64                        `json:"_benchmark"`
	Meta      IndexMeta                      `json:"_meta"`
	Data      []client.ApiVulnerabilityAlias `json:"data"`
}

func (c *Client) GetIndexVulnerabilityAliases(queryParameters ...IndexQueryParameters) (responseJSON *IndexVulnerabilityAliasesResponse, err error) {

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

type IndexVyaireResponse struct {
	Benchmark float64                         `json:"_benchmark"`
	Meta      IndexMeta                       `json:"_meta"`
	Data      []client.AdvisoryVYAIREAdvisory `json:"data"`
}

func (c *Client) GetIndexVyaire(queryParameters ...IndexQueryParameters) (responseJSON *IndexVyaireResponse, err error) {

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

type IndexWatchguardResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryWatchGuard `json:"data"`
}

func (c *Client) GetIndexWatchguard(queryParameters ...IndexQueryParameters) (responseJSON *IndexWatchguardResponse, err error) {

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

type IndexWhatsappResponse struct {
	Benchmark float64                   `json:"_benchmark"`
	Meta      IndexMeta                 `json:"_meta"`
	Data      []client.AdvisoryWhatsApp `json:"data"`
}

func (c *Client) GetIndexWhatsapp(queryParameters ...IndexQueryParameters) (responseJSON *IndexWhatsappResponse, err error) {

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

type IndexWibuResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryWibu `json:"data"`
}

func (c *Client) GetIndexWibu(queryParameters ...IndexQueryParameters) (responseJSON *IndexWibuResponse, err error) {

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

type IndexWiresharkResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryWireshark `json:"data"`
}

func (c *Client) GetIndexWireshark(queryParameters ...IndexQueryParameters) (responseJSON *IndexWiresharkResponse, err error) {

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

type IndexWithSecureResponse struct {
	Benchmark float64                     `json:"_benchmark"`
	Meta      IndexMeta                   `json:"_meta"`
	Data      []client.AdvisoryWithSecure `json:"data"`
}

func (c *Client) GetIndexWithSecure(queryParameters ...IndexQueryParameters) (responseJSON *IndexWithSecureResponse, err error) {

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

type IndexWolfiResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryWolfi `json:"data"`
}

func (c *Client) GetIndexWolfi(queryParameters ...IndexQueryParameters) (responseJSON *IndexWolfiResponse, err error) {

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

type IndexWolfsslResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryWolfSSL `json:"data"`
}

func (c *Client) GetIndexWolfssl(queryParameters ...IndexQueryParameters) (responseJSON *IndexWolfsslResponse, err error) {

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

type IndexWordfenceResponse struct {
	Benchmark float64                    `json:"_benchmark"`
	Meta      IndexMeta                  `json:"_meta"`
	Data      []client.AdvisoryWordfence `json:"data"`
}

func (c *Client) GetIndexWordfence(queryParameters ...IndexQueryParameters) (responseJSON *IndexWordfenceResponse, err error) {

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

type IndexXenResponse struct {
	Benchmark float64              `json:"_benchmark"`
	Meta      IndexMeta            `json:"_meta"`
	Data      []client.AdvisoryXen `json:"data"`
}

func (c *Client) GetIndexXen(queryParameters ...IndexQueryParameters) (responseJSON *IndexXenResponse, err error) {

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

type IndexXeroxResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryXerox `json:"data"`
}

func (c *Client) GetIndexXerox(queryParameters ...IndexQueryParameters) (responseJSON *IndexXeroxResponse, err error) {

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

type IndexXiaomiResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryXiaomi `json:"data"`
}

func (c *Client) GetIndexXiaomi(queryParameters ...IndexQueryParameters) (responseJSON *IndexXiaomiResponse, err error) {

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

type IndexXylemResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryXylem `json:"data"`
}

func (c *Client) GetIndexXylem(queryParameters ...IndexQueryParameters) (responseJSON *IndexXylemResponse, err error) {

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

type IndexYokogawaResponse struct {
	Benchmark float64                           `json:"_benchmark"`
	Meta      IndexMeta                         `json:"_meta"`
	Data      []client.AdvisoryYokogawaAdvisory `json:"data"`
}

func (c *Client) GetIndexYokogawa(queryParameters ...IndexQueryParameters) (responseJSON *IndexYokogawaResponse, err error) {

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

type IndexYubicoResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryYubico `json:"data"`
}

func (c *Client) GetIndexYubico(queryParameters ...IndexQueryParameters) (responseJSON *IndexYubicoResponse, err error) {

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

type IndexZdiResponse struct {
	Benchmark float64                          `json:"_benchmark"`
	Meta      IndexMeta                        `json:"_meta"`
	Data      []client.AdvisoryZeroDayAdvisory `json:"data"`
}

func (c *Client) GetIndexZdi(queryParameters ...IndexQueryParameters) (responseJSON *IndexZdiResponse, err error) {

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

type IndexZeroscienceResponse struct {
	Benchmark float64                              `json:"_benchmark"`
	Meta      IndexMeta                            `json:"_meta"`
	Data      []client.AdvisoryZeroScienceAdvisory `json:"data"`
}

func (c *Client) GetIndexZeroscience(queryParameters ...IndexQueryParameters) (responseJSON *IndexZeroscienceResponse, err error) {

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

type IndexZimbraResponse struct {
	Benchmark float64                 `json:"_benchmark"`
	Meta      IndexMeta               `json:"_meta"`
	Data      []client.AdvisoryZimbra `json:"data"`
}

func (c *Client) GetIndexZimbra(queryParameters ...IndexQueryParameters) (responseJSON *IndexZimbraResponse, err error) {

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

type IndexZoomResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryZoom `json:"data"`
}

func (c *Client) GetIndexZoom(queryParameters ...IndexQueryParameters) (responseJSON *IndexZoomResponse, err error) {

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

type IndexZscalerResponse struct {
	Benchmark float64                  `json:"_benchmark"`
	Meta      IndexMeta                `json:"_meta"`
	Data      []client.AdvisoryZscaler `json:"data"`
}

func (c *Client) GetIndexZscaler(queryParameters ...IndexQueryParameters) (responseJSON *IndexZscalerResponse, err error) {

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

type IndexZusoResponse struct {
	Benchmark float64               `json:"_benchmark"`
	Meta      IndexMeta             `json:"_meta"`
	Data      []client.AdvisoryZuso `json:"data"`
}

func (c *Client) GetIndexZuso(queryParameters ...IndexQueryParameters) (responseJSON *IndexZusoResponse, err error) {

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

type IndexZyxelResponse struct {
	Benchmark float64                `json:"_benchmark"`
	Meta      IndexMeta              `json:"_meta"`
	Data      []client.AdvisoryZyxel `json:"data"`
}

func (c *Client) GetIndexZyxel(queryParameters ...IndexQueryParameters) (responseJSON *IndexZyxelResponse, err error) {

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
