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
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))
	displayUserResult := utils.ConvertToBool(c.QueryParam("displayUserResult"))
	mcqSubmissions := quizzesDBInteractions.GetMCQSubmissionsByQuizID(userID, quizID)
	longAnswerSubmissions := quizzesDBInteractions.GetLongAnswerSubmissionsByQuizID(userID, quizID)
	if !displayUserResult {
		for _, submissionItem := range mcqSubmissions {
			submissionItem.UserResult = 0
		}
		for _, submissionItem := range longAnswerSubmissions {
			submissionItem.UserResult = ""
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"mcqSubmissions":        mcqSubmissions,
		"longAnswerSubmissions": longAnswerSubmissions,
	})
}

// CreateMCQSubmission creates a new submission for an mcq question.
func CreateMCQSubmission(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	mcqID := utils.ConvertToUInt(c.FormValue("mcqID"))
	userResult := utils.ConvertToUInt(c.FormValue("userResult"))

	submission := quizzesModel.Submission{
		UserID:     userID,
		QuestionID: mcqID,
	}

	mcqSubmission := quizzesModel.MCQSubmission{
		Submission: submission,
		UserResult: userResult,
	}

	quizzesDBInteractions.CreateMCQSubmission(&mcqSubmission)

	return c.JSON(http.StatusOK, echo.Map{
		"message":       "MCQ submission created successfully",
		"mcqSubmission": mcqSubmission,
	})
}

// UpdateMCQSubmission updates a previous submission for an mcq question.
func UpdateMCQSubmission(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	mcqID := utils.ConvertToUInt(c.FormValue("mcqID"))
	userResult := utils.ConvertToUInt(c.FormValue("userResult"))

	mcqSubmission := quizzesDBInteractions.GetMCQSubmissionByQuestionID(mcqID, userID)
	mcqSubmission.UserResult = userResult

	quizzesDBInteractions.UpdateMCQSubmission(&mcqSubmission)

	return c.JSON(http.StatusOK, echo.Map{
		"message":       "MCQ submission updated successfully",
		"mcqSubmission": mcqSubmission,
	})
}
