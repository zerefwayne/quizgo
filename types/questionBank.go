package types

import "fmt"

type QuestionBank struct {
	questions []*Question
	numberOfQuestions int
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
		fmt.Printf("%d. %s (%d Points)\n", idx + 1, val.Statement, val.Score)
	}
}