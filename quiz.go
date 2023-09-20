package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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

func AskQuestions(questions []Question, timeLimit int) int {
	correctAnswers := 0

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i, question := range questions {
		fmt.Printf("Problem #%d: %s = ", i+1, question.content)
		answerCh := make(chan string)
		go func() {
			var userAnswer string
			fmt.Scan(&userAnswer)
			answerCh <- userAnswer
		}()
		select {
		case <-timer.C:
			DisplayScore(correctAnswers, len(questions))
			os.Exit(0)
		case userAnswer := <- answerCh:
			if userAnswer == question.answer {
				fmt.Println("Correct! \u2713")
				correctAnswers++
			} else {
				fmt.Println("Wrong! \u2715")
			}
		}
	}

	return correctAnswers
}

func DisplayScore(correctAnswers int, questionsCount int) {
	fmt.Println("\n" + "Your score is: " + strconv.Itoa(correctAnswers) + "/" + strconv.Itoa(questionsCount) + "\n")
}

func main() {
	var filepath string
	var timeLimit int
	flag.StringVar(&filepath, "filepath", "problems.csv", "Filepath to a csv where questions are stored")
	flag.IntVar(&timeLimit, "timelimit", 30, "Time limit to answer the questions in seconds")
	flag.Parse()

	records := GetRecords(filepath)
	questions := GetQuestions(records)
	correctAnswers := AskQuestions(questions, timeLimit)
	DisplayScore(correctAnswers, len(questions))
}
