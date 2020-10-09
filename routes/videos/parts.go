package videos

import (
	partsController "backend/controllers/videos"
	"github.com/labstack/echo")

func InitializePartsRoutes(videos *echo.Group, adminVideos *echo.Group) {
	videos.GET("/parts", partsController.GetPartsByVideo)
	videos.GET("/part", partsController.GetPartLink)

	adminVideos.POST("/parts", partsController.CreatePart)
	adminVideos.PUT("/parts", partsController.UpdatePart)
	adminVideos.DELETE("/parts", partsController.DeletePart)
}
