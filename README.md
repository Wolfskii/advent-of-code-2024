# Advent of Code 2024 - Go Solutions

My solutions for Advent of Code 2024 using Golang.

## Setup

1. Ensure you have Go 1.23+ installed
2. Clone this repository
3. Navigate to the project directory

## Makefile Commands

### Running Solutions

Run a specific day's solution:
```bash
make run 01      # Runs day01's solution
make run 02      # Runs day02's solution
```

Alternative run method:
```bash
make run DAY=day01
```

### Building

Build all day solutions:
```bash
make build       # Builds all day solutions
```

Clean build artifacts:
```bash
make clean       # Removes build artifacts for all days
```

## Manual Running

Alternatively, you can manually run a solution:
```bash
go run day01/main.go
```

## Directory Structure

- `dayXX/`: Contains solution for each day
  - `main.go`: The solution
  - `input.txt`: Puzzle input
  - `README.md`: Day-specific notes

## Tools

- VSCode recommended
- Go extension installed
- Recommended Go tools:
  - `gofmt` for formatting
  - `golangci-lint` for linting

## Development Workflow

1. Create a new day directory (e.g., `day01`)
2. Implement solution in `main.go`
3. Add input in `input.txt`
4. Run and validate using Makefile commands