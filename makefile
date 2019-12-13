## Go parameters ##
GOCMD=go
GORUN=$(GOCMD) run
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

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
clean:
	# go clean
	$(GOCLEAN)
	# remove binary
	-rm $(GOPATH)/bin/$(BINARY_NAME)
clean-output:
	# delete all generated images
	-rm -r $(OUTPUT_DIR)
