package organizations

import (
	organizationsDBInteractions "backend/database/organizations"
	"github.com/labstack/echo"
	"net/http"
)

func GetOrganizationAdminsByClass(c echo.Context) error {
	class := c.QueryParam("selectedClass")
	admins := organizationsDBInteractions.GetOrganizationAdminsByClass(class)
	return c.JSON(http.StatusOK, echo.Map{
		"admins": admins,
	})
}

