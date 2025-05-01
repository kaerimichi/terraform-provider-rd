VERSION ?= 1.0.0
BINARY_NAME := terraform-provider-rd_v$(VERSION)
BUILD_DIR := .build

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)

help:
	@echo "Available targets:"
	@echo "  all      - Build the project (default)"
	@echo "  build    - Build the binary"
	@echo "  clean    - Clean build artifacts"
	@echo "  help     - Show this help message"

.PHONY: all build clean help
