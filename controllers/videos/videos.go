package videos

import (
	videosDBInterations "backend/database/videos"
	videosModel "backend/models/videos"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

func CreateVideo(c echo.Context) error {
	availableFrom := utils.ConvertToTime(c.FormValue("availableFrom"))
	video := videosModel.Video{
		AvailableFrom: availableFrom,
	}
	videosDBInterations.CreateVideo(&video)
	return c.JSON(http.StatusOK, echo.Map{
		"video": video,
	})
}
