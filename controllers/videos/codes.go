package videos

import (
	authController "backend/controllers/auth"
	paginationController "backend/controllers/pagination"
	codesDBInteractions "backend/database/videos"
	codesModel "backend/models/videos"
	"backend/utils"
	"crypto/rand"
	"encoding/hex"
	"github.com/labstack/echo"
	"net/http"
)

func GetCodesByVideo(c echo.Context) error {
	videoID := utils.ConvertToUInt(c.QueryParam("videoID"))
	paginationData := paginationController.ExtractPaginationData(c)
	search := c.QueryParam("search")
	codes, numberOfCodes := codesDBInteractions.GetCodesByVideo(paginationData, search, videoID)
	return c.JSON(http.StatusOK, echo.Map{
		"codes": codes,
		"total": numberOfCodes,
	})
}

func GenerateCodes(c echo.Context) error {
	videoID := utils.ConvertToUInt(c.FormValue("videoID"))
	createdByUserID := authController.FetchLoggedInUserID(c)
	numberOfCodes := utils.ConvertToUInt(c.FormValue("numberOfCodes"))
	codes := make([]codesModel.Code, numberOfCodes)
	for i := uint(0); i < numberOfCodes; i++ {
		bytes := make([]byte, 6)
		rand.Read(bytes)
		codes[i] = codesModel.Code{
			Value:           hex.EncodeToString(bytes),
			VideoID:         videoID,
			CreatedByUserID: createdByUserID,
		}
	}
	err := codesDBInteractions.GenerateCodes(codes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (codes not generated)",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Codes generated successfully",
	})
}
