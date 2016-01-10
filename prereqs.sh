#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

main() {
  go get github.com/mattn/gom
}

main
