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

	file, err := os.Open(inputCsv)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	numQuestions := len(records)
	var userAnswer string
	var correctAnswers int

	fmt.Printf("Welcome to the quiz! You have _ seconds to answer the following %d questions:\n", numQuestions)
	for i, question := range records {
		fmt.Printf("Problem %d: %s = ", i, question[0])
		fmt.Scan(&userAnswer)
		if userAnswer == question[1] {
			correctAnswers++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, numQuestions)

}
