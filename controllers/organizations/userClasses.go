package organizations

import (
	authController "backend/controllers/auth"
	paginationController "backend/controllers/pagination"
	classUsersDBInteractions "backend/database/organizations"
	classUsersModel "backend/models/organizations"
	"backend/utils"
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

func GetStudentsByClass(c echo.Context) error {
	searchByValue := c.QueryParam("searchByValue")
	searchByField := c.QueryParam("searchByField")
	pending := utils.ConvertToBool(c.QueryParam("pending"))
	class := c.QueryParam("selectedClass")

	paginationData := paginationController.ExtractPaginationData(c)
	users, totalUsers := classUsersDBInteractions.GetStudentsByClass(paginationData, searchByValue, searchByField, class, pending)

	return c.JSON(http.StatusOK, echo.Map{
		"students":      users,
		"totalStudents": totalUsers,
	})
}

func AcceptStudents(c echo.Context) error {
	id := utils.ConvertToUInt(c.FormValue("classUserID"))
	status := utils.ConvertToBool(c.FormValue("status"))
	classUser := classUsersDBInteractions.GetClassUserByID(id)
	if !status {
		err := classUsersDBInteractions.DeleteClassUser(&classUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Unexpected error occurred (user not deleted), please try again",
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "User rejected successfully",
		})
	}
	classUser.Activated = true
	err := classUsersDBInteractions.UpdateClassUser(&classUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (user not saved), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User saved successfully",
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