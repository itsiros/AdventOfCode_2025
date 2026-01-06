# Advent of Code 2025

Progress tracking and solutions for **Advent of Code 2025** in Go.

## Overview

[Advent of Code](https://adventofcode.com/2025) is an annual coding challenge with 25 daily puzzles (Dec 1–25). Each day has two parts with increasing difficulty.

This repository contains Go solutions organized by day, with:
- `main.go` — solution code
- `puzzle_input` — full puzzle input
- `puzzle_input_easy` — test/example input
- `puzzle_instructions_part1.txt` & `puzzle_instructions_part2.txt` — problem descriptions

## Progress

| Day | Part 1 | Part 2 | Status |
|-----|--------|--------|--------|
| 1 | ✅ | ✅ | Complete |
| 2 | ✅ | ✅ | Complete |
| 3 | ✅ | ✅ | Complete |
| 4 | ✅ | ✅ | Complete |
| 5 | ✅ | ✅ | Complete |
| 6 | ✅ | ✅ | Complete |
| 7 | ✅ | ✅ | Complete |
| 8 | ✅ | ⏳ | In Progress |
| 9 | ✅ | ⏳| Complete |
| 10 | ⏳ | ⏳ | In Progress |
| 11 | ✅ | ⏳ | In Progress |
| 12 | ⏳ | ⏳ | In Progress |

## Project Structure

```
AdventOfCode_2025/
├── go.mod                   # Go module definition
├── README.md               # this file
└── cmd/
    ├── day1/
    │   ├── main.go         # solution
    │   ├── puzzle_input    # full input (1000+ lines)
    │   ├── puzzle_input_easy  # test input
    │   ├── puzzle_instructions_part1.txt
    │   └── puzzle_instructions_part2.txt
    ├── day2/
    │   └── ... (same structure)
    │
    └── day10/
        └── ... (same structure)
```

## Requirements

- **Go 1.25.4** or later

### Install Go (Linux):
```bash
# via package manager (Arch)
sudo pacman -S go

# or download from golang.org
wget https://go.dev/dl/go1.25.4.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.25.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

## Build & Run

### Run a single day:
```bash
# run day 1
cd cmd/day1 && go run main.go

# or build executable
cd cmd/day1 && go build -o day1 main.go && ./day1
```

### Run with custom input:
```bash
# test with easy input
cd cmd/day1 && go run main.go puzzle_input_easy

# or with full input
cd cmd/day1 && go run main.go puzzle_input
```

### Run all days (loop):
```bash
for day in {1..10}; do
  if [ -f "cmd/day$day/main.go" ]; then
    echo "=== Day $day ==="
    cd "cmd/day$day"
    go run main.go puzzle_input
    cd - > /dev/null
  fi
done
```

### Build all:
```bash
cd cmd
for day in day*/; do
  echo "Building $day..."
  go build -o "$day/$(basename $day)" "$day/main.go"
done
```

## Go Coding Style

All solutions follow:
- Standard library where possible (no external deps)
- Gofmt for formatting (`go fmt ./...`)
- Clear variable names (avoid single-letter except loops)
- Comments for non-obvious logic

## Resources

- **AoC Official:** https://adventofcode.com/2025
- **Go Docs:** https://golang.org/doc/
- **Bufio (input parsing):** https://golang.org/pkg/bufio/

## License

MIT — feel free to reference or adapt for learning.