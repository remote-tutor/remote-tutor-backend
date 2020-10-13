package organizations

import (
	authController "backend/controllers/auth"
	classUsersDBInteractions "backend/database/organizations"
	classUsersModel "backend/models/organizations"
	"github.com/labstack/echo"
	"net/http"
)

func GetClassesByUser(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	classes := classUsersDBInteractions.GetClassesByUser(userID)
	return c.JSON(http.StatusOK, echo.Map{
		"classes": classes,
	})
}

func Enroll(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	class := c.FormValue("selectedClass")
	classUser := classUsersModel.ClassUser{
		UserID:    userID,
		ClassHash: class,
	}
	err := classUsersDBInteractions.EnrollUser(&classUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (you've not been enrolled), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "You've enrolled successfully! - wait for admin verification -",
	})
}