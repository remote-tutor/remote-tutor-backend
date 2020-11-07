package quizzes

import (
	paginationController "backend/controllers/pagination"
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"net/http"

	"github.com/jinzhu/now"
	"github.com/labstack/echo"
)

// CreateQuiz adds new quiz
func CreateQuiz(c echo.Context) error {
	title := c.FormValue("title")
	class := c.FormValue("selectedClass")
	startTime := utils.ConvertToTime(c.FormValue("startTime"))
	endTime := utils.ConvertToTime(c.FormValue("endTime"))

	quiz := quizzesModel.Quiz{
		Title:     title,
		ClassHash: class,
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
	quiz.ClassHash = c.FormValue("selectedClass")
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
	class := c.QueryParam("selectedClass")
	paginationData := paginationController.ExtractPaginationData(c)
	pastQuizzes, totalQuizzes := quizzesDBInteractions.GetPastQuizzes(paginationData, class)
	return c.JSON(http.StatusOK, echo.Map{
		"pastQuizzes":  pastQuizzes,
		"totalQuizzes": totalQuizzes,
	})
}

//GetFutureQuizzes retrieves list of future quizzes for the logged in user
func GetFutureQuizzes(c echo.Context) error {
	class := c.QueryParam("selectedClass")
	paginationData := paginationController.ExtractPaginationData(c)
	futureQuizzes, totalQuizzes := quizzesDBInteractions.GetFutureQuizzes(paginationData, class)
	return c.JSON(http.StatusOK, echo.Map{
		"futureQuizzes": futureQuizzes,
		"totalQuizzes":  totalQuizzes,
	})
}

//GetCurrentQuizzes retrieves list of current quizzes for the logged in user
func GetCurrentQuizzes(c echo.Context) error {
	class := c.QueryParam("selectedClass")
	paginationData := paginationController.ExtractPaginationData(c)
	currentQuizzes, totalQuizzes := quizzesDBInteractions.GetCurrentQuizzes(paginationData, class)
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

// GetQuizByHash returns a quiz with the specific passed id
func GetQuizByHash(c echo.Context) error {
	hash := c.QueryParam("quizHash")
	quiz := quizzesDBInteractions.GetQuizByHash(hash)
	return c.JSON(http.StatusOK, echo.Map{
		"quiz": quiz,
	})
}

// GetQuizzesByClassMonthAndYear gets the quizzes within a month period.
func GetQuizzesByClassMonthAndYear(c echo.Context) error {
	class := c.QueryParam("selectedClass")
	date := utils.ConvertToTime(c.QueryParam("date"))
	endOfMonth := now.With(date).EndOfMonth()
	startOfMonth := now.With(date).BeginningOfMonth()
	quizzes := quizzesDBInteractions.GetQuizzesByClassAndMonthAndYear(class, startOfMonth, endOfMonth)
	return c.JSON(http.StatusOK, echo.Map{
		"quizzes": quizzes,
	})
}
