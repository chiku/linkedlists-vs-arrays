#!/bin/sh

set -euo pipefail
IFS=$'\n\t'

main() {
  export PATH=$GOROOT/bin:$PATH
  ./output/compare | tee output/result.csv
  ./output/create_report
}

main
