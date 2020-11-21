package videos

import (
	authController "backend/controllers/auth"
	classUsersDBInteractions "backend/database/organizations"
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
	class := c.QueryParam("selectedClass")
	classUser := classUsersDBInteractions.GetClassUserByUserIDAndClass(userID, class)
	if classUser.Admin {
		return c.JSON(http.StatusOK, echo.Map{})
	}
	partID := utils.ConvertToUInt(c.FormValue("videoPartID"))
	currentTime := time.Now()
	video := watchesDBInteractions.GetVideoByPartID(partID)
	validTill := getSmallestDate(video.AvailableTo, time.Now().Add(time.Duration(video.StudentHours) * time.Hour))
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

func getSmallestDate(first, second time.Time) time.Time {
	if first.Before(second) {
		return first
	}
	return second
}