package controllers

import (
	authController "backend/controllers/auth"
	paginationController "backend/controllers/pagination"
	announcementsDBInteractions "backend/database/announcements"
	anouncementsModel "backend/models/announcements"
	"backend/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

//GetAnnouncementsByYear retrieves list of announcements with the client specified length.
func GetAnnouncementsByYear(c echo.Context) error {
	var year int
	if authController.FetchLoggedInUserAdminStatus(c) {
		year = utils.ConvertToInt(c.QueryParam("year"))
	} else {
		year = authController.FetchLoggedInUserYear(c)
	}
	title := c.QueryParam("title")
	topic := c.QueryParam("topic")
	content := c.QueryParam("content")
	paginationData := paginationController.ExtractPaginationData(c)
	announcements, numberOfAnnouncements := announcementsDBInteractions.
		GetAnnouncementsByYear(paginationData, title, topic, content, year)
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
	year := utils.ConvertToInt(c.FormValue("year"))

	announcement := anouncementsModel.Announcement{
		UserID:  userid,
		Title:   title,
		Topic:   topic,
		Content: content,
		Year:    year,
	}

	announcement.ID = 8
	err := announcementsDBInteractions.CreateAnnouncement(&announcement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (announcement not created), please try again",
		})
	}
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
	year := utils.ConvertToInt(c.FormValue("year"))

	announcement := announcementsDBInteractions.GetAnnouncementByID(announcementID)
	announcement.Title = title
	announcement.Topic = topic
	announcement.Content = content
	announcement.Year = year

	err := announcementsDBInteractions.UpdateAnnouncement(&announcement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (announcement not updated), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":      "Announcement updated successfully",
		"announcement": announcement,
	})
}

// DeleteAnnouncement deletes the announcement that the user selects
func DeleteAnnouncement(c echo.Context) error {
	announcementID := utils.ConvertToUInt(c.FormValue("id"))
	announcement := announcementsDBInteractions.GetAnnouncementByID(announcementID)
	err := announcementsDBInteractions.DeleteAnnouncement(&announcement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Unexpected error occurred (announcement not deleted), please try again",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Announcement deleted successfully",
	})
}
