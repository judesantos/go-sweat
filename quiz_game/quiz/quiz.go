package quiz

// quiz.go

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"yourtechy.com/go-sweat/utils/logger"
)

var (
	log = logger.NewLogger()
)

const (
	DefaultCsv       = "../csvs/deploy.csv" // default CSV file name
	MaxQuizItems     = 100
	DefaultTimeLimit = 30 // 30 seconds
)

type problem struct {
	q string
	a string
}

type quiz struct {
	csvFile   string
	timeLimit int
	score     int
	total     int
}

/**************************************************************************
 *
 *	IQuiz interface methods implementation
 *
 *************************************************************************/

type IQuiz interface {
	Start() error
	GetResult() (int, int)
}

// NewQuiz - Get new quiz instance
func NewQuiz(csvFile string, limit int) (*quiz, error) {

	// check if not blank

	if csvFile == "" {
		msg := "No csv file provided"
		//log.Trace(msg)
		return nil, errors.New(msg)
	}

	// Check if file exists

	_, err := ioutil.ReadFile(csvFile)
	if err != nil {
		//log.Trace(err.Error())
		return nil, err
	}

	return &quiz{
		csvFile:   csvFile,
		timeLimit: limit,
		score:     0,
		total:     0,
	}, nil
}

// Start - IQuiz interface method implementation
// Start game engine.
func (q *quiz) Start() error {

	// parse file and return all quiz items

	lines, err := getQuizItems(&q.csvFile)
	if err != nil {
		return err
	}

	// get number of questions
	q.total = len(*lines)

	// block until all quiz items are read, or time out expires
	return runGame(q, lines, q.timeLimit)
}

// GetResult - Get quiz result
// @return score, total quiz items
func (q *quiz) GetResult() (int, int) {
	return q.score, q.total
}

/**************************************************************************
 *
 *	Private helper methods
 *
 *************************************************************************/

// getQuizItems - Get quiz items from CSV file.
// Validate if quiz item count is within game rule limit.
//
// @return quiz items
func getQuizItems(csvFileName *string) (*[]problem, error) {

	f, err := os.Open(*csvFileName)
	if err != nil {
		//log.Trace(err.Error())
		return nil, err
	}

	// Load and extract lines from CSV file

	r := csv.NewReader(f)

	// read and store all lines
	lines, err := r.ReadAll()

	// check if read ok
	if err != nil {
		return nil, err
	}

	// Validate quiz items limit

	if len(lines) > MaxQuizItems {
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

// runGame - Run the game, continue until all quiz items are posed and answered,
// or until time runs out.
func runGame(q *quiz, problems *[]problem, timeLimit int) error {

	var timer = time.NewTimer(time.Duration(timeLimit) * time.Second)

	// Loop through quiz items
	var timesUp bool = false

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
			fmt.Printf("\n\nYour time is up. Game over!\n")
			timesUp = true
			break
		case answer := <-answerChan:
			if answer == p.a {
				q.score++
			}
		}

		if timesUp {
			break
		}

	} // end for

	return nil
}
