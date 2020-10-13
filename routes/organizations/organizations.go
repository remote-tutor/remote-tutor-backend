package organizations

import (
	userClassesController "backend/controllers/organizations"
	"github.com/labstack/echo"
)

func InitializeOrganizationsRoutes(classes *echo.Group, adminClasses *echo.Group) {
	adminClasses.GET("/admins", userClassesController.GetOrganizationAdminsByClass)
}

