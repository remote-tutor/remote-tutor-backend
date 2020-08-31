package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// CheckAdmin checks if the request is comming from an admin user
func CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		admin := FetchLoggedInUserAdminStatus(c)
		if admin {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Unauthorized admin access",
		})
	}
}
