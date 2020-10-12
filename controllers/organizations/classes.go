package organizations

import (
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
	classes, numberOfClasses := classesDBInteractions.GetAllClasses(paginationData, className, subject, teacherName, year)
	return c.JSON(http.StatusOK, echo.Map{
		"classes": classes,
		"total":   numberOfClasses,
	})
}
