package videos

import (
	partsController "backend/controllers/videos"
	"github.com/labstack/echo")

func InitializePartsRoutes(videos *echo.Group, adminVideos *echo.Group) {
	adminVideos.POST("/parts", partsController.CreatePart)
}
