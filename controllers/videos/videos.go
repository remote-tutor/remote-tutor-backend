package videos

import (
	authController "backend/controllers/auth"
	partsFiles "backend/controllers/files/videos"
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

func GetVideoByID(c echo.Context) error {
	videoID := utils.ConvertToUInt(c.QueryParam("id"))
	video := videosDBInterations.GetVideoByID(videoID)
	return c.JSON(http.StatusOK, echo.Map{
		"video": video,
	})
}

func CreateVideo(c echo.Context) error {
	availableFrom := utils.ConvertToTime(c.FormValue("availableFrom"))
	year := utils.ConvertToInt(c.FormValue("year"))
	title := c.FormValue("title")
	video := videosModel.Video{
		AvailableFrom: now.With(availableFrom).BeginningOfDay(),
		Year: year,
		Title: title,
	}
	err := videosDBInterations.CreateVideo(&video)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (video not created), please try again",
		})
	}
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
	video.Title = c.FormValue("title")

	err := videosDBInterations.UpdateVideo(&video)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (video not updated), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"video": video,
		"message": "Video Updated Successfully",
	})
}

func DeleteVideo(c echo.Context) error {
	id := utils.ConvertToUInt(c.FormValue("id"))
	video := videosDBInterations.GetVideoByID(id)
	typedTitle := c.FormValue("typedTitle")
	if video.Title != typedTitle {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Sorry, you've entered a wrong video title, please check your selection and try again",
		})
	}
	parts := videosDBInterations.GetPartsByVideo(video.ID)
	err := partsFiles.DeleteVideo(&video, parts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred while trying to delete the video files, please try again later",
		})
	}
	err = videosDBInterations.DeleteVideo(&video)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (video not deleted), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "The video and its parts are deleted successfully",
	})
}