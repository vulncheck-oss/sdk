<img src="/logo-sdk.png" align="right" alt="VulnCheck Logo" width="150" />

SDK to interact with VulnCheck API.

[![Release](https://img.shields.io/github/v/release/vulncheck-oss/sdk)](https://github.com/vulncheck-oss/sdk/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/vulncheck-oss/sdk)](https://goreportcard.com/report/github.com/vulncheck-oss/sdk)
[![Go Reference](https://pkg.go.dev/badge/github.com/vulncheck-oss/sdk.svg)](https://pkg.go.dev/github.com/vulncheck-oss/sdk)
[![Lint](https://github.com/vulncheck-oss/sdk/actions/workflows/lint.yml/badge.svg)](https://github.com/vulncheck-oss/sdk/actions/workflows/lint.yml)
[![Tests](https://github.com/vulncheck-oss/sdk/actions/workflows/test.yml/badge.svg)](https://github.com/vulncheck-oss/sdk/actions/workflows/test.yml)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/vulncheck-oss/sdk/pulls)

# The VulnCheck SDK

<br />
<br />
<br />

## Usage

### Install the VulnCheck SDK Client

```bash
go get github.com/vulncheck-oss/sdk
```

### Connecting Client

```go
package main

import (
	"fmt"
	"github.com/vulncheck-oss/sdk"
)

func main() {
    client := sdk.Connect("https://api.vulncheck.com", "vulncheck_token")
    fmt.Println(client.GetBaseURL())
}
```


### Available Methods


### PURL
```go
response, err := client.GetPurl("pkg:hex/coherence@0.1.2")

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

### CPE
```go
response, err := client.GetCpe("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

### BACKUP
```go
response, err := client.GetIndexBackup("initial-access")

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

### INDICES
```go
response, err := client.GetIndices()

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

### INDEX
```go
queryParams := sdk.IndexQueryParameters{}

response, err := client.GetIndex("a10", queryParams)

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

### INDEX (looking up a CVE in the vulncheck-nvd2 index)
```go
response, err := client.GetIndexVulncheckNvd2(
    sdk.IndexQueryParameters{
        Cve: "CVE-2019-19781",
    }
)

if err != nil {
    return err
}

description := (*response.Data[0].Descriptions)[0].Value
cvssBaseScore := (*response.Data[0].Metrics.CvssMetricV31)[0].CvssData.BaseScore
```

### Cursor INDEX
```go
queryParams := sdk.IndexQueryParameters{}

response, err := client.GetCursorIndex("vulncheck-nvd2", "cursor_string", queryParams)

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

## Changelog

Please see CHANGELOG for more information on what has changed recently.

## Contributing

Please see CONTRIBUTING for details.

## Security Vulnerabilities

If you discover any security related issues, please use issue tracker.

## Sponsorship

Development of this package is sponsored by VulnCheck learn more about us!

## License

Apache License 2.0. Please see License File for more information.
