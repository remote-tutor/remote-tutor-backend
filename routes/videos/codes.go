package videos

import (
	codesController "backend/controllers/videos"
	"github.com/labstack/echo"
)


func InitializeCodesRoutes(videos *echo.Group, adminVideos *echo.Group) {
	adminVideos.GET("/codes", codesController.GetCodesByVideo)
	adminVideos.POST("/codes", codesController.GenerateCodes)
}
