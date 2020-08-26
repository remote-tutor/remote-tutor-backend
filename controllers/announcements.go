package controllers

import (
	dbInteractions "backend/database"
	md "backend/models"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//GetAnnouncements retrieves the non activated users to view to the admin
func GetAnnouncements(c echo.Context) error {
	title := c.QueryParam("title")
	topic := c.QueryParam("topic")
	content := c.QueryParam("content")
	length := utils.ConvertToInt(c.QueryParam("length"))
	currentPage := utils.ConvertToInt(c.QueryParam("currentPage"))
	announcements, numberOfAnnouncements := dbInteractions.GetAnnouncements(title, topic, content, length, currentPage)
	return c.JSON(http.StatusOK, echo.Map{
		"announcements": announcements,
		"total":         numberOfAnnouncements,
	})
}

// CreateAnnouncement creates a new announcement based on the user input
func CreateAnnouncement(c echo.Context) error {
	userid := fetchLoggedInUserID(c)

	title := c.FormValue("title")
	topic := c.FormValue("topic")
	content := c.FormValue("content")

	announcement := md.Announcement{
		UserID:  userid,
		Title:   title,
		Topic:   topic,
		Content: content,
	}

	dbInteractions.CreateAnnouncement(&announcement)
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

	announcement := dbInteractions.GetAnnouncementById(announcementID)
	announcement.Title = title
	announcement.Topic = topic
	announcement.Content = content

	dbInteractions.SaveAnnouncement(&announcement)
	return c.JSON(http.StatusOK, echo.Map{
		"message":      "Announcement updated successfully",
		"announcement": announcement,
	})
}

// DeleteAnnouncement deletes the announcement that the user selects
func DeleteAnnouncement(c echo.Context) error {
	announcementID := utils.ConvertToUInt(c.FormValue("id"))
	announcement := dbInteractions.GetAnnouncementById(announcementID)
	dbInteractions.DeleteAnnouncement(&announcement)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Announcement deleted successfully",
	})
}
