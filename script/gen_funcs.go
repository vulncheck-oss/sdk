package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"unicode"
)

var topOfFile = `package sdk

import (
	"encoding/json"
 	"fmt"
	"github.com/vulncheck-oss/sdk/pkg/client"
	"net/http"
	"net/url"
)

// THIS FILE IS GENERATED. DO NOT EDIT

`
var insideFunction = `
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", c.GetUrl()+"/v3/index/"+url.QueryEscape("::INDEX::"), nil)
	if err != nil {
		panic(err)
	}

	c.SetAuthHeader(req)

	query := req.URL.Query()
	setIndexQueryParameters(query, queryParameters...)
	req.URL.RawQuery = query.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var metaError MetaError
		_ = json.NewDecoder(resp.Body).Decode(&metaError)

		return nil, fmt.Errorf("error: %v", metaError.Errors)
	}

	_ = json.NewDecoder(resp.Body).Decode(&responseJSON)

	return responseJSON, nil
`

type docsIndex struct {
	Name   string `yaml:"name"`
	Title  string `yaml:"title"`
	Desc   string `yaml:"desc"`
	Struct string `yaml:"struct"`
	Large  bool   `yaml:"large"`
	Silent bool   `yaml:"silent"`
}

type docsIndices struct {
	Indices []docsIndex `yaml:"indices"`
}

// This points to the indices file in th docs repo
var indicesFilename = "../../docs/packages/indices/indices.yaml"

var funcsFilename = "index_funcs.go"

func main() {

	indicesFile, err := os.Open(indicesFilename)

	if err != nil {
		panic(err)
		return
	}

	var indices docsIndices
	err = yaml.NewDecoder(indicesFile).Decode(&indices)
	if err != nil {
		panic(err)
		return
	}

	funcsFile, err := os.OpenFile(funcsFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
		return
	}
	defer funcsFile.Close()

	// Write the top of the file
	funcsFile.WriteString(topOfFile)

	for _, index := range indices.Indices {

		if index.Silent || index.Name == "advisories" {
			continue
		}

		// Generate the type
		typeStr := fmt.Sprintf("type Index%sResponse struct {\n", toCamelCase(index.Name))
		typeStr += "\tBenchmark float64 `json:\"_benchmark\"`\n"
		typeStr += "\tMeta IndexMeta `json:\"_meta\"`\n"
		typeStr += fmt.Sprintf("\tData []client.%s `json:\"data\"`\n", strings.TrimSuffix(toCamelCase(index.Struct), "{}"))
		typeStr += "}\n\n"

		if _, err := funcsFile.WriteString(typeStr); err != nil {
			panic(err)
		}

		// Generate the function
		funcName := fmt.Sprintf("GetIndex%s", toCamelCase(index.Name))
		funcStr := fmt.Sprintf("// GetIndex%s -  %s\n", toCamelCase(index.Name), index.Desc)
		funcStr += fmt.Sprintf("func (c *Client) %s(queryParameters ...IndexQueryParameters) (responseJSON *Index%sResponse, err error) {\n", funcName, toCamelCase(index.Name))
		funcStr += strings.Replace(insideFunction, "::INDEX::", index.Name, -1)
		funcStr += "}\n\n"
		if _, err := funcsFile.WriteString(funcStr); err != nil {
			panic(err)
		}
	}

}

func toCamelCase(s string) string {
	// Split the string by dash
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '-' || r == '.'
	})

	// Convert each word to title case (first letter uppercase)
	for i, word := range words {
		if unicode.IsLower(rune(word[0])) {
			words[i] = strings.ToTitle(word[:1]) + word[1:]
		} else {
			words[i] = word
		}
	}

	// Join the words together
	result := strings.Join(words, "")

	return result
}
