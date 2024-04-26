test:
	go test -v ./...

format:
	go fmt ./...

client:
	cp ../../api/pkg/client/vulncheck.go pkg/client/client.go
	@go run script/gen_funcs.go

ray:
	@go get github.com/octoper/go-ray
