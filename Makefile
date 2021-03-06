.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf $(BIN_DIR)/*

dependencies:
	go mod download

build: dependencies build-api

build-api: 
	go build -tags $(LIBRARY_ENV) -o ./bin/api api/main.go
	docker build -t stonktendency-api .

linux-binaries:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -tags "$(LIBRARY_ENV) netgo" -installsuffix netgo -o $(BIN_DIR)/api api/main.go

ci: dependencies test

build-mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	@~/go/bin/mockgen -source=usecase/book/interface.go -destination=usecase/book/mock/book.go -package=mock

test:
	go test -tags testing ./...
