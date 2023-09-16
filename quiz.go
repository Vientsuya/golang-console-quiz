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
	userAnswers := make([]string, len(questions))

	// ask the questions
	for i, question := range questions {
		fmt.Print(question.content + ": ")
		fmt.Scan(&userAnswers[i])
	}

	// check how many answers are correct
	correctAnswers := 0
	for i, answer := range userAnswers {
		if answer == questions[i].answer {
			correctAnswers++
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
