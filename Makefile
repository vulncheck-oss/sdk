test:
	go test -v ./...

format:
	go fmt ./...

client:
	cp ../../api/pkg/client/vulncheck.go pkg/client/client.go

ray:
	@go get github.com/octoper/go-ray
