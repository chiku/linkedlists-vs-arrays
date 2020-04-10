#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

main() {
  go test -cover ./...
  go build -o ./arraylist/bin/compare-go ./compare
  go build -o ./arraylist/bin/create_report-go ./create_report
}

main
