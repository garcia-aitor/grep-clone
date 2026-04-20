# grep-clone

`grep-clone` is a concurrent, grep-style command-line tool written in Go for fast recursive text search across files. It scans directories, searches file contents in parallel using worker goroutines, and prints matches in a structured format.

## Features

- Recursive directory traversal
- Parallel file scanning with worker goroutines
- Simple substring search (`strings.Contains`)
- Grep-style output: `path:line:content`
- Lightweight codebase with minimal dependencies

## Requirements

- Go `1.26.2` or newer

## Installation

Clone the repository and build:

```bash
go build -o mgrep ./mgrep
```

Or run directly without building:

```bash
go run ./mgrep <search-term> <directory>
```

## Usage

```bash
./mgrep <search-term> <directory>
```

Example:

```bash
./mgrep "TODO" .
```

## Output Format

Each match is printed as:

```text
<file-path>:<line-number>:<line-content>
```

Example output:

```text
./worker/worker.go:37:if strings.Contains(scanner.Text(), find) {
```

## How It Works

1. The tool recursively discovers files from the target directory.
2. File paths are pushed into a buffered worklist.
3. Multiple worker goroutines read file paths from the worklist and search each file line by line.
4. Matches are sent to a results channel and printed as they arrive.

## Current Behavior Notes

- Search is case-sensitive.
- Search is substring-based (not regex-based).
- Binary files are not explicitly filtered.
- Permission or read errors are printed and skipped.

## Project Structure

- `mgrep/main.go` - CLI entry point, worker orchestration, and output handling
- `worker/worker.go` - line-by-line file scanning and match collection
- `worklist/worklist.go` - buffered job queue used by workers

## Roadmap Ideas

- Add case-insensitive search option
- Add regex support
- Add include/exclude patterns
- Add context lines around matches
- Add colorized terminal output

