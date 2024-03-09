LOCAL_BIN:=$(CURDIR)/bin

install-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	GOBIN=$(LOCAL_BIN) go install golang.org/x/vuln/cmd/govulncheck@latest

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml
	GOBIN=$(LOCAL_BIN) govulncheck ./...

install-deps:
	go install github.com/gojuno/minimock/v3/cmd/minimock@v3.3.6

mock-generate:
	go generate	./...