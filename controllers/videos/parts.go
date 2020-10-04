package videos

import (
	partsFiles "backend/controllers/files/videos"
	partsDBInteractions "backend/database/videos"
	partsModel "backend/models/videos"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

func CreatePart(c echo.Context) error {
	videoID := utils.ConvertToUInt(c.FormValue("videoID"))
	number := utils.ConvertToInt(c.FormValue("number"))
	fileLocation, err := partsFiles.UploadVideoPart(c, videoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Sorry, we're unable to upload the part right now, please try again later",
		})
	}
	part := partsModel.VideoPart{
		VideoID: videoID,
		Link: fileLocation,
		Number: number,
	}
	partsDBInteractions.CreatePart(&part)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "VideoPart Uploaded Successfully",
	})
}
