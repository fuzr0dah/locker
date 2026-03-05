.PHONY: all build clean run generate

BINARY_NAME := locker
BUILD_DIR := .build
MAIN_PATH := ./cmd/locker

LDFLAGS := -s -w

all: build

build:
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Built: $(BUILD_DIR)/$(BINARY_NAME)"

run: build
	$(BUILD_DIR)/$(BINARY_NAME) server --dev

clean:
	rm -rf $(BUILD_DIR)
	@echo "Clean completed"

generate:
	rm -rf internal/infrastructure/storage/sqlite/sqlitegen/*.go
	sqlc generate
	@go build ./... || (echo "Generated code is broken!" && exit 1)
	@echo "Generated: sqlc code (verified)"
