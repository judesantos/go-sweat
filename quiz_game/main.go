/*
	Quiz-game

	is a command line application that reads quize items from
	csv file and shows question one at a time, waiting for user to type answer.

	Game will run until all quiz items have been answered or timer expires.


	sourcefile: main.go
*/
package main

import (
	"flag"
	"fmt"
	"os"

	"yourtechy.com/go-sweat/quiz_game/quiz"
)

func main() {

	// Get CLI args - get csv file location in '-csv' arg.

	csvFile := flag.String("csv", quiz.DefaultCsv,
		"The csv file path and filename. Default: deploy.csv")
	timeLimit := flag.Int("timer", quiz.DefaultTimeLimit,
		"The expiration time in seconds. Default: 30s")

	flag.Parse()

	_quiz, err := quiz.NewQuiz(*csvFile, *timeLimit)
	if err != nil {
		exit(err)
	}

	if err := _quiz.Start(); err != nil {
		exit(err)
	}

	score, total := _quiz.GetResult()

	fmt.Printf("\nGame complete! You scored %d out of %d\n\n", score, total)
}

func exit(err error) {
	fmt.Println("Error:", err.Error())
	os.Exit(1)
}
