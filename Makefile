REPOSITORY := github.com/mytheta/go-layerd-architecture-sample
GOLIST := $(shell go list ./...)
PACKAGES := $(shell go list ./... | grep -v -e tools)

.PHONY: build
build:
	GO111MODULE=on CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' interface/api/server/main.go

.PHONY: run
run:
	go run interface/api/server/main.go -e ./interface/env/conf


.PHONY: tools
tools:
	go build -o bin/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint
	go build -o bin/errcheck github.com/kisielk/errcheck
	go build -o bin/staticcheck honnef.co/go/tools/cmd/staticcheck

.PHONY: fmt
fmt:
	! gofmt -s -d $(subst $(REPOSITORY),.,$(PACKAGES)) | grep -E '^'

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: lint
lint:
	./bin/golangci-lint run --config .golangci.yml ./...

.PHONY: test
test:
	go test -cover $(GOLIST)

.PHONY: errcheck
errcheck:
	./bin/errcheck $(PACKAGES)

.PHONY: staticcheck
staticcheck:
	./bin/staticcheck $(PACKAGES)

.PHONY: mod
mod:
	go mod download