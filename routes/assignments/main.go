package assignments

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	assignments := e.Group("/assignments", middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))))
	adminAssignments := adminRouter.Group("/assignments")

	InitializeAssignmentsRoutes(assignments, adminAssignments)
	InitializeSubmissionsRoutes(assignments, adminAssignments)
}
