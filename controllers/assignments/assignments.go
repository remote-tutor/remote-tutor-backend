package assignments

import (
	authController "backend/controllers/auth"
	assignmentsDBInteractions "backend/database/assignments"
	usersDBInteractions "backend/database/users"
	assignmentsModel "backend/models/assignments"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

func GetAssignments(c echo.Context) error {
	admin := authController.FetchLoggedInUserAdminStatus(c)
	var year int
	if admin {
		year = utils.ConvertToInt(c.QueryParam("year"))
	} else {
		userID := authController.FetchLoggedInUserID(c)
		user := usersDBInteractions.GetUserByUserID(userID)
		year = user.Year
	}
	assignments, totalAssignments := assignmentsDBInteractions.GetAssignments(c, year)
	return c.JSON(http.StatusOK, echo.Map{
		"assignments": assignments,
		"totalAssignments": totalAssignments,
	})
}


func CreateAssignment(c echo.Context) error {
	assignment := new(assignmentsModel.Assignment)
	if err := c.Bind(&assignment); err != nil {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Error reading assignment data from user",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{})
}
