package controllers

import (
	authController "backend/controllers/auth"
	announcementsDBInteractions "backend/database/announcements"
	anouncementsModel "backend/models/announcements"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//GetAnnouncements retrieves list of announcements with the client specified length.
func GetAnnouncements(c echo.Context) error {
	title := c.QueryParam("title")
	topic := c.QueryParam("topic")
	content := c.QueryParam("content")
	length := utils.ConvertToInt(c.QueryParam("length"))
	currentPage := utils.ConvertToInt(c.QueryParam("currentPage"))
	announcements, numberOfAnnouncements := announcementsDBInteractions.GetAnnouncements(title, topic, content, length, currentPage)
	return c.JSON(http.StatusOK, echo.Map{
		"announcements": announcements,
		"total":         numberOfAnnouncements,
	})
}

// CreateAnnouncement creates a new announcement based on the user input
func CreateAnnouncement(c echo.Context) error {
	userid := authController.FetchLoggedInUserID(c)

	title := c.FormValue("title")
	topic := c.FormValue("topic")
	content := c.FormValue("content")

	announcement := anouncementsModel.Announcement{
		UserID:  userid,
		Title:   title,
		Topic:   topic,
		Content: content,
	}

	announcementsDBInteractions.CreateAnnouncement(&announcement)
	return c.JSON(http.StatusOK, echo.Map{
		"message":      "Announcement created successfully",
		"announcement": announcement,
	})
}

// UpdateAnnouncement updates the announcement that the user selects
func UpdateAnnouncement(c echo.Context) error {
	announcementID := utils.ConvertToUInt(c.FormValue("id"))
	fmt.Println(announcementID)
	title := c.FormValue("title")
	topic := c.FormValue("topic")
	content := c.FormValue("content")

	announcement := announcementsDBInteractions.GetAnnouncementByID(announcementID)
	announcement.Title = title
	announcement.Topic = topic
	announcement.Content = content

	announcementsDBInteractions.UpdateAnnouncement(&announcement)
	return c.JSON(http.StatusOK, echo.Map{
		"message":      "Announcement updated successfully",
		"announcement": announcement,
	})
}

// DeleteAnnouncement deletes the announcement that the user selects
func DeleteAnnouncement(c echo.Context) error {
	announcementID := utils.ConvertToUInt(c.FormValue("id"))
	announcement := announcementsDBInteractions.GetAnnouncementByID(announcementID)
	announcementsDBInteractions.DeleteAnnouncement(&announcement)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Announcement deleted successfully",
	})
}
