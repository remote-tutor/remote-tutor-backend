package videos

import (
	authController "backend/controllers/auth"
	usersDBInteractions "backend/database/users"
	watchesDBInteractions "backend/database/videos"
	watchesModel "backend/models/videos"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func GetWatchByUserAndPart(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	partID := utils.ConvertToUInt(c.QueryParam("partID"))
	watch := watchesDBInteractions.GetUserWatchByUserAndPart(userID, partID)
	return c.JSON(http.StatusOK, echo.Map{
		"watch": watch,
	})
}

func CreateUserWatch(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	user := usersDBInteractions.GetUserByUserID(userID)
	if user.Admin {
		return c.JSON(http.StatusOK, echo.Map{})
	}
	partID := utils.ConvertToUInt(c.FormValue("videoPartID"))
	currentTime := time.Now()
	threeHoursDuration, _ := time.ParseDuration("3h")
	validTill := currentTime.Add(threeHoursDuration)
	userWatch := watchesModel.UserWatch{
		UserID: userID,
		VideoPartID: partID,
		StartAt: currentTime,
		ValidTill: validTill,
	}
	err := watchesDBInteractions.CreateUserWatch(&userWatch)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred, please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"watch": userWatch,
	})
}
