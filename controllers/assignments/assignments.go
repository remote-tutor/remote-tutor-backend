package assignments

import (
	authController "backend/controllers/auth"
	filesUtils "backend/controllers/files"
	assignmentsFiles "backend/controllers/files/assignments"
	paginationController "backend/controllers/pagination"
	assignmentsDBInteractions "backend/database/assignments"
	usersDBInteractions "backend/database/users"
	assignmentsModel "backend/models/assignments"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

// GetAssignments gets the assignments of the logged in user OR selected year (IF ADMIN)
func GetAssignments(c echo.Context) error {
	admin := authController.FetchLoggedInUserAdminStatus(c)
	var year int
	if admin {
		year = utils.ConvertToInt(c.QueryParam("year"))
	} else {
		userID := authController.FetchLoggedInUserID(c)
		user := usersDBInteractions.GetUserByUserID(userID)
		year = user.Year
	}
	paginationData := paginationController.ExtractPaginationData(c)
	assignments, totalAssignments := assignmentsDBInteractions.GetAssignments(paginationData, year)
	return c.JSON(http.StatusOK, echo.Map{
		"assignments":      assignments,
		"totalAssignments": totalAssignments,
	})
}

func CreateOrUpdateAssignment(c echo.Context) error {
	method := c.Request().Method
	assignment := new(assignmentsModel.Assignment)
	if err := c.Bind(assignment); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Error reading assignment data from user",
		})
	}
	assignment.Deadline = utils.ConvertToTime(c.FormValue("deadline"))
	if method == "POST" {
		err := assignmentsDBInteractions.CreateAssignment(assignment)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Unexpected error occurred (assignment not created), please try again",
			})
		}
	}
	questionsFilePath, questionsErr := assignmentsFiles.UploadQuestionsFile(c, assignment)
	modelAnswerFilePath, modelAnswerErr := assignmentsFiles.UploadModelAnswerFile(c, assignment)
	if questionsErr != nil || modelAnswerErr != nil {
		if method == "POST" {
			err := assignmentsDBInteractions.DeleteAssignment(assignment)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "Unexpected error occurred (assignment not deleted), please try again",
				})
			}
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred when trying to upload the files. Please try again later",
		})
	}
	if questionsFilePath != "" {
		assignment.Questions = questionsFilePath
	}
	if modelAnswerFilePath != "" {
		assignment.ModelAnswer = modelAnswerFilePath
	}
	if method == "PUT" {
		assignment.CreatedAt = utils.ConvertToTime(c.FormValue("CreatedAt"))
	}
	err := assignmentsDBInteractions.UpdateAssignment(assignment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (assignment not created/updated), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Assignment Saved Successfully",
	})
}

// GetAssignmentByAssignmentID gets the assignment object by the passed ID
func GetAssignmentByAssignmentID(c echo.Context) error {
	assignmentID := utils.ConvertToUInt(c.QueryParam("assignmentID"))
	assignment := assignmentsDBInteractions.GetAssignmentByID(assignmentID)
	return c.JSON(http.StatusOK, echo.Map{
		"assignment": assignment,
	})
}

func GetQuestionsFile(c echo.Context) error {
	questionsPath := c.QueryParam("file")
	bytes, err := filesUtils.GetFile(questionsPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{})
	}
	c.Response().Header().Set("Content-Type", http.DetectContentType(bytes))
	return c.Blob(http.StatusOK, http.DetectContentType(bytes), bytes)
}
