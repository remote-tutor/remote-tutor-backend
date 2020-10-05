package videos

import (
	partsFiles "backend/controllers/files/videos"
	partsDBInteractions "backend/database/videos"
	partsModel "backend/models/videos"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

func GetPartsByVideo(c echo.Context) error {
	videoID := utils.ConvertToUInt(c.QueryParam("videoID"))
	parts := partsDBInteractions.GetPartsByVideo(videoID)
	return c.JSON(http.StatusOK, echo.Map{
		"parts": parts,
	})
}

func CreatePart(c echo.Context) error {
	videoID := utils.ConvertToUInt(c.FormValue("videoID"))
	number := utils.ConvertToInt(c.FormValue("number"))
	fileName := c.FormValue("fileName")
	fileLocation, err := partsFiles.UploadVideoPart(c, videoID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Sorry, we're unable to upload the part right now, please try again later",
		})
	}
	part := partsModel.VideoPart{
		VideoID: videoID,
		Link:    fileLocation,
		Number:  number,
		Name:    fileName,
	}
	partsDBInteractions.CreatePart(&part)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Video Part Uploaded Successfully",
	})
}

func UpdatePart(c echo.Context) error {
	id := utils.ConvertToUInt(c.FormValue("id"))
	part := partsDBInteractions.GetPartByID(id)
	part.Number = utils.ConvertToInt(c.FormValue("number"))
	partsDBInteractions.UpdatePart(&part)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Video Part Updated Successfully",
	})
}

func DeletePart(c echo.Context) error {
	id := utils.ConvertToUInt(c.FormValue("id"))
	typedName := c.FormValue("typedName")
	part := partsDBInteractions.GetPartByID(id)
	if typedName != part.Name {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Sorry, you've entered a wrong filename, please check your selection and try again",
		})
	}
	err := partsFiles.DeleteVideoPart(&part)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred while trying to delete the video part, please try again later",
		})
	}
	partsDBInteractions.DeletePart(&part)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Vide Part Deleted Successfully",
	})
}