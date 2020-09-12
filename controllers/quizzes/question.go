package quizzes

import (
	authController "backend/controllers/auth"
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"
)

//CreateMCQ creates a new mcq question for a quiz
func CreateMCQ(c echo.Context) error {
	question, err := constructQuestion(c)
	if err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, echo.Map{
			"message": "Error uploading the image, please try again later",
		})
	}

	correctAnswer := utils.ConvertToUInt(c.FormValue("correctAnswer"))
	mcq := quizzesModel.MCQ{
		Question:      question,
		CorrectAnswer: correctAnswer,
	}
	quizzesDBInteractions.CreateMCQ(&mcq)

	err = createImageFile(&mcq.Question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "An unexpected error occured when tring to save the image, please try again later",
		})
	}
	updateQuizTotalMark(1, c)
	return c.JSON(http.StatusOK, echo.Map{
		"mcq": mcq,
	})
}

//CreateLongAnswer creates a new long answer question for a quiz
func CreateLongAnswer(c echo.Context) error {
	question, err := constructQuestion(c)
	if err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, echo.Map{
			"message": "Error uploading the image, please try again later",
		})
	}

	correctAnswer := c.FormValue("correctAnswer")
	longAnswer := quizzesModel.LongAnswer{
		Question:      question,
		CorrectAnswer: correctAnswer,
	}
	quizzesDBInteractions.CreateLongAnswer(&longAnswer)

	updateQuizTotalMark(question.TotalMark, c)
	return c.JSON(http.StatusOK, echo.Map{
		"longAnswer": longAnswer,
	})
}

//UpdateMCQ updates mcq question for a quiz
func UpdateMCQ(c echo.Context) error {
	question, err := constructQuestion(c)
	if err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, echo.Map{
			"message": "Error uploading the image, please try again later",
		})
	}

	mcq := quizzesDBInteractions.GetMCQByID(utils.ConvertToUInt(c.FormValue("id")))
	mcq.CorrectAnswer = utils.ConvertToUInt(c.FormValue("correctAnswer"))
	mcq.Question.TotalMark = question.TotalMark
	mcq.Question.Text = question.Text
	mcq.Question.ImagePath = question.ImagePath
	mcq.Question.Image = question.Image

	err = createImageFile(&mcq.Question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "An unexpected error occured when tring to save the image, please try again later",
		})
	}
	quizzesDBInteractions.UpdateMCQ(&mcq)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "MCQ question updated successfully",
		"mcq":     mcq,
	})
}

//UpdateLongAnswer updates long answer question for a quiz
func UpdateLongAnswer(c echo.Context) error {
	longAnswer := quizzesDBInteractions.GetLongAnswerByID(utils.ConvertToUInt(c.FormValue("id")))
	oldTotalMark := longAnswer.Question.TotalMark
	longAnswer.CorrectAnswer = c.FormValue("correctAnswer")
	longAnswer.Question.TotalMark = utils.ConvertToInt(c.FormValue("totalMark"))
	longAnswer.Question.Text = c.FormValue("text")
	quizzesDBInteractions.UpdateLongAnswer(&longAnswer)

	updateQuizTotalMark(longAnswer.Question.TotalMark-oldTotalMark, c)
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "LongAnswer question created successfully",
		"longAnswer": longAnswer,
	})
}

// DeleteMCQ deletes mcq question for a quiz
func DeleteMCQ(c echo.Context) error {
	mcq := quizzesDBInteractions.GetMCQByID(utils.ConvertToUInt(c.FormValue("id")))
	quizzesDBInteractions.DeleteMCQ(&mcq)

	updateQuizTotalMark(-1, c)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "MCQ question deleted successfully",
	})
}

// DeleteLongAnswer deletes long answer question for a quiz
func DeleteLongAnswer(c echo.Context) error {
	longAnswer := quizzesDBInteractions.GetLongAnswerByID(utils.ConvertToUInt(c.FormValue("id")))
	quizzesDBInteractions.DeleteLongAnswer(&longAnswer)

	updateQuizTotalMark(-longAnswer.Question.TotalMark, c)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Long Answer question deleted successfully",
	})
}

//GetQuestionsByQuiz retrieves all questions for a quiz
func GetQuestionsByQuiz(c echo.Context) error {
	quizID := utils.ConvertToUInt(c.QueryParam("quizID"))
	mcqs := quizzesDBInteractions.GetMCQsByQuiz(quizID)
	quiz := quizzesDBInteractions.GetQuizByID(quizID)
	isAdmin := authController.FetchLoggedInUserAdminStatus(c)
	if !isAdmin && time.Now().Before(quiz.EndTime) {
		utils.ShuffleQuestions(&mcqs)
	}
	longAnswers := quizzesDBInteractions.GetLongAnswersByQuiz(quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"mcqs":        mcqs,
		"longAnswers": longAnswers,
	})
}

func constructQuestion(c echo.Context) (quizzesModel.Question, error) {
	text := c.FormValue("text")
	totalMark := utils.ConvertToInt(c.FormValue("totalMark"))
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))
	question := quizzesModel.Question{
		Text:      text,
		TotalMark: totalMark,
		QuizID:    quizID,
	}

	image, err := c.FormFile("image")
	if err != nil {
		return question, err
	}
	src, err := image.Open()
	if err != nil {
		return question, err
	}
	defer src.Close()
	question.Image = src
	question.ImagePath = image.Filename[strings.LastIndex(image.Filename, ".")+1:]

	return question, nil
}

func createImageFile(question *quizzesModel.Question) error {
	createDirectoryIfNotExist(fmt.Sprintf("quizzesImages/quiz %d", question.QuizID))

	question.ImagePath = fmt.Sprintf("quizzesImages/quiz %d/question %d.%s",
		question.QuizID, question.ID, question.ImagePath)
	dst, err := os.Create(question.ImagePath)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, question.Image); err != nil {
		return err
	}
	return nil
}

func createDirectoryIfNotExist(directoryName string) {
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		os.Mkdir(directoryName, os.FileMode(int(0777)))
	}
}

func updateQuizTotalMark(markDifference int, c echo.Context) {
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))
	quiz := quizzesDBInteractions.GetQuizByID(quizID)
	quiz.TotalMark += markDifference
	quizzesDBInteractions.UpdateQuiz(&quiz)
}
