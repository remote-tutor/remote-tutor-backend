package organizations

import (
	authController "backend/controllers/auth"
	organizationsDBInteractions "backend/database/organizations"
	"github.com/labstack/echo"
	"net/http"
)

func GetClassesByUser(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	classes := organizationsDBInteractions.GetClassesByUser(userID)
	return c.JSON(http.StatusOK, echo.Map{
		"classes": classes,
	})
}
