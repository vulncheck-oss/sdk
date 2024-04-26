package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

var topOfFile = `package sdk

import (
	"encoding/json"
 	"fmt"
	"github.com/vulncheck-oss/sdk/pkg/client"
	"net/http"
	"net/url"
)

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

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "pkg/client/client.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Open the index_funcs.go file in write mode
	file, err := os.OpenFile("index_funcs.go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write the top of the file
	file.WriteString(topOfFile)

	ast.Inspect(node, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		if strings.HasPrefix(typeSpec.Name.Name, "GetIndex") && strings.HasSuffix(typeSpec.Name.Name, "Response") {
			generateTypeAndFunction(file, typeSpec.Name.Name, structType)
		}

		return true
	})
}

func generateTypeAndFunction(file *os.File, name string, structType *ast.StructType) {
	fmt.Println("Generating type and function for:", name)

	// Find the JSON200 field in the struct
	var json200Type string
	for _, field := range structType.Fields.List {
		if field.Names[0].Name == "JSON200" {
			json200Type = fmt.Sprintf("%s", field.Type)
			break
		}
	}

	// Check if JSON200 type starts with FumeResponse and ends with PaginatePagination
	if strings.Contains(json200Type, "FumeResponse") && strings.HasSuffix(json200Type, "PaginatePagination}") {

		startPos := strings.Index(json200Type, "FumeResponse") + len("FumeResponse")
		endPos := strings.Index(json200Type, "PaginatePagination")
		json200Type = json200Type[startPos:endPos]

		if !strings.HasPrefix(json200Type, "Advisory") && !strings.HasPrefix(json200Type, "Api") {
			return
		}

		// Generate the type
		typeStr := fmt.Sprintf("type %s struct {\n", name)
		typeStr += "\tBenchmark float64 `json:\"_benchmark\"`\n"
		typeStr += "\tMeta IndexMeta `json:\"_meta\"`\n"
		typeStr += fmt.Sprintf("\tData []client.%s `json:\"data\"`\n", json200Type)
		typeStr += "}\n\n"
		if _, err := file.WriteString(typeStr); err != nil {
			panic(err)
		}

		indexName := formatIndexName(name)

		// Generate the function
		funcName := strings.TrimSuffix(name, "Response")
		funcStr := fmt.Sprintf("func (c *Client) %s(queryParameters ...IndexQueryParameters) (responseJSON *%s, err error) {\n", funcName, name)
		funcStr += strings.Replace(insideFunction, "::INDEX::", indexName, -1)
		funcStr += "}\n\n"
		if _, err := file.WriteString(funcStr); err != nil {
			panic(err)
		}
	}
}

func formatIndexName(indexName string) string {
	indexName = strings.TrimPrefix(indexName, "GetIndex")
	indexName = strings.TrimSuffix(indexName, "Response")
	var formattedName strings.Builder
	for i, char := range indexName {
		if char >= 'A' && char <= 'Z' {
			if i != 0 {
				formattedName.WriteRune('-')
			}
			formattedName.WriteRune(char + 'a' - 'A')
		} else {
			formattedName.WriteRune(char)
		}
	}
	return formattedName.String()
}
