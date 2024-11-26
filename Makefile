.PHONY: all clean test run

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run

# Directories
DAYS := $(wildcard day*/)

# Build all day solutions
build:
	@for dir in $(DAYS); do \
		echo "Building $$dir"; \
		(cd $$dir && $(GOBUILD) ./...); \
	done

# Clean build artifacts
clean:
	@for dir in $(DAYS); do \
		echo "Cleaning $$dir"; \
		(cd $$dir && rm -f *.out); \
	done

# Run a specific day's solution
run:
	@if [ -z "$(filter 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25, $(filter-out run, $(MAKECMDGOALS)))" ]; then \
		echo "Please specify a day (01-25), e.g., make run 01"; \
		exit 1; \
	fi
	@(cd day$(filter 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25, $(filter-out run, $(MAKECMDGOALS))) && $(GORUN) main.go)

# Dummy target to allow make run 01 syntax
01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25:
	@: