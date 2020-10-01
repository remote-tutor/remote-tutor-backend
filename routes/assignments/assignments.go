package assignments

import (
	assignmentsController "backend/controllers/assignments"
	"github.com/labstack/echo"
)

func InitializeAssignmentsRoutes(assignments *echo.Group, adminAssignments *echo.Group) {
	assignments.GET("", assignmentsController.GetAssignments)
	assignments.GET("/assignment", assignmentsController.GetAssignmentByAssignmentID)
	assignments.GET("/assignment/file", assignmentsController.GetQuestionsFile)

	adminAssignments.POST("", assignmentsController.CreateOrUpdateAssignment)
	adminAssignments.PUT("", assignmentsController.CreateOrUpdateAssignment)
}
