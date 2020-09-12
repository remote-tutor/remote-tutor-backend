package utils

import (
	quizzesModel "backend/models/quizzes"
	"math/rand"
	"sort"
	"time"
)

// ShuffleQuestions takes the questions array and shuffle it
func ShuffleQuestions(questions []quizzesModel.MCQ) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	for i := 0; i < len(questions); i++ {
		orderChoices(questions[i].Choices)
	}
}

func orderChoices(choices []quizzesModel.Choice) {
	sort.Slice(choices, func(i, j int) bool {
		return choices[i].Text < choices[j].Text
	})
}
