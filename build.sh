#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

main() {
  go test -cover ./...
  go build -o ./bin/compare ./cmd/compare
  go build -o ./bin/create-report ./cmd/createreport
}

main
