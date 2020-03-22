package types

// Question implements the basic type into which the CSV will be read into. It has a question, answer, score
type Question struct {
	Statement string
	Answer string
	Score int
}

func (question *Question) CheckAnswer(givenAnswer string) (bool, int){
	if givenAnswer == question.Answer {
		return true, question.Score
	}
	return false, 0
}

