package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	inputCsv := flag.String("input", "problems.csv", "Input CSV file name")
	timeLimit := flag.Int("limit", 30, "Time limit")
	flag.Parse()

	file, err := os.Open(*inputCsv)
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

	fmt.Printf("Welcome to the quiz! You have %d seconds to answer the following %d questions:\n", *timeLimit, numQuestions)
	for i, question := range records {
		fmt.Printf("Problem %d: %s = ", i, question[0])
		fmt.Scan(&userAnswer)
		if userAnswer == question[1] {
			correctAnswers++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, numQuestions)

}
