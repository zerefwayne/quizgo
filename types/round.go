package types

type Round struct {
	questions *QuestionBank
	score int
	currentQuestion int
	answeredQuestions int
	totalQuestions int
}

// NewRound returns a new round, it takes in questionBank pointer as an input
func NewRound(questions *QuestionBank) *Round{
	round := new(Round)
	round.questions = questions
	round.score = 0
	round.currentQuestion = 0
	round.answeredQuestions = 0
	round.totalQuestions, _ = questions.Size()

	return round
}
