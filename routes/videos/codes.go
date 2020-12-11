package videos

import (
	codesController "backend/controllers/videos"
	"github.com/labstack/echo"
)


func InitializeCodesRoutes(videos *echo.Group, adminVideos *echo.Group) {
	videos.GET("/codes/code", codesController.GetCodeByUserAndVideo)
	videos.PUT("/codes/code", codesController.GrantStudentAccess)

	adminVideos.GET("/codes", codesController.GetCodesByVideo)
	adminVideos.POST("/codes", codesController.GenerateCodes)
	adminVideos.POST("/codes/manual", codesController.AddManualAccess)
}
