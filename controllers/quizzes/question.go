package quizzes

import (
	authController "backend/controllers/auth"
	questionsFiles "backend/controllers/files/quizzes"
	classUsersDBInteractions "backend/database/organizations"
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

//CreateMCQ creates a new mcq question for a quiz
func CreateMCQ(c echo.Context) error {
	question := constructQuestion(c)
	mcq := quizzesModel.MCQ{
		Question:      question,
	}
	err := quizzesDBInteractions.CreateMCQ(&mcq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (MCQ not created), please try again",
		})
	}

	quiz := quizzesDBInteractions.GetQuizByID(mcq.QuizID)
	mcq.Quiz = quiz
	class := classUsersDBInteractions.GetClassByHash(quiz.ClassHash)
	imageFilePath, err := questionsFiles.UploadQuestionImage(&c, &mcq, &class)
	if err != nil {
		quizzesDBInteractions.DeleteMCQ(&mcq)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "An unexpected error occurred when trying to save the image, please try again later",
		})
	}
	mcq.ImagePath = imageFilePath
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
	question := constructQuestion(c)

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
	question := constructQuestion(c)

	mcq := quizzesDBInteractions.GetMCQByID(utils.ConvertToUInt(c.FormValue("id")))
	mcq.CorrectAnswer = utils.ConvertToUInt(c.FormValue("correctAnswer"))
	mcq.Question.TotalMark = question.TotalMark
	mcq.Question.Text = question.Text
	class := classUsersDBInteractions.GetClassByHash(mcq.Quiz.ClassHash)
	imageFilePath, err := questionsFiles.UploadQuestionImage(&c, &mcq, &class)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "An unexpected error occurred when trying to save the image, please try again later",
		})
	}
	mcq.ImagePath = imageFilePath

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
	class := classUsersDBInteractions.GetClassByHash(mcq.Quiz.ClassHash)
	err := questionsFiles.DeleteQuestionImage(&mcq, &class)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (MCQ not deleted), please try again",
		})
	}
	err = quizzesDBInteractions.DeleteMCQ(&mcq)
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
