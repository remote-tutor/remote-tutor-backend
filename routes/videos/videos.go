package videos

import (
	videosController "backend/controllers/videos"
	"github.com/labstack/echo")

func InitializeVideosRoutes(videos *echo.Group, adminVideos *echo.Group) {
	videos.GET("", videosController.GetVideosByClassAndMonthAndYear)
	videos.GET("/video", videosController.GetVideoByID)

	adminVideos.POST("", videosController.CreateVideo)
	adminVideos.PUT("", videosController.UpdateVideo)
	adminVideos.DELETE("", videosController.DeleteVideo)
}