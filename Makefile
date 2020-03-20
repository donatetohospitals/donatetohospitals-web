# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=main
BINARY_UNIX=$(BINARY_NAME)_unix
TARGET_DIR=./

build:
	GOOS=linux $(GOBUILD) -o $(TARGET_DIR)/$(BINARY_NAME) -v ./main.go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(TARGET_DIR)/$(BINARY_NAME)
	rm -f $(TARGET_DIR)/$(BINARY_UNIX)

run:
	$(GORUN) ./main.go
