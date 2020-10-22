package quizzes

import (
	authController "backend/controllers/auth"
	quizzesDBInteractions "backend/database/quizzes"
	quizzesModel "backend/models/quizzes"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

// GetSubmissionsByQuizAndUser retrieves all mcq and long answer submission for a specific quiz for a specific user
func GetSubmissionsByQuizAndUser(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	quizID := utils.ConvertToUInt(c.QueryParam("quizID"))

	mcqSubmissions := quizzesDBInteractions.GetMCQSubmissionsByQuizID(userID, quizID)
	longAnswerSubmissions := quizzesDBInteractions.GetLongAnswerSubmissionsByQuizID(userID, quizID)

	return c.JSON(http.StatusOK, echo.Map{
		"mcqSubmissions":        mcqSubmissions,
		"longAnswerSubmissions": longAnswerSubmissions,
	})
}

// CreateOrUpdateMCQSubmission updates a previous submission for an mcq question.
func CreateOrUpdateMCQSubmission(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	mcqID := utils.ConvertToUInt(c.FormValue("mcqID"))
	userResult := utils.ConvertToUInt(c.FormValue("userResult"))

	submission := quizzesModel.Submission{
		UserID: userID,
	}
	mcqSubmission := quizzesModel.MCQSubmission{
		Submission: submission,
		UserResult: userResult,
		MCQID:      mcqID,
	}

	mcqQuestion := quizzesDBInteractions.GetMCQByID(mcqID)
	if userResult == mcqQuestion.CorrectAnswer {
		mcqSubmission.Submission.Grade = mcqQuestion.TotalMark
	} else {
		mcqSubmission.Submission.Grade = 0
	}
	err := quizzesDBInteractions.CreateOrUpdateMCQSubmission(&mcqSubmission)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (submission not saved), please try again",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"mcqSubmission": mcqSubmission,
	})
}
