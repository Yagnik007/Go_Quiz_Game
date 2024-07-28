package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open the CSV file: %s", *csvFileName)
		os.Exit(1)
	}
	_ = file
}
