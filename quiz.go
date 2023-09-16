package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Question struct {
	content string
	answer  string
}

func GetRecords(filepath string) [][]string {

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading questions")
	}

	return records
}

func GetQuestions(records [][]string) []Question {

	questions := make([]Question, len(records))

	for i, record := range records {
		questions[i].content = record[0]
		questions[i].answer = record[1]
	}

	return questions
}

func AskQuestions(questions []Question) int {
	correctAnswers := 0

	for i, question := range questions {
		var userAnswer string
		fmt.Printf("Problem #%d: %s = ", i+1, question.content)
		fmt.Scan(&userAnswer)

		if userAnswer == question.answer {
			fmt.Println("Correct! \u2713")
			correctAnswers++
		} else {
			fmt.Println("Wrong! \u2715")
		}
	}

	return correctAnswers
}

func DisplayScore(correctAnswers int, questionsCount int) {
	fmt.Println("Your score is: " + strconv.Itoa(correctAnswers) + "/" + strconv.Itoa(questionsCount))
}

func main() {
	var filepath string
	flag.StringVar(&filepath, "filepath", "problems.csv", "Filepath to a csv where questions are stored")
	flag.Parse()

	records := GetRecords(filepath)
	questions := GetQuestions(records)
	correctAnswers := AskQuestions(questions)

	DisplayScore(correctAnswers, len(questions))
}
