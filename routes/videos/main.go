package videos

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	videos := e.Group("/videos", middleware.JWT([]byte(os.Getenv("JWT_TOKEN"))))
	adminVideos := adminRouter.Group("/videos")

	InitializeVideosRoutes(videos, adminVideos)
	InitializePartsRoutes(videos, adminVideos)
	InitializeWatchesRoutes(videos, adminVideos)
}
