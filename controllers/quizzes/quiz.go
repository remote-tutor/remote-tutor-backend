package quizzes

import (
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
