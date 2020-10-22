package quizzes

import (
	authController "backend/controllers/auth"
	classUsersDBInteractions "backend/database/organizations"
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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
	err = quizzesDBInteractions.CreateMCQ(&mcq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (MCQ not created), please try again",
		})
	}

	err = createImageFile(&mcq.Question, mcq.Question.ImagePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "An unexpected error occurred when trying to save the image, please try again later",
		})
	}
	err = quizzesDBInteractions.UpdateMCQ(&mcq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (MCQ not created), please try again",
		})
	}
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
	if question.ImagePath != "" {
		mcq.Question.ImagePath = question.ImagePath
		mcq.Question.Image = question.Image
	}

	err = createImageFile(&mcq.Question, question.ImagePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "An unexpected error occured when tring to save the image, please try again later",
		})
	}

	submissions := quizzesDBInteractions.GetMCQSubmissionsByQuestionID(mcq.ID)
	for _, submission := range submissions {
		if submission.UserResult == mcq.CorrectAnswer {
			submission.Grade = mcq.TotalMark
		} else {
			submission.Grade = 0
		}
		quizzesDBInteractions.CreateOrUpdateMCQSubmission(&submission)
	}
	err = quizzesDBInteractions.UpdateMCQ(&mcq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (MCQ not updated), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "MCQ question updated successfully",
		"mcq":     mcq,
	})
}

//UpdateLongAnswer updates long answer question for a quiz
func UpdateLongAnswer(c echo.Context) error {
	longAnswer := quizzesDBInteractions.GetLongAnswerByID(utils.ConvertToUInt(c.FormValue("id")))
	longAnswer.CorrectAnswer = c.FormValue("correctAnswer")
	longAnswer.Question.TotalMark = utils.ConvertToInt(c.FormValue("totalMark"))
	longAnswer.Question.Text = c.FormValue("text")
	quizzesDBInteractions.UpdateLongAnswer(&longAnswer)

	return c.JSON(http.StatusOK, echo.Map{
		"message":    "LongAnswer question created successfully",
		"longAnswer": longAnswer,
	})
}

// DeleteMCQ deletes mcq question for a quiz
func DeleteMCQ(c echo.Context) error {
	mcq := quizzesDBInteractions.GetMCQByID(utils.ConvertToUInt(c.FormValue("id")))
	err := quizzesDBInteractions.DeleteMCQ(&mcq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (MCQ not deleted), please try again",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "MCQ question deleted successfully",
	})
}

// DeleteLongAnswer deletes long answer question for a quiz
func DeleteLongAnswer(c echo.Context) error {
	longAnswer := quizzesDBInteractions.GetLongAnswerByID(utils.ConvertToUInt(c.FormValue("id")))
	quizzesDBInteractions.DeleteLongAnswer(&longAnswer)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Long Answer question deleted successfully",
	})
}

//GetQuestionsByQuiz retrieves all questions for a quiz
func GetQuestionsByQuiz(c echo.Context) error {
	quizID := utils.ConvertToUInt(c.QueryParam("quizID"))
	mcqs := quizzesDBInteractions.GetMCQsByQuiz(quizID)
	userID := authController.FetchLoggedInUserID(c)
	class := c.QueryParam("selectedClass")
	classUser := classUsersDBInteractions.GetClassUserByUserIDAndClass(userID, class)
	utils.ShuffleQuestions(mcqs, &classUser)
	longAnswers := quizzesDBInteractions.GetLongAnswersByQuiz(quizID)
	return c.JSON(http.StatusOK, echo.Map{
		"mcqs":        mcqs,
		"longAnswers": longAnswers,
	})
}

// GetQuestionImage gets the image of the specified question
func GetQuestionImage(c echo.Context) error {
	imagePath := c.Param("imagePath")
	quizID := c.Param("quizID")
	questionID := c.Param("questionID")
	fullPath := imagePath + "/" + quizID + "/" + questionID
	if _, err := os.Stat(fullPath); err == nil {
		bytes, err := ioutil.ReadFile(fullPath)
		if err == nil {
			c.Response().Header().Set("Content-Type", "image/*")
			return c.Blob(http.StatusOK, "image/*", bytes)
		}
	}
	return c.JSON(http.StatusInternalServerError, echo.Map{
		"message": "Error while retrieving the file",
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
		if err.Error() == "http: no such file" {
			return question, nil
		}
		return question, err
	}
	src, err := image.Open()
	if err != nil {
		return question, err
	}
	defer src.Close()
	bytes, err := ioutil.ReadAll(src)
	if err != nil {
		return question, err
	}
	question.Image = bytes
	question.ImagePath = image.Filename[strings.LastIndex(image.Filename, ".")+1:]

	return question, nil
}

func createImageFile(question *quizzesModel.Question, newPath string) error {
	if newPath == "" {
		return nil
	}
	createDirectoryIfNotExist(fmt.Sprintf("quizzesImages/quiz %d", question.QuizID))

	question.ImagePath = fmt.Sprintf("quizzesImages/quiz %d/question %d.%s",
		question.QuizID, question.ID, question.ImagePath)
	dst, err := os.Create(question.ImagePath)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = dst.Write(question.Image); err != nil {
		return err
	}
	return nil
}

func createDirectoryIfNotExist(directoryName string) {
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		os.Mkdir(directoryName, os.FileMode(int(0777)))
	}
}
