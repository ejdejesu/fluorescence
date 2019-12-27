## Go parameters ##
GOCMD=go
GORUN=$(GOCMD) run
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet

GOLINT=golint

BINARY_NAME=fluorescence
OUTPUT_DIR=./output/

## Actions ##
all: clean install run

restart: flush all

run:
	# execute binary
	$(GOPATH)/bin/$(BINARY_NAME)
install:
	# go install
	$(GOINSTALL) 
build:
	# go build
	$(GOBUILD) -o $(BINARY_NAME) main.go
test:
	$(GOTEST) -v ./...
benchmark:
	$(GOTEST) -v -bench . -run xxx ./...
vet:
	# go vet
	$(GOVET)
lint:
	# golint
	$(GOLINT) ./...
perfect: lint vet
clean:
	# go clean
	$(GOCLEAN)
	# remove binary
	-rm $(GOPATH)/bin/$(BINARY_NAME)
flush:
	# delete all generated images
	-rm -r $(OUTPUT_DIR)