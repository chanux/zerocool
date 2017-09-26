GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=zerocool
BINARY_OSX=$(BINARY_NAME)_osx
BINARY_LINUX=$(BINARY_NAME)_linux

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_LINUX) -v
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_OSX) -v

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

build-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME) -a -ldflags '-s' -v
	docker build -t chanux/zerocool .
