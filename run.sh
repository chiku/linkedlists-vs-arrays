#!/bin/sh

set -euo pipefail
IFS=$'\n\t'

main() {
  ./output/compare | tee output/result.csv
  ./output/create_report
}

main
