package assignments

import (
	assignmentsController "backend/controllers/assignments"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	assignments := e.Group("/assignments", middleware.JWT([]byte("secret")))
	assignments.GET("", assignmentsController.GetAssignments)

	adminAssignments := adminRouter.Group("/assignments")
	adminAssignments.POST("", assignmentsController.CreateAssignment)
}
