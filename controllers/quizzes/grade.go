package quizzes

import (
	authController "backend/controllers/auth"
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

// CreateQuizGrade creates a quiz grade record for the logged-in user
func CreateQuizGrade(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))
	quizGrade := quizzesModel.QuizGrade{
		Grade:  0,
		QuizID: quizID,
		UserID: userID,
	}
	quizzesDBInteractions.CreateGrade(&quizGrade)
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "Quiz grade created successfully",
		"quizGrade": quizGrade,
	})
}

// GetGradesByQuiz Fetches logged-in user's grade for a quiz
func GetGradesByQuiz(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))

	quizGrade := quizzesDBInteractions.GetGradesByQuizID(userID, quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"quizGrade": quizGrade,
	})
}

// GetGradesByQuizForAllUsers fetches all class grades for a quiz
// func GetGradesByQuizForAllUsers(c echo.Context) {

// }
