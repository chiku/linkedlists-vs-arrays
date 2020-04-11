#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

main() {
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out -o coverage.html
  go build -o ./bin/compare ./cmd/compare
  go build -o ./bin/create-report ./cmd/createreport
}

main
