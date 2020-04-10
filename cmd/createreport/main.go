package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"text/template"
)

type result struct {
	ContainerSizes      string
	ArrayDurations      string
	LinkedListDurations string
}

func main() {
	file, err := os.Open("output/result.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV file: %v\n", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	var res result

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading CSV file: %v\n", err)
			return
		}
		res.ContainerSizes += (record[0] + ",")
		res.ArrayDurations += (record[1] + ",")
		res.LinkedListDurations += (record[2] + ",")
	}

	template, err := template.New("comparison.js.tmpl").ParseFiles("public/comparison.js.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening template file: %v\n", err)
		return
	}

	output, err := os.Create("public/comparison.js")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing result file: %v\n", err)
		return
	}
	defer output.Close()

	err = template.Execute(output, res)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error populating template: %v\n", err)
	}
}
