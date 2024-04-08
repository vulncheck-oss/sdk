# VulnCheck SDK

SDK to interact with VulnCheck API.

## Usage

### Init the VulnCheck SDK Client

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

### INDEXES
```go
response, err := client.GetIndexes()

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
