package utils

import (
	"github.com/zerefwayne/quizgo/types"
	"strconv"
)

// LineToOWA Parses Line to OneWordAnswerType
func LineToOWA(line []string) *types.Question{
	score, _ := strconv.ParseInt(line[2], 10, 64)
	question := &types.Question{Statement: line[0], Answer: line[1], Score: int(score)}
	return question
}
