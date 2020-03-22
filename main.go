package main

import (
	"bufio"
	"encoding/csv"
	"github.com/zerefwayne/quizgo/types"
	"github.com/zerefwayne/quizgo/utils"
	"io"
	"log"
	"os"
)

func ParseCSVToQuestionBank(csvFilePath string, questionBank *types.QuestionBank) error {

	csvFile, err := os.Open(csvFilePath)

	if err != nil {
		return err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		question := utils.LineToOWA(line)
		_ = questionBank.AddQuestion(question)
	}

	return nil
}


func main() {

	questionBank := new(types.QuestionBank)

	err := ParseCSVToQuestionBank("./problems.csv", questionBank)

	if err != nil {
		log.Fatal(err)
	}

	questionBank.DisplayQuestions()

}
