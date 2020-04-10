#!/bin/sh

set -euo pipefail
IFS=$'\n\t'

main() {
  ./arraylist/bin/compare-go | tee ./output/result.csv
  ./arraylist/bin/create_report-go
}

main
