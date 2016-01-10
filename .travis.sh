#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

golang_prereqs() {
  go get github.com/mattn/gom
}

golang_build() {
  cd arraylist && gom test -coverprofile=coverage.out && gom tool cover -html=coverage.out -o coverage.html && cd ..
  cd linkedlist && gom test -coverprofile=coverage.out && gom tool cover -html=coverage.out -o coverage.html && cd ..
  gom install github.com/chiku/linkedlists-vs-arrays/compare
  gom install github.com/chiku/linkedlists-vs-arrays/create_report
}

main() {
  golang_prereqs
  golang_build
}

main
