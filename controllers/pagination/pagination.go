package pagination

import (
	dbPagination "backend/database/scopes"
	"backend/utils"
	"github.com/labstack/echo"
)

func ExtractPaginationData(c echo.Context) *dbPagination.PaginationData {
	queryParams := c.Request().URL.Query()
	sortDesc := utils.ConvertToBoolArray(queryParams["sortDesc[]"])
	sortBy := queryParams["sortBy[]"]
	page := utils.ConvertToInt(queryParams["page"][0])
	itemsPerPage := utils.ConvertToInt(queryParams["itemsPerPage"][0])
	return &dbPagination.PaginationData{
		SortDesc:     sortDesc,
		SortBy:       sortBy,
		Page:         page,
		ItemsPerPage: itemsPerPage,
	}
}
