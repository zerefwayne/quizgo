package types

import (
	"flag"
	"fmt"
	"log"
)

type Driver struct {
	round *Round
}

func NewDriver() (*Driver, error) {

	driver := new(Driver)

	// Path of the input file
	var filePath string
	// Type of the input file (csv,json)
	var sourceType string

	flag.StringVar(&filePath, "file", "./problems/addition.csv",
		"a csv file in the format of 'question,answer'")

	// acceptable values are csv and json
	flag.StringVar(&sourceType, "type", "csv",
		"the type of input used for the questions")
	flag.Parse()

	questionBank, err := InitialiseQuestionBank(filePath, sourceType)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	driver.round = NewRound(questionBank)
	return driver, nil
}

func (driver *Driver) Start() {

	round := driver.round
	totalQuestions := round.TotalQuestions

	for currentQuestion := 0; currentQuestion < totalQuestions; currentQuestion++ {

		question := round.QuestionAtIndex(currentQuestion)
		fmt.Printf("Question %d/%d:\n\n", currentQuestion+1, totalQuestions)
		fmt.Printf("%s\n\n", question.Statement)

		answer := ""

		_, err := fmt.Scanf("%s", &answer)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println()

		isCorrect, pointsScored := question.CheckAnswer(answer)

		round.Score += pointsScored

		fmt.Println("\033[H\033[2J")

		if isCorrect {
			fmt.Printf("Correct Answer! You scored %d points. Total points: %d\n\n", pointsScored, round.Score)
		} else {
			fmt.Printf("Incorrect Answer! Total Points: %d\n\n", round.Score)
		}

	}

}
