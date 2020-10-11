package organizations

import (
	organizationsController "backend/controllers/organizations"
	"github.com/labstack/echo"
)

func InitializeUerClassesRoutes(classes *echo.Group, adminClasses *echo.Group) {
	classes.GET("", organizationsController.GetClassesByUser)
}
