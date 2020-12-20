package videos

import (
	watchesController "backend/controllers/videos"
	"github.com/labstack/echo"
)

func InitializeWatchesRoutes(videos *echo.Group, adminVideos *echo.Group) {
	videos.GET("/watches/watch", watchesController.GetWatchByUserAndPart)
	adminVideos.GET("/watches/part", watchesController.GetPartWatchesForAllUsers)
	adminVideos.GET("/watches/part/pdf", watchesController.GetPartWatchesPDF)
	videos.POST("/watches", watchesController.CreateUserWatch)
}
