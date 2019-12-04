## Go parameters ##
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get

BINARY_NAME=fluorescence
OUTPUT_DIR=./output/

## Actions ##
all: clean run

run: build
	./$(BINARY_NAME)
build:
	$(GOBUILD) -o $(BINARY_NAME) main.go
clean:
	# goworker
	-rm $(BINARY_NAME)
	-rm -r $(OUTPUT_DIR)
