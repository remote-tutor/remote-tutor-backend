package utils

import (
	quizzesModel "backend/models/quizzes"
	usersModel "backend/models/users"
	"math/rand"
	"sort"
)

// ShuffleQuestions takes the questions array and shuffle it
func ShuffleQuestions(questions []quizzesModel.MCQ, user *usersModel.User) {
	if user.Admin { // if the current user is an admin, then sort the questions in the order of creation
		sort.Slice(questions, func(i, j int) bool {
			return questions[i].ID < questions[j].ID
		})
	} else { // if the current user is a student, shuffle the questions using his created_at timestamp
		rand.Seed(user.CreatedAt.UnixNano())
		rand.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}

	for i := 0; i < len(questions); i++ { // in both cases order the choices to display (A, B, C, D)
		orderChoices(questions[i].Choices)
	}
}

func orderChoices(choices []quizzesModel.Choice) {
	sort.Slice(choices, func(i, j int) bool {
		return choices[i].Text < choices[j].Text
	})
}
