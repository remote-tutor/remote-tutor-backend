package quizzes

import (
	authController "backend/controllers/auth"
	quizzesDBInteractions "backend/database/quizzes"
	usersDBInteractions "backend/database/users"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"net/http"

	"github.com/jinzhu/now"
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

	err := quizzesDBInteractions.CreateQuiz(&quiz)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (quiz not created), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Quiz created successfully",
		"quiz":    quiz,
	})
}

// UpdateQuiz updtes an existing quiz
func UpdateQuiz(c echo.Context) error {
	quizID := utils.ConvertToUInt(c.FormValue("id"))

	quiz := quizzesDBInteractions.GetQuizByID(quizID)
	quiz.Title = c.FormValue("title")
	quiz.Year = utils.ConvertToInt(c.FormValue("year"))
	quiz.StartTime = utils.ConvertToTime(c.FormValue("startTime"))
	quiz.EndTime = utils.ConvertToTime(c.FormValue("endTime"))

	err := quizzesDBInteractions.UpdateQuiz(&quiz)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (quiz not updated), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Quiz updated successfully",
		"quiz":    quiz,
	})
}

//GetPastQuizzes retrieves list of past quizzes for the logged in user
func GetPastQuizzes(c echo.Context) error {
	userid := authController.FetchLoggedInUserID(c)
	user := usersDBInteractions.GetUserByUserID(userid)
	var year int
	if user.Admin {
		year = utils.ConvertToInt(c.QueryParam("year"))
	} else {
		year = user.Year
	}
	pastQuizzes, totalQuizzes := quizzesDBInteractions.GetPastQuizzes(c, year)
	return c.JSON(http.StatusOK, echo.Map{
		"pastQuizzes":  pastQuizzes,
		"totalQuizzes": totalQuizzes,
	})
}

//GetFutureQuizzes retrieves list of future quizzes for the logged in user
func GetFutureQuizzes(c echo.Context) error {
	userid := authController.FetchLoggedInUserID(c)
	user := usersDBInteractions.GetUserByUserID(userid)
	var year int
	if user.Admin {
		year = utils.ConvertToInt(c.QueryParam("year"))
	} else {
		year = user.Year
	}
	futureQuizzes, totalQuizzes := quizzesDBInteractions.GetFutureQuizzes(c, year)
	return c.JSON(http.StatusOK, echo.Map{
		"futureQuizzes": futureQuizzes,
		"totalQuizzes":  totalQuizzes,
	})
}

//GetCurrentQuizzes retrieves list of current quizzes for the logged in user
func GetCurrentQuizzes(c echo.Context) error {
	userid := authController.FetchLoggedInUserID(c)
	user := usersDBInteractions.GetUserByUserID(userid)
	var year int
	if user.Admin {
		year = utils.ConvertToInt(c.QueryParam("year"))
	} else {
		year = user.Year
	}
	currentQuizzes, totalQuizzes := quizzesDBInteractions.GetCurrentQuizzes(c, year)
	return c.JSON(http.StatusOK, echo.Map{
		"currentQuizzes": currentQuizzes,
		"totalQuizzes":   totalQuizzes,
	})
}

// DeleteQuiz deletes the quiz that the user selects
func DeleteQuiz(c echo.Context) error {
	quizID := utils.ConvertToUInt(c.FormValue("id"))
	quiz := quizzesDBInteractions.GetQuizByID(quizID)
	err := quizzesDBInteractions.DeleteQuiz(&quiz)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (quiz not deleted), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Quiz deleted successfully",
	})
}

// GetQuizByID returns a quiz with the specific passed id
func GetQuizByID(c echo.Context) error {
	quizID := utils.ConvertToUInt(c.QueryParam("id"))
	quiz := quizzesDBInteractions.GetQuizByID(quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"quiz": quiz,
	})
}

// GetQuizzesByMonthAndYear gets the quizzes within amonth period.
func GetQuizzesByMonthAndYear(c echo.Context) error {
	isAdmin := authController.FetchLoggedInUserAdminStatus(c)
	var year int
	if isAdmin {
		year = utils.ConvertToInt(c.QueryParam("year"))
	} else {
		year = authController.FetchLoggedInUserYear(c)
	}
	date := utils.ConvertToTime(c.QueryParam("date"))
	endOfMonth := now.With(date).EndOfMonth()
	startOfMonth := now.With(date).BeginningOfMonth()
	quizzes := quizzesDBInteractions.GetQuizzesByMonthAndYear(year, startOfMonth, endOfMonth)
	return c.JSON(http.StatusOK, echo.Map{
		"quizzes": quizzes,
	})
}
