GOPATH := $(shell go env GOPATH)
AIR := $(GOPATH)/bin/air
PORT := 8080

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run: stop-run
	@echo "Running..."
	@go run cmd/api/main.go &

# Stop the running application
stop-run:
	@echo "Stopping application running on port $(PORT)..."
	@lsof -i :$(PORT) -t | xargs kill -9 || echo "No running process found on port $(PORT)."

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Watch the cmd
run-air: stop-air
	@if [ -x "$(AIR)" ]; then \
		$(AIR); \
	else \
		read -p "air is not installed. Do you want to install it now? (y/n) " choice; \
		if [ "$$choice" = "y" ]; then \
			go install github.com/air-verse/air@latest; \
			if [ -x "$(AIR)" ]; then \
				$(AIR); \
			else \
				echo "Error: air binary not found after installation"; \
				exit 1; \
			fi; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi

# Stop the running air process
stop-air:
	@echo "Stopping air running on port $(PORT)..."
	@lsof -i :$(PORT) -t | xargs kill -9 || echo "No running process found on port $(PORT)."

.PHONY: serve stop-run stop-air
serve:
	./tmp/cmd/api/main
