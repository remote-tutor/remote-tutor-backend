package utils

import (
	quizzesModel "backend/models/quizzes"
	"math/rand"
	"time"
)

// ShuffleQuestions takes the questions array and shuffle it
func ShuffleQuestions(questions *[]quizzesModel.MCQ) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*questions), func(i, j int) {
		(*questions)[i], (*questions)[j] = (*questions)[j], (*questions)[i]
	})
	for i := 0; i < len(*questions); i++ {
		shuffleChoices(&(*questions)[i].Choices)
	}
}

func shuffleChoices(choices *[]quizzesModel.Choice) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*choices), func(i, j int) {
		(*choices)[i], (*choices)[j] = (*choices)[j], (*choices)[i]
	})
}
