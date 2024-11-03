# Go Build Info Extractor

A command-line tool to extract and display build information from Go binary files.

## Description

This tool reads and displays detailed build information from compiled Go binaries, including:

- Main module information
- Dependencies
- Build settings
- Go Version
- Module replacements (if any)

## Installation

```bash
go install github.com/maa3x/buildinfo@latest
```

## Usage

```bash
buildinfo <path-to-binary> [additional-binaries...]
```

The tool accepts one or more binary files as arguments and displays their build information in a formatted table.
