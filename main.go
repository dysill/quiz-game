package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
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
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	answerChan := make(chan string)
	var correctAnswers int

	fmt.Printf("Welcome to the quiz! You have %d seconds to answer the following %d questions:\n", *timeLimit, numQuestions)
loop:
	for i, question := range records {
		fmt.Printf("Problem %d: %s = ", i, question[0])
		go func() {
			var userAnswer string
			fmt.Scan(&userAnswer)
			answerChan <- userAnswer
		}()
		select {
		case answer := <-answerChan:
			if answer == question[1] {
				correctAnswers++
			}
		case <-timer.C:
			fmt.Println("\nYou ran out of time!")
			break loop
		}

	}
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, numQuestions)

}
