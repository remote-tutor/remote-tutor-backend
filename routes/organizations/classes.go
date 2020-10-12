package organizations

import (
	classesController "backend/controllers/organizations"
	"github.com/labstack/echo"
)

func InitializeClassesRoutes(classes *echo.Group, adminClasses *echo.Group) {
	classes.GET("/all", classesController.GetAllClasses)
}
