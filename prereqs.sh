#!/bin/bash

set -euo pipefail
IFS=$'\n\t'

main() {
  export PATH=$GOROOT/bin:$PATH
  go get github.com/mattn/gom
}

main
