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

restart: clean clean-output run

run: build
	# execute binary
	./$(BINARY_NAME)
build:
	# go build
	$(GOBUILD) -o $(BINARY_NAME) main.go
clean:
	# remove binary
	-rm $(BINARY_NAME)
clean-output:
	# delete all generated images
	-rm -r $(OUTPUT_DIR)
