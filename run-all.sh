#!/bin/sh

cd arraylist
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

cd ..

cd linkedlist
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

cd ..

go install github.com/chiku/linkedlists-vs-arrays/compare
go install github.com/chiku/linkedlists-vs-arrays/create_report

compare | tee output/result.csv
create_report
