package announcements

import (
	announcementsController "backend/controllers/announcements"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for announcements.
func InitializeRoutes(e *echo.Echo, adminRouter *echo.Group) {
	announcements := e.Group("/announcements", middleware.JWT([]byte("secret")))
	announcements.GET("", announcementsController.GetAnnouncements)

	adminAnnouncements := adminRouter.Group("/announcements")
	adminAnnouncements.POST("", announcementsController.CreateAnnouncement)
	adminAnnouncements.PUT("", announcementsController.UpdateAnnouncement)
	adminAnnouncements.DELETE("", announcementsController.DeleteAnnouncement)
}
