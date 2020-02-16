# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=chat
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_DARWIN=$(BINARY_NAME)_darwin
BINARY_RPI=$(BINARY_NAME)_rpi

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_DARWIN)
	rm -f $(BINARY_RPI)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/labstack/echo
	$(GOGET) golang.org/x/sys/unix
	$(GOGET) github.com/google/uuid
	$(GOGET) golang.org/x/net/websocket

# Cross compilation
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_DARWIN) -v
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
build-rpi:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -o $(BINARY_RPI) -v
