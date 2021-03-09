.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev
CGO_ENABLED=0

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf $(BIN_DIR)/*

dependencies:
	go mod download

build: dependencies build-api

build-api: 
	docker build -t stonktendency-api .

ci: dependencies test

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=usecase/book/interface.go -destination=usecase/book/mock/book.go -package=mock

test:
	go test -tags testing ./...
