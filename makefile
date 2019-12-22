## Go parameters ##
GOCMD=go
GORUN=$(GOCMD) run
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet

BINARY_NAME=fluorescence
OUTPUT_DIR=./output/

## Actions ##
all: clean install run

restart: clean clean-output install run

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
clean:
	# go clean
	$(GOCLEAN)
	# remove binary
	-rm $(GOPATH)/bin/$(BINARY_NAME)
flush:
	# delete all generated images
	-rm -r $(OUTPUT_DIR)