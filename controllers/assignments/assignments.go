package assignments

import (
	"backend/aws"
	assignmentsFiles "backend/controllers/files/assignments"
	paginationController "backend/controllers/pagination"
	assignmentsDBInteractions "backend/database/assignments"
	classesDBInteractions "backend/database/organizations"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

// GetAssignmentsByClass gets the assignments of the logged in user OR selected year (IF ADMIN)
func GetAssignmentsByClass(c echo.Context) error {
	class := c.QueryParam("selectedClass")
	paginationData := paginationController.ExtractPaginationData(c)
	assignments, totalAssignments := assignmentsDBInteractions.GetAssignmentsByClass(paginationData, class)
	return c.JSON(http.StatusOK, echo.Map{
		"assignments":      assignments,
		"totalAssignments": totalAssignments,
	})
}

func CreateOrUpdateAssignment(c echo.Context) error {
	method := c.Request().Method
	id := utils.ConvertToUInt(c.FormValue("id"))
	assignment := assignmentsDBInteractions.GetAssignmentByID(id)
	assignment.Title = c.FormValue("title")
	assignment.TotalMark = utils.ConvertToInt(c.FormValue("totalMark"))
	assignment.Deadline = utils.ConvertToTime(c.FormValue("deadline"))
	if method == "POST" {
		assignment.ClassHash = c.FormValue("selectedClass")
		err := assignmentsDBInteractions.CreateAssignment(&assignment)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Unexpected error occurred (assignment not created), please try again",
			})
		}
	}
	class := classesDBInteractions.GetClassByHash(assignment.ClassHash)
	questionsFilePath, questionsErr := assignmentsFiles.UploadQuestionsFile(c, &assignment, &class)
	modelAnswerFilePath, modelAnswerErr := assignmentsFiles.UploadModelAnswerFile(c, &assignment, &class)
	if questionsErr != nil || modelAnswerErr != nil {
		if method == "POST" {
			err := assignmentsDBInteractions.DeleteAssignment(&assignment)
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
	err := assignmentsDBInteractions.UpdateAssignment(&assignment)
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

// GetAssignmentByAssignmentHash gets the assignment object by the passed hash
func GetAssignmentByAssignmentHash(c echo.Context) error {
	assignmentHash := c.FormValue("assignmentHash")
	assignment := assignmentsDBInteractions.GetAssignmentByHash(assignmentHash)
	return c.JSON(http.StatusOK, echo.Map{
		"assignment": assignment,
	})
}

func GetQuestionsFile(c echo.Context) error {
	originalUrl := c.QueryParam("originalUrl")
	url, err := aws.GenerateSignedURL(originalUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred when trying to get the link, please try again latter",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"url": url,
	})
}
