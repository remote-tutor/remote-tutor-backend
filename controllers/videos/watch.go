package videos

import (
	authController "backend/controllers/auth"
	"backend/controllers/pagination"
	classUsersDBInteractions "backend/database/organizations"
	watchesDBInteractions "backend/database/videos"
	watchesModel "backend/models/videos"
	watchesPDFHandler "backend/pdf/handlers/videos"
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

func GetPartWatchesForAllUsers(c echo.Context) error {
	partID := utils.ConvertToUInt(c.QueryParam("partID"))
	paginationData := pagination.ExtractPaginationData(c)
	watches, total := watchesDBInteractions.GetPartWatchesForAllUsers(partID, paginationData)
	return c.JSON(http.StatusOK, echo.Map{
		"watches": watches,
		"total":   total,
	})
}

func GetPartWatchesPDF(c echo.Context) error {
	partID := utils.ConvertToUInt(c.QueryParam("partID"))
	part := watchesDBInteractions.GetPartByID(partID)
	parts := watchesDBInteractions.GetPartsByVideo(part.VideoID)
	paginationData := pagination.ExtractPaginationData(c)
	watches, _ := watchesDBInteractions.GetPartWatchesForAllUsers(partID, paginationData)
	pdfGenerator, err := watchesPDFHandler.DeliverWatchesPDF(&part, parts, watches)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{})
	}
	c.Response().Header().Set("Content-Type", "application/pdf")
	return c.Blob(http.StatusOK, "application/pdf", pdfGenerator.Bytes())
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
	validTill := time.Now().Add(time.Duration(video.StudentHours) * time.Hour)
	userWatch := watchesModel.UserWatch{
		UserID:      userID,
		VideoPartID: partID,
		StartAt:     currentTime,
		ValidTill:   validTill,
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

func DeleteUserWatch(c echo.Context) error {
	partID := utils.ConvertToUInt(c.QueryParam("partID"))
	userToDelete := utils.ConvertToUInt(c.QueryParam("userID"))
	userWatchToDelete := watchesModel.UserWatch{
		UserID:      userToDelete,
		VideoPartID: partID,
	}
	err := watchesDBInteractions.DeleteUserWatch(&userWatchToDelete)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred, please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{})
}
