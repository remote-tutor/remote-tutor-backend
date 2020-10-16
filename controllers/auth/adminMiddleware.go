package controllers

import (
	classUsersDBInteractions "backend/database/organizations"
	"net/http"

	"github.com/labstack/echo"
)

// CheckAdmin checks if the request is comming from an admin user
func CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := FetchLoggedInUserID(c)
		class := c.QueryParam("selectedClass")
		classUser := classUsersDBInteractions.GetClassUserByUserIDAndClass(userID, class)
		if classUser.Admin {
			return next(c)
		}
		return c.JSON(http.StatusForbidden, echo.Map{
			"message": "Unauthorized admin access",
		})
	}
}
