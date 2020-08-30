package quizzes

import (
	choiceDBInteractions "backend/database/quizzes"
	choiceModel "backend/models/quizzes"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo"
)

// CreateChoice creates a new choice
func CreateChoice(c echo.Context) error {
	mcqID := utils.ConvertToUInt(c.FormValue("mcqID"))
	text := c.FormValue("text")

	choice := choiceModel.Choice{
		MCQID: mcqID,
		Text:  text,
	}

	choiceDBInteractions.CreateChoice(&choice)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Choice created successfully",
		"choice":  choice,
	})
}

// UpdateChoice updates an existing choice
func UpdateChoice(c echo.Context) error {
	id := utils.ConvertToUInt(c.FormValue("id"))
	text := c.FormValue("text")

	choice := choiceDBInteractions.GetChoiceByID(id)
	choice.Text = text

	choiceDBInteractions.UpdateChoice(&choice)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Choice updates successfully",
		"choice":  choice,
	})
}
