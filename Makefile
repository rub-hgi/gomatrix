# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOLINT=golint
GOGET=$(GOCMD) get
BINARY_PATH=./bin/
BINARY_NAME=pcs-implementation
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_WINDOWS=$(BINARY_NAME).exe
BINARY_MACOS=$(BINARY_NAME)_macos
TEST_FILES := $(shell $(GOCMD) list ./...)

all: deps test
build: 
		$(GOBUILD) -o $(BINARY_PATH)$(BINARY_NAME) -v
test: 
		$(GOTEST) -v -count=1 -short -covermode=count $(TEST_FILES)
		$(GOLINT) -set_exit_status $(TEST_FILES)
		CC=clang $(GOTEST) -v -count=1 -msan -short $(TEST_FILES)
test-coverage:
		$(GOTEST) -v -count=1 -short -covermode=count -coverprofile report/cover.out $(TEST_FILES)
		$(GOCMD) tool cover -html=report/cover.out -o report/cover.html
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_PATH)$(BINARY_NAME)
		rm -f $(BINARY_PATH)$(BINARY_LINUX)
		rm -f $(BINARY_PATH)$(BINARY_WINDOWS)
		rm -f $(BINARY_PATH)$(BINARY_MACOS)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		GO111MODULE=on $(GOCMD) mod vendor
