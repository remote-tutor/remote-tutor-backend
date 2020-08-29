package quizzes

import (
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

//CreateMCQ creates a new mcq question for a quiz
func CreateMCQ(c echo.Context) error {
	correctAnswer := utils.ConvertToUInt(c.FormValue("correctAnswer"))
	question := constructQuestion(c)
	mcq := quizzesModel.MCQ{
		Question:      question,
		CorrectAnswer: correctAnswer,
	}
	quizzesDBInteractions.CreateMCQ(&mcq)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "MCQ question created successfully",
		"mcq":     mcq,
	})
}

//CreateLongAnswer creates a new long answer question for a quiz
func CreateLongAnswer(c echo.Context) error {
	correctAnswer := c.FormValue("correctAnswer")
	question := constructQuestion(c)
	longAnswer := quizzesModel.LongAnswer{
		Question:      question,
		CorrectAnswer: correctAnswer,
	}
	quizzesDBInteractions.CreateLongAnswer(&longAnswer)
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "Long Answer question created successfully",
		"longAnswer": longAnswer,
	})
}

//UpdateMCQ updates mcq question for a quiz
func UpdateMCQ(c echo.Context) error {
	mcq := quizzesDBInteractions.GetMCQByID(utils.ConvertToUInt(c.FormValue("id")))
	mcq.CorrectAnswer = utils.ConvertToUInt(c.FormValue("correctAnswer"))
	mcq.Question.TotalMark = utils.ConvertToInt(c.FormValue("totalMark"))
	mcq.Question.Text = c.FormValue("text")
	return c.JSON(http.StatusOK, echo.Map{
		"message": "MCQ question created successfully",
		"mcq":     mcq,
	})
}

//UpdateLongAnswer updates long answer question for a quiz
func UpdateLongAnswer(c echo.Context) error {
	longAnswer := quizzesDBInteractions.GetLongAnswerByID(utils.ConvertToUInt(c.FormValue("id")))
	longAnswer.CorrectAnswer = c.FormValue("correctAnswer")
	longAnswer.Question.TotalMark = utils.ConvertToInt(c.FormValue("totalMark"))
	longAnswer.Question.Text = c.FormValue("text")
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "LongAnswer question created successfully",
		"longAnswer": longAnswer,
	})
}

//GetQuestionsByQuiz retrieves all questions for a quiz
func GetQuestionsByQuiz(c echo.Context) error {
	quizID := utils.ConvertToUInt(c.QueryParam("quizID"))
	mcqs := quizzesDBInteractions.GetMCQQuestionsByQuiz(quizID)
	longAnswers := quizzesDBInteractions.GetLongAnswerQuestionsByQuiz(quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"mcqs":        mcqs,
		"longAnswers": longAnswers,
	})
}

func constructQuestion(c echo.Context) quizzesModel.Question {
	text := c.FormValue("text")
	totalMark := utils.ConvertToInt(c.FormValue("totalMark"))
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))

	question := quizzesModel.Question{
		Text:      text,
		TotalMark: totalMark,
		QuizID:    quizID,
	}
	return question
}
