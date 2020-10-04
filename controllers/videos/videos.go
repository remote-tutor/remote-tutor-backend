package videos

import (
	authController "backend/controllers/auth"
	videosDBInterations "backend/database/videos"
	videosModel "backend/models/videos"
	"backend/utils"
	"github.com/jinzhu/now"
	"github.com/labstack/echo"
	"net/http"
)

func GetVideosByMonthAndYear(c echo.Context) error {
	isAdmin := authController.FetchLoggedInUserAdminStatus(c)
	var year int
	if isAdmin {
		year = utils.ConvertToInt(c.QueryParam("year"))
	} else {
		year = authController.FetchLoggedInUserYear(c)
	}
	date := utils.ConvertToTime(c.QueryParam("date"))
	endOfMonth := now.With(date).EndOfMonth()
	startOfMonth := now.With(date).BeginningOfMonth()
	videos := videosDBInterations.GetVideosByMonthAndYear(year, startOfMonth, endOfMonth)
	return c.JSON(http.StatusOK, echo.Map{
		"videos": videos,
	})
}

func CreateVideo(c echo.Context) error {
	availableFrom := utils.ConvertToTime(c.FormValue("availableFrom"))
	year := utils.ConvertToInt(c.FormValue("year"))
	video := videosModel.Video{
		AvailableFrom: now.With(availableFrom).BeginningOfDay(),
		Year: year,
	}
	videosDBInterations.CreateVideo(&video)
	return c.JSON(http.StatusOK, echo.Map{
		"video": video,
		"message": "Video Created Successfully",
	})
}

func UpdateVideo(c echo.Context) error {
	id := utils.ConvertToUInt(c.FormValue("id"))
	video := videosDBInterations.GetVideoByID(id)
	video.AvailableFrom = utils.ConvertToTime(c.FormValue("availableFrom"))
	video.Year = utils.ConvertToInt(c.FormValue("year"))

	videosDBInterations.UpdateVideo(&video)
	return c.JSON(http.StatusOK, echo.Map{
		"video": video,
		"message": "Video Updated Successfully",
	})
}