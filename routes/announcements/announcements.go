package announcements

import (
	announcementsController "backend/controllers/announcements"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for announcements.
func InitializeRoutes(e *echo.Echo) {
	announcements := e.Group("/announcements")
	announcements.Use(middleware.JWT([]byte("secret")))
	announcements.GET("", announcementsController.GetAnnouncements)
	announcements.POST("", announcementsController.CreateAnnouncement)
	announcements.PUT("", announcementsController.UpdateAnnouncement)
	announcements.DELETE("", announcementsController.DeleteAnnouncement)
}
