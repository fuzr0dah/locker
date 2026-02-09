.PHONY: all build clean run

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
	$(BUILD_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)
	@echo "Clean completed"
