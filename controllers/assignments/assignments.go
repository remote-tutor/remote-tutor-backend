package assignments

import (
	authController "backend/controllers/auth"
	assignmentsFiles "backend/controllers/files/assignments"
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
	assignments, totalAssignments := assignmentsDBInteractions.GetAssignments(c, year)
	return c.JSON(http.StatusOK, echo.Map{
		"assignments":      assignments,
		"totalAssignments": totalAssignments,
	})
}

// CreateAssignment creates a new assignment with the given data
func CreateAssignment(c echo.Context) error {
	assignment := new(assignmentsModel.Assignment)
	if err := c.Bind(assignment); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Error reading assignment data from user",
		})
	}
	assignment.Deadline = utils.ConvertToTime(c.FormValue("deadline"))
	assignmentsDBInteractions.CreateAssignment(assignment)
	questionsFilePath, questionsErr := assignmentsFiles.UploadQuestionsFile(c, assignment)
	modelAnswerFilePath, modelAnswerErr := assignmentsFiles.UploadModelAnswerFile(c, assignment)
	if questionsErr != nil || modelAnswerErr != nil {
		assignmentsDBInteractions.DeleteAssignment(assignment)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred when trying to upload the files. Please try again later",
		})
	}
	assignment.Questions = questionsFilePath
	assignment.ModelAnswer = modelAnswerFilePath
	assignmentsDBInteractions.UpdateAssignment(assignment)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Assignment Created Successfully",
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
	bytes, err := assignmentsFiles.GetFile(questionsPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Sorry we cannot download the requested file now, please try again later",
		})
	}
	c.Response().Header().Set("Content-Type", http.DetectContentType(bytes))
	return c.Blob(http.StatusOK, http.DetectContentType(bytes), bytes)
}
