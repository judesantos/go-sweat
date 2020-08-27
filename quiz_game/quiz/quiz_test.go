package quiz

// quiz_test.go

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuizConstants(t *testing.T) {

	assert.EqualValues(t, "../csvs/deploy.csv", DefaultCsv,
		"DefaultCsv should be './csvs/deploy.csv'")

	assert.EqualValues(t, 100, MaxQuizItems,
		"MaxQuizItems value should be 100")

	assert.EqualValues(t, 30, DefaultTimeLimit,
		"DefaultTimeLimit value should be 30")
}

func TestCreateNewQuizWithDefaultArgs(t *testing.T) {

	quiz, err := NewQuiz(DefaultCsv, DefaultTimeLimit)

	assert.Equal(t, nil, err, "Error creating new quiz")
	assert.NotEqual(t, nil, quiz, "Quiz returned nil")
}

func TestCreateNewQuizMissingCsv(t *testing.T) {

	quiz, err := NewQuiz("missing.csv", DefaultTimeLimit)

	assert.False(t, nil == err, "Error creating new quiz")
	assert.True(t, nil == quiz, "Quiz returned nil")
}

func TestTimeoutQuiz(t *testing.T) {

	const timeout int = 15 // timeout in seconds

	quiz, err := NewQuiz(DefaultCsv, timeout)

	assert.True(t, nil == err, "Error creating new quiz")
	assert.False(t, nil == quiz, "NewQuiz returned nil, quiz object expected")

	var stdin bytes.Buffer
	stdin.Write([]byte("atay\n"))

	// test me

	err = quiz.Start()

	assert.True(t, nil != err, "Start is expected to return with error")
	//assert.EqualValues(t, "Your time is up. Game over!", err.Error(),
	//	"Exprected error message mismatch")
}
