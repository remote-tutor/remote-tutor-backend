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

func GetCodeByUserAndVideo(c echo.Context) error {
	videoID := utils.ConvertToUInt(c.QueryParam("videoID"))
	userID := authController.FetchLoggedInUserID(c)
	code := codesDBInteractions.GetCodeByUserAndVideo(userID, videoID)
	if code.Value == "" {
		return c.JSON(http.StatusOK, echo.Map{
			"status": false,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": true,
	})
}

func GrantStudentAccess(c echo.Context) error {
	userID := authController.FetchLoggedInUserID(c)
	videoID := utils.ConvertToUInt(c.FormValue("videoID"))
	codeString := c.FormValue("code")
	code := codesDBInteractions.GetCodeByValueAndVideo(codeString, videoID)
	if code.Value == "" {
		return c.JSON(http.StatusNotAcceptable, echo.Map{
			"message": "Make sure you've entered the correct code",
		})
	}
	if code.UsedByUserID != 0 {
		return c.JSON(http.StatusConflict, echo.Map{
			"message": "This code has been used before",
		})
	}
	code.UsedByUserID = userID
	err := codesDBInteractions.UpdateCode(&code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred, please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Code used successfully",
	})
}

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

func AddManualAccess(c echo.Context) error {
	videoID := utils.ConvertToUInt(c.FormValue("videoID"))
	usersToGiveAccess := utils.ConvertToUIntArray(utils.ConvertToFormArray(c.FormValue("addedTo[]")))
	codes := make([]codesModel.Code, len(usersToGiveAccess))
	for index, userID := range usersToGiveAccess {
		bytes := make([]byte, 6)
		rand.Read(bytes)
		codes[index] = codesModel.Code{
			Value:           hex.EncodeToString(bytes),
			VideoID:         videoID,
			CreatedByUserID: authController.FetchLoggedInUserID(c),
			UsedByUserID:    userID,
			Manual:          true,
		}
	}
	err := codesDBInteractions.CreateManualAccess(codes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (access may not be fully given)",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Video access has been updated successfully",
	})
}
