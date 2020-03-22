package types

type Round struct {
	Questions *QuestionBank
	Score int
	CurrentQuestion int
	AnsweredQuestions int
	TotalQuestions int
}

// NewRound returns a new round, it takes in questionBank pointer as an input
func NewRound(questions *QuestionBank) *Round{
	round := new(Round)
	round.Questions = questions
	round.Score = 0
	round.AnsweredQuestions = 0
	round.TotalQuestions, _ = questions.Size()
	return round
}

// Returns the question at index
func (round *Round) QuestionAtIndex(index int) *Question{
	question, _ := round.Questions.QuestionAtIndex(index)
	return question
}