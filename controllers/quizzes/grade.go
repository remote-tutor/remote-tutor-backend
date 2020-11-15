package quizzes

import (
	authController "backend/controllers/auth"
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// CreateQuizGrade creates a quiz grade record for the logged-in user
func CreateQuizGrade(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))
	quiz := quizzesDBInteractions.GetQuizByID(quizID)
	quizGrade := quizzesModel.QuizGrade{
		Grade:     0,
		QuizID:    quizID,
		UserID:    userID,
		StartAt:   time.Now(),
		ValidTill: time.Now().Add(time.Duration(quiz.StudentTime) * time.Minute),
	}
	err := quizzesDBInteractions.CreateGrade(&quizGrade)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred, please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"validTill": quizGrade.ValidTill,
	})
}

func GetStudentRemainingTime(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	quizID := utils.ConvertToUInt(c.QueryParam("quizID"))

	remainingTime, recordFound := quizzesDBInteractions.GetStudentRemainingTime(userID, quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"studentTime": remainingTime,
		"recordFound": recordFound,
	})
}

// GetGradesByQuiz Fetches logged-in user's grade for a quiz
func GetGradesByQuiz(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	quizID := utils.ConvertToUInt(c.QueryParam("quizID"))

	quizGrade := quizzesDBInteractions.GetGradesByQuizID(userID, quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"quizGrade": []quizzesModel.QuizGrade{quizGrade},
	})
}

// GetGradesByQuizForAllUsers fetches all class grades for a quiz
func GetGradesByQuizForAllUsers(c echo.Context) error {
	quizID := utils.ConvertToUInt(c.QueryParam("quizID"))
	quizGrades := quizzesDBInteractions.GetGradesByQuizForAllUsers(quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"quizGrades": quizGrades,
	})
}
