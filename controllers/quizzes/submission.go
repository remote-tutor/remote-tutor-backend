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

	// displayUserResult := utils.ConvertToBool(c.QueryParam("displayUserResult"))
	mcqSubmissions := quizzesDBInteractions.GetMCQSubmissionsByQuizID(userID, quizID)
	longAnswerSubmissions := quizzesDBInteractions.GetLongAnswerSubmissionsByQuizID(userID, quizID)

	return c.JSON(http.StatusOK, echo.Map{
		"mcqSubmissions":        mcqSubmissions,
		"longAnswerSubmissions": longAnswerSubmissions,
	})
}

// CreateMCQSubmission creates a new submission for an mcq question.
func CreateMCQSubmission(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	mcqID := utils.ConvertToUInt(c.FormValue("mcqID"))
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))
	userResult := utils.ConvertToUInt(c.FormValue("userResult"))

	submission := quizzesModel.Submission{
		UserID: userID,
	}

	mcqSubmission := quizzesModel.MCQSubmission{
		Submission: submission,
		UserResult: userResult,
		MCQID:      mcqID,
	}

	quizzesDBInteractions.CreateMCQSubmission(&mcqSubmission)

	mcqQuestion := quizzesDBInteractions.GetMCQByID(mcqID)
	if userResult == mcqQuestion.CorrectAnswer {
		quizGrade := quizzesDBInteractions.GetGradesByQuizID(userID, quizID)
		quizGrade.Grade += mcqQuestion.TotalMark
		quizzesDBInteractions.UpdateGrade(&quizGrade)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":       "MCQ submission created successfully",
		"mcqSubmission": mcqSubmission,
	})
}

// UpdateMCQSubmission updates a previous submission for an mcq question.
func UpdateMCQSubmission(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	mcqID := utils.ConvertToUInt(c.FormValue("mcqID"))
	quizID := utils.ConvertToUInt(c.FormValue("quizID"))
	userResult := utils.ConvertToUInt(c.FormValue("userResult"))

	mcqSubmission := quizzesDBInteractions.GetMCQSubmissionByQuestionID(userID, mcqID)
	previousResult := mcqSubmission.UserResult
	if userResult != previousResult {
		mcq := quizzesDBInteractions.GetMCQByID(mcqID)
		quizGrade := quizzesDBInteractions.GetGradesByQuizID(userID, quizID)
		if previousResult == mcq.CorrectAnswer {
			quizGrade.Grade -= mcq.TotalMark
		} else if userResult == mcq.CorrectAnswer {
			quizGrade.Grade += mcq.TotalMark
		}
		mcqSubmission.UserResult = userResult
		quizzesDBInteractions.UpdateMCQSubmission(&mcqSubmission)
		quizzesDBInteractions.UpdateGrade(&quizGrade)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message":       "MCQ submission updated successfully",
		"mcqSubmission": mcqSubmission,
	})
}
