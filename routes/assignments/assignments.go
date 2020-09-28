package assignments

import (
	assignmentsController "backend/controllers/assignments"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	assignments := e.Group("/assignments", middleware.JWT([]byte("secret")))
	assignments.GET("", assignmentsController.GetAssignments)
	assignments.GET("/assignment", assignmentsController.GetAssignmentByAssignmentID)
	assignments.GET("/assignment/file", assignmentsController.GetQuestionsFile)

	adminAssignments := adminRouter.Group("/assignments")
	adminAssignments.POST("", assignmentsController.CreateAssignment)
}
