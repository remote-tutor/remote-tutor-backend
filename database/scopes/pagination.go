package scopes

import (
	"backend/utils"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Paginate is a middleware used to paginate every reponse without duplicating the code
func Paginate(c echo.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		queryParams := c.Request().URL.Query()
		sortDesc := utils.ConvertToBoolArray(queryParams["sortDesc[]"])
		sortBy := queryParams["sortBy[]"]
		page := utils.ConvertToInt(queryParams["page"][0])
		itemsPerPage := utils.ConvertToInt(queryParams["itemsPerPage"][0])

		if sortBy != nil {
			for i := 0; i < len(sortBy); i++ {
				if sortDesc[i] {
					db = db.Order(sortBy[i] + " DESC")
				} else {
					db = db.Order(sortBy[i])
				}
			}
		}
		offset := (page - 1) * itemsPerPage
		return db.Offset(offset).Limit(itemsPerPage)
	}
}
