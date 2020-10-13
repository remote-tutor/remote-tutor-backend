package organizations

import (
	controllers "backend/controllers/auth"
	paginationController "backend/controllers/pagination"
	classesDBInteractions "backend/database/organizations"
	"backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

func GetAllClasses(c echo.Context) error {
	paginationData := paginationController.ExtractPaginationData(c)
	subject := c.QueryParam("subject")
	teacherName := c.QueryParam("teacherName")
	className := c.QueryParam("className")
	year := utils.ConvertToInt(c.QueryParam("year"))
	userID := controllers.FetchLoggedInUserID(c)
	classes, numberOfClasses := classesDBInteractions.GetAllClasses(
		paginationData, className, subject, teacherName, year, userID,
	)
	return c.JSON(http.StatusOK, echo.Map{
		"classes": classes,
		"total":   numberOfClasses,
	})
}
