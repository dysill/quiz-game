package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// handle flags for time and file (later)
	inputCsv := "problems.csv"

	// Read the CSV
	file, err := os.Open(inputCsv)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(records)
	fmt.Println("Welcome to the quiz! You have _ seconds to answer the following _ questions:")

}
