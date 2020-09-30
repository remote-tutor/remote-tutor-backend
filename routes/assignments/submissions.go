package assignments

import (
	submissionsController "backend/controllers/assignments"
	"github.com/labstack/echo"
)

func InitializeSubmissionsRoutes(assignments *echo.Group, adminAssignments *echo.Group) {
	assignments.GET("/submission", submissionsController.GetSubmissionByUserAndAssignment)

	assignments.POST("/submissions", submissionsController.CreateOrUpdateSubmission)
	assignments.PUT("/submissions", submissionsController.CreateOrUpdateSubmission)
}
