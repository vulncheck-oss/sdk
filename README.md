# VulnCheck SDK

## Init the VulnCheck SDK Client
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

## PURL
```go
response, err := client.GetPurl("pkg:hex/coherence@0.1.2")

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

## CPE
```go
response, err := client.GetCpe("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

## BACKUP
```go
response, err := client.GetIndexBackup("initial-access")

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

## INDEXES
```go
response, err := client.GetIndexes()

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```

## INDEX
```go
queryParams := sdk.IndexQueryParameters{}

response, err := client.GetIndex("a10", queryParams)

if err != nil {
    panic(err)
}

fmt.Println(response.GetData())
```
