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
	pending := utils.ConvertToBool(c.QueryParam("pending"))
	class := c.QueryParam("selectedClass")

	paginationData := paginationController.ExtractPaginationData(c)
	users, totalUsers := classUsersDBInteractions.GetStudentsByClass(paginationData, searchByValue, class, pending)

	return c.JSON(http.StatusOK, echo.Map{
		"students":      users,
		"totalStudents": totalUsers,
	})
}

func AcceptStudents(c echo.Context) error {
	id := utils.ConvertToUInt(c.FormValue("classUserID"))
	status := utils.ConvertToBool(c.FormValue("status"))
	admin := utils.ConvertToBool(c.FormValue("admin"))
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
	if admin {
		classUser.Admin = true
	} else {
		classUser.Admin = false
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
	class := c.FormValue("newSelectedClass")
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

func CheckIfUserEnrolledToClass(c echo.Context) error {
	class := c.QueryParam("selectedClass")
	userID := utils.ConvertToUInt(c.QueryParam("userID"))
	classUser := classUsersDBInteractions.GetClassUserByUserIDAndClass(userID, class)
	var enrolled bool
	if classUser.ID == 0 {
		enrolled = false
	} else {
		enrolled = true
	}
	return c.JSON(http.StatusOK, echo.Map{
		"classUser": classUser,
		"enrolled":  enrolled,
	})
}