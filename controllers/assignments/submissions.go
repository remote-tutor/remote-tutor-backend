package assignments

import (
	authController "backend/controllers/auth"
	submissionsFiles "backend/controllers/files/assignments"
	paginationController "backend/controllers/pagination"
	submissionsDBInteractions "backend/database/assignments"
	classesDBInteractions "backend/database/organizations"
	submissionsModel "backend/models/assignments"
	assignmentsPDFHandler "backend/pdf/handlers/assignments"
	"backend/utils"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

func GetSubmissionByUserAndAssignment(c echo.Context) error {
	assignmentHash := c.QueryParam("assignmentHash")
	userID := authController.FetchLoggedInUserID(c)
	submission := submissionsDBInteractions.GetSubmissionByUserAndAssignment(userID, assignmentHash)
	return c.JSON(http.StatusOK, echo.Map{
		"submission": submission,
	})
}

func CreateOrUpdateSubmission(c echo.Context) error {
	method := c.Request().Method
	submission := new(submissionsModel.AssignmentSubmission)
	submission.UserID = authController.FetchLoggedInUserID(c)
	submission.AssignmentID = utils.ConvertToUInt(c.FormValue("assignmentID"))
	submission.UploadedAt = time.Now()
	assignmentHash := c.FormValue("assignmentHash")
	if method == "POST" {
		err := submissionsDBInteractions.CreateSubmission(submission)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Unexpected error occurred (submission not created), please try again",
			})
		}
	} else {
		originalSubmission := submissionsDBInteractions.GetSubmissionByUserAndAssignment(submission.UserID, assignmentHash)
		if originalSubmission.Graded {
			return c.JSON(http.StatusForbidden, echo.Map{
				"message": "Sorry this assignment has been graded, you can't change the submission",
			})
		}
	}
	class := submissionsDBInteractions.GetClassByAssignmentHash(assignmentHash)
	submissionFilePath, submissionErr := submissionsFiles.
		UploadUserSubmissionFile(c, submission.UserID, assignmentHash, &class)
	if submissionErr != nil {
		if method == "POST" {
			err := submissionsDBInteractions.DeleteSubmission(submission)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Unexpected error occurred (submission not deleted), please try again",
				})
			}
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred when trying to upload the submission. Please try again later",
		})
	}
	if submissionFilePath != "" {
		submission.File = submissionFilePath
		submission.UploadedAt = time.Now()
	}
	err := submissionsDBInteractions.UpdateSubmission(submission)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (submission not created/updated), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "AssignmentSubmission Saved Successfully",
	})
}

func GetSubmissionsByAssignmentForAllUsers(c echo.Context) error {
	assignmentID := utils.ConvertToUInt(c.QueryParam("assignmentID"))
	fullNameSearch := c.QueryParam("searchBy")
	paginationData := paginationController.ExtractPaginationData(c)
	submissions, totalSubmissions := submissionsDBInteractions.GetSubmissionsByAssignmentForAllUsers(paginationData, assignmentID, fullNameSearch)
	return c.JSON(http.StatusOK, echo.Map{
		"submissions": submissions,
		"totalSubmissions": totalSubmissions,
	})
}

func UpdateSubmissionByAdmin(c echo.Context) error {
	userID := utils.ConvertToUInt(c.FormValue("userID"))
	assignmentHash := c.FormValue("assignmentHash")
	submission := submissionsDBInteractions.GetSubmissionByUserAndAssignment(userID, assignmentHash)
	submission.Graded = true
	submission.Mark = utils.ConvertToInt(c.FormValue("mark"))
	submission.Feedback = c.FormValue("feedback")
	err := submissionsDBInteractions.UpdateSubmission(&submission)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (submission not updated/marked), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Submission Updated Successfully",
		"submission": submission,
	})
}

func GetSubmissionsPDFByAssignment(c echo.Context) error {
	assignmentID := utils.ConvertToUInt(c.QueryParam("assignmentID"))
	classHash := c.QueryParam("selectedClass")
	class := classesDBInteractions.GetClassByHash(classHash)
	paginationData := paginationController.ExtractPaginationData(c)
	submissions, _ := submissionsDBInteractions.GetSubmissionsByAssignmentForAllUsers(paginationData, assignmentID, "")
	assignment := submissionsDBInteractions.GetAssignmentByID(assignmentID)
	pdfGenerator, err := assignmentsPDFHandler.DeliverAssignmentsPDF(&assignment, submissions, &class)
	if err != nil {
		log.Fatal(err)
		//return c.JSON(http.StatusInternalServerError, echo.Map{})
	}
	c.Response().Header().Set("Content-Type", "application/pdf")
	return c.Blob(http.StatusOK, "application/pdf", pdfGenerator.Bytes())
}