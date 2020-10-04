package videos

import (
	videosController "backend/controllers/videos"
	"github.com/labstack/echo")

func InitializeVideosRoutes(videos *echo.Group, adminVideos *echo.Group) {
	adminVideos.POST("", videosController.CreateVideo)
}