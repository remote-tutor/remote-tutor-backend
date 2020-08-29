package quizzes

import (
	"backend/controllers"
	dbInteractions "backend/database"
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

// CreateQuiz adds new quiz
func CreateQuiz(c echo.Context) error {
	title := c.FormValue("title")
	year := utils.ConvertToInt(c.FormValue("year"))
	startTime := utils.ConvertToTime(c.FormValue("startTime"))
	endTime := utils.ConvertToTime(c.FormValue("endTime"))

	quiz := quizzesModel.Quiz{
		Title:     title,
		Year:      year,
		StartTime: startTime,
		EndTime:   endTime,
	}

	quizzesDBInteractions.CreateQuiz(&quiz)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Quiz created successfully",
		"quiz":    quiz,
	})
}

//GetPastQuizzes retrieves list of past quizzes for the logged in user
func GetPastQuizzes(c echo.Context) error {
	userid := controllers.FetchLoggedInUserID(c)
	user := dbInteractions.GetUserByUserID(userid)
	pastQuizzes := quizzesDBInteractions.GetPastQuizzes(user.Year)
	return c.JSON(http.StatusOK, echo.Map{
		"pastQuizzes": pastQuizzes,
	})
}

//GetFutureQuizzes retrieves list of future quizzes for the logged in user
func GetFutureQuizzes(c echo.Context) error {
	userid := controllers.FetchLoggedInUserID(c)
	user := dbInteractions.GetUserByUserID(userid)
	futureQuizzes := quizzesDBInteractions.GetFutureQuizzes(user.Year)
	return c.JSON(http.StatusOK, echo.Map{
		"futureQuizzes": futureQuizzes,
	})
}

//GetCurrentQuizzes retrieves list of current quizzes for the logged in user
func GetCurrentQuizzes(c echo.Context) error {
	userid := controllers.FetchLoggedInUserID(c)
	user := dbInteractions.GetUserByUserID(userid)
	currentQuizzes := quizzesDBInteractions.GetCurrentQuizzes(user.Year)
	return c.JSON(http.StatusOK, echo.Map{
		"currentQuizzes": currentQuizzes,
	})
}
