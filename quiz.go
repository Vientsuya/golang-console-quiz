package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"flag"
)

type Question struct {
	content string
	answer  string
}

func main() {

	var filepath string

	flag.StringVar(&filepath, "filepath", "problems.csv", "Filepath to a csv where questions are stored")

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	// closes the file
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading questions")
	}

	questions := make([]Question, len(records))
	answers := make([]string, len(records))

	// populate questions based on csv file
	for i, record := range records {
		questions[i].content = record[0]
		questions[i].answer = record[1]
	}

	// ask the questions
	for i, question := range questions {
		fmt.Print(question.content + ": ")
		fmt.Scan(&answers[i])
	}

	// check how many correct answers
	correctAnswers := 0
	for i, answer := range answers {
		if answer == questions[i].answer {
			correctAnswers++
		}
	}

	// display the score
	fmt.Println("Your score is: " + strconv.Itoa(correctAnswers) + "/" + strconv.Itoa(len(records)))
}
