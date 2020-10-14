package organizations

import (
	organizationAdminsController "backend/controllers/organizations"
	"github.com/labstack/echo"
)

func InitializeOrganizationAdminsRoutes(classes *echo.Group, adminClasses *echo.Group) {
	adminClasses.GET("/admins", organizationAdminsController.GetOrganizationAdminsByClass)
	adminClasses.POST("/admins", organizationAdminsController.AddAdminToOrganization)
}
