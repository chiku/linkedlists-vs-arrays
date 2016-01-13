#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

main() {
  export PATH=$GOROOT/bin:$PATH
  cd arraylist && gom test -coverprofile=coverage.out && gom tool cover -html=coverage.out -o coverage.html && cd ..
  cd linkedlist && gom test -coverprofile=coverage.out && gom tool cover -html=coverage.out -o coverage.html && cd ..
  gom build -o output/compare github.com/chiku/linkedlists-vs-arrays/compare
  gom build -o output/create_report github.com/chiku/linkedlists-vs-arrays/create_report
}

main
