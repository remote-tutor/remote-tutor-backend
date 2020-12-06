package videos

import (
	codesController "backend/controllers/videos"
	"github.com/labstack/echo"
)


func InitializeCodesRoutes(videos *echo.Group, adminVideos *echo.Group) {
	adminVideos.POST("/codes", codesController.GenerateCodes)
}
