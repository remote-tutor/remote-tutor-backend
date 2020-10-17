package utils

import (
	classUsersModel "backend/models/organizations"
	quizzesModel "backend/models/quizzes"
	"math/rand"
	"sort"
)

// ShuffleQuestions takes the questions array and shuffle it
func ShuffleQuestions(questions []quizzesModel.MCQ, classUser *classUsersModel.ClassUser) {
	if classUser.Admin { // if the current classUser is an admin, then sort the questions in the order of creation
		sort.Slice(questions, func(i, j int) bool {
			return questions[i].ID < questions[j].ID
		})
	} else { // if the current classUser is a student, shuffle the questions using his created_at timestamp
		rand.Seed(classUser.CreatedAt.UnixNano())
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
