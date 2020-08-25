/*
	Quiz-game

	is a command line application that reads quize items from
	csv file and shows question one at a time, waiting for user to type answer.

	Game will run until all quiz items have been answered or timer expires.


	sourcefile: main.go
*/
package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	defaultCsv       = "./csvs/deploy.csv" // default CSV file name
	maxQuizItems     = 100
	defaultTimeLimit = 30 // 30 seconds
)

type problem struct {
	q string
	a string
}

var timeLimit *int

func main() {

	//
	// Get CLI args - get csv file location in '-csv' arg.

	csvFile := flag.String("csv", defaultCsv,
		"The csv file path and filename. Default: deploy.csv")
	timeLimit = flag.Int("timer", defaultTimeLimit,
		"The expiration time in seconds. Default: 30s")

	flag.Parse()

	//
	// parse file and return all quiz items

	lines, err := getQuizItems(csvFile)
	if err != nil {
		exit(err)
	}

	//
	// start game

	fmt.Printf("\nWelcome to Quiz Game!\nStarting game...\nEnjoy!!!\n\n")

	runGame(lines)

}

//
// getQuizItems - Get quiz items from CSV file.
// Validate if quiz item count is within game rule limit.
//
// @return quiz items
//
func getQuizItems(csvFileName *string) (*[]problem, error) {

	//
	// Check if file exists

	_, err := ioutil.ReadFile(*csvFileName)
	if err != nil {
		exit(err)
	}

	f, err := os.Open(*csvFileName)
	if err != nil {
		exit(err)
	}

	//
	// Load and extract lines from CSV file

	r := csv.NewReader(f)

	// read and store all lines
	lines, err := r.ReadAll()

	// check if read ok
	if err != nil {
		return nil, err
	}

	// Validate quiz items limit

	if len(lines) > maxQuizItems {
		return nil, errors.New("Quiz item exceeds game limit.")
	}

	// Pre process each record into question and answer items

	problems := parseLines(&lines)

	return problems, nil
}

func parseLines(lines *[][]string) *[]problem {

	// Create container with specified length
	ret := make([]problem, len(*lines))

	for i, line := range *lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return &ret
}

//
// runGame - Run the game, continue until all quiz items are posed and answered,
// or until time runs out.
func runGame(problems *[]problem) {

	var correct int = 0
	var timer = time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// Loop through quiz items

	for i, p := range *problems {

		// Pose question

		fmt.Printf("Problem %d: %s = ", i+1, p.q)

		// Pause to read user answer

		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		// Wait for timeout, or user answer

		select {
		case <-timer.C:
			fmt.Printf("\n\nYou're time is up. Game over!\n")
			return
		case answer := <-answerChan:
			if answer == p.a {
				correct++
			}
		}

	} // end for

	fmt.Printf("\nGame complete! You scored %d out of %d.\n\n", correct, len(*problems))
}

func exit(err error) {
	fmt.Println("Error:", err.Error())
	os.Exit(1)
}
