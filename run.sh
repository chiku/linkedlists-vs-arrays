#!/bin/sh

set -euo pipefail
IFS=$'\n\t'

main() {
  ./bin/compare | tee ./output/result.csv
  ./bin/create-report
}

main
