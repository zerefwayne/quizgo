package types

import (
	"encoding/json"
	"fmt"
	"log"
)

type QuestionBank struct {
	questions         []*Question
	numberOfQuestions int
}

// Takes a Questions input and returns a QuestionsBank with those questions
func NewQuestionBank(questions Questions) QuestionBank {
	questionBank := new(QuestionBank)
	for i := range questions.Questions {
		questionBank.questions = append(questionBank.questions, &questions.Questions[i])
	}
	questionBank.numberOfQuestions = len(questions.Questions)
	return *questionBank
}

// AddQuestion adds a question to the question bank from the read query
func (questionBank *QuestionBank) AddQuestion(question *Question) error {
	questionBank.numberOfQuestions++
	questionBank.questions = append(questionBank.questions, question)
	return nil
}

// Size returns the size of the question bank
func (questionBank *QuestionBank) Size() (int, error) {
	return questionBank.numberOfQuestions, nil
}

// DisplayQuestions function prints the questions in the question bank
func (questionBank *QuestionBank) DisplayQuestions() {
	for idx, val := range questionBank.questions {
		fmt.Printf("%d. %s (%d Points)\n", idx+1, val.Statement, val.Score)
	}
}

// Returns the question at index i
func (questionBank *QuestionBank) QuestionAtIndex(index int) (*Question, error) {
	return questionBank.questions[index], nil
}

// Takes a file path and file type and returns a QuestionBank with the questions
// from the file. Supports csv and json formatted files
func InitialiseQuestionBank(filePath string, sourceType string) (*QuestionBank, error) {

	var err error
	questionBank := new(QuestionBank)

	switch sourceType {
	case "csv":
		{
			err = getQuestionBankFromCSV(filePath, questionBank)
			if err != nil {
				return nil, err
			}
			break
		}
	case "json":
		{
			err = getQuestionBankFromJSON(filePath, questionBank)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	return questionBank, nil
}

// Gets contents from a csv file path and parses them into a QuestionBank
func getQuestionBankFromCSV(filePath string, questionBank *QuestionBank) error {

	csv, err := getCSVFileContents(filePath)

	if err != nil {
		log.Fatal(err)
		return err
	}
	err = parseStrings(csv, questionBank)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// Gets contents from a json file path and unmarshals them into a QuestionBank
func getQuestionBankFromJSON(filePath string, questionBank *QuestionBank) error {

	byteValue, err := getJSONFileContents(filePath)

	if err != nil {
		log.Fatal(err)
		return err
	}

	// we initialize our Questions array
	var questions Questions

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &questions)

	if err != nil {
		log.Fatal(err)
		return err
	}

	*questionBank = NewQuestionBank(questions)
	return nil
}
