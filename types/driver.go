package types

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"quizgo/repositories"
)

type Driver struct {
	round *Round
}

func NewDriver() (*Driver, error) {

	driver := new(Driver)

	var filePath string

	flag.StringVar(&filePath, "csv", "./problems/addition.csv",
		"a csv file in the format of 'question,answer'")
	flag.Parse()

	questionBank, err := InitialiseQuestionBank(filePath)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	driver.round = NewRound(questionBank)
	return driver, nil
}

func (driver *Driver) Start() (int, error) {

	round := driver.round
	totalQuestions := round.TotalQuestions
	localDb, err := sql.Open("sqlite3", "quizgo.sqlite")
	if err != nil {
		return 0, err
	}

	user, err := repositories.Get(localDb, 1)
	fmt.Println("----------Player Info -----------")
	fmt.Println("Score Loaded:", user.Score)
	fmt.Println("Player:", user.Name)
	fmt.Println("Last game:", user.LastGame.Format("02-01-2006"))
	fmt.Println("--------------------------------- \n")

	if err != nil {
		return 0, err
	}

	for currentQuestion := 0; currentQuestion < totalQuestions; currentQuestion++ {

		question := round.QuestionAtIndex(currentQuestion)
		fmt.Printf("Question %d/%d:\n\n", currentQuestion+1, totalQuestions)
		fmt.Printf("%s\n\n", question.Statement)

		answer := ""

		fmt.Scanf("%s", &answer)
		fmt.Println()

		isCorrect, pointsScored := question.CheckAnswer(answer)

		round.Score += pointsScored
		user.Score += pointsScored

		fmt.Println("\033[H\033[2J")

		if isCorrect {
			fmt.Printf("Correct Answer! You scored %d points. Round points: %d Total points: %d\n\n", pointsScored, round.Score, user.Score)
		} else {
			fmt.Printf("Incorrect Answer! Total Points: %d\n\n", round.Score)
		}

	}

	repositories.Update(localDb, user.Score, 1) // updating score and last_game played

	return user.Score, nil

}
