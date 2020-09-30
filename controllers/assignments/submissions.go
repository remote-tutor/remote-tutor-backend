package assignments

import (
	authController "backend/controllers/auth"
	submissionsFiles "backend/controllers/files/assignments"
	submissionsDBInteractions "backend/database/assignments"
	submissionsModel "backend/models/assignments"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func GetSubmissionByID(c echo.Context) error {
	submissionID := utils.ConvertToUInt(c.QueryParam("submissionID"))
	submission := submissionsDBInteractions.GetSubmissionByID(submissionID)
	return c.JSON(http.StatusOK, echo.Map{
		"submission": submission,
	})
}

func CreateOrUpdateSubmission(c echo.Context) error {
	method := c.Request().Method
	submission := new(submissionsModel.AssignmentSubmission)
	if err := c.Bind(submission); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Error reading assignment data from user",
		})
	}
	submission.UserID = authController.FetchLoggedInUserID(c)
	if method == "POST" {
		submission.UploadedAt = time.Now()
		submissionsDBInteractions.CreateSubmission(submission)
	}
	submissionFilePath, submissionErr := submissionsFiles.UploadUserSubmissionFile(c, submission)
	if submissionErr != nil {
		if method == "POST" {
			submissionsDBInteractions.DeleteSubmission(submission)
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred when trying to upload the submission. Please try again later",
		})
	}
	if submissionFilePath != "" {
		submission.File = submissionFilePath
		submission.UploadedAt = time.Now()
	}
	submissionsDBInteractions.UpdateSubmission(submission)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "AssignmentSubmission Saved Successfully",
	})
}