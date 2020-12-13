package quizzes

import (
	quizzesModel "backend/models/quizzes"
	"time"
)

type Grades struct {
	Grades           []map[string]interface{}
	GradesOnly       [][]int
	TeacherName      string
	ClassName        string
	StartDate        time.Time
	EndDate          time.Time
	Quizzes          []quizzesModel.Quiz
	QuizzesTotalMark int
}
