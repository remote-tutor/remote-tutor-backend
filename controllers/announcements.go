package controllers

import (
	db "backend/database"
	md "backend/models"
	"net/http"
	"github.com/labstack/echo"
)

// GetAnnouncements retrieves the non activated users to view to the admin
func GetAnnouncements(c echo.Context) error {
	announcements := db.GetAnnouncements()
	return c.JSON(http.StatusOK, echo.Map{
		"announcements":      announcements,
	})
}

func CreateAnnouncement(c echo.Context) error {
	userid := uint(1)
	title := c.FormValue("title")
	topic := c.FormValue("topic")
	content := c.FormValue("content")

	announcement := md.Announcement{
		UserID: userid,
		Title: title,
		Topic: topic,
		Content: content,
	}

	db.CreateAnnouncement(&announcement)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Announcement created successfully",
	})
}