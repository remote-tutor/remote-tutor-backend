package videos

import (
	partsFiles "backend/controllers/files/videos"
	paginationController "backend/controllers/pagination"
	classesDBInteractions "backend/database/organizations"
	videosDBInterations "backend/database/videos"
	videosModel "backend/models/videos"
	"backend/utils"
	"github.com/jinzhu/now"
	"github.com/labstack/echo"
	"net/http"
)

func GetVideosByClassAndMonthAndYear(c echo.Context) error {
	class := c.QueryParam("selectedClass")
	date := utils.ConvertToTime(c.QueryParam("date"))
	endOfMonth := now.With(date).EndOfMonth()
	startOfMonth := now.With(date).BeginningOfMonth()
	videos := videosDBInterations.GetVideosByClassAndMonthAndYear(class, startOfMonth, endOfMonth)
	return c.JSON(http.StatusOK, echo.Map{
		"videos": videos,
	})
}

func GetVideoByHash(c echo.Context) error {
	videoHash := c.QueryParam("videoHash")
	video := videosDBInterations.GetVideoByHash(videoHash)
	return c.JSON(http.StatusOK, echo.Map{
		"video": video,
	})
}

func GetNonAccessedStudents(c echo.Context) error {
	videoID := utils.ConvertToUInt(c.QueryParam("videoID"))
	paginationData := paginationController.ExtractPaginationData(c)
	classHash := c.QueryParam("selectedClass")
	search := c.QueryParam("searchByValue")
	students, numberOfStudents := videosDBInterations.GetNonAccessedStudents(paginationData, classHash, search, videoID)
	return c.JSON(http.StatusOK, echo.Map{
		"students": students,
		"total": numberOfStudents,
	})
}

func CreateVideo(c echo.Context) error {
	availableFrom := utils.ConvertToTime(c.FormValue("availableFrom"))
	class := c.FormValue("selectedClass")
	title := c.FormValue("title")
	studentHours := utils.ConvertToUInt(c.FormValue("studentTime"))
	if studentHours <= 0 {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Make sure that you've entered a correct time for each part",
		})
	}
	video := videosModel.Video{
		AvailableFrom: now.With(availableFrom).BeginningOfDay(),
		ClassHash: class,
		Title: title,
		StudentHours: studentHours,
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
	availableFrom := utils.ConvertToTime(c.FormValue("availableFrom"))
	video.AvailableFrom = now.With(availableFrom).BeginningOfDay()
	video.Title = c.FormValue("title")
	studentHours := utils.ConvertToUInt(c.FormValue("studentTime"))
	if studentHours <= 0 {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Make sure that you've entered a correct time for each part",
		})
	}
	video.StudentHours = studentHours

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
	class := classesDBInteractions.GetClassByHash(video.ClassHash)
	err := partsFiles.DeleteVideo(&video, parts, &class)
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