package scopes

import (
	"gorm.io/gorm"
)

type PaginationData struct {
	SortDesc []bool
	SortBy []string
	Page int
	ItemsPerPage int
}

// Paginate is a middleware used to paginate every reponse without duplicating the code
func Paginate(paginationData *PaginationData) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if paginationData.SortBy != nil {
			for i := 0; i < len(paginationData.SortBy); i++ {
				if paginationData.SortDesc[i] {
					db = db.Order(paginationData.SortBy[i] + " DESC")
				} else {
					db = db.Order(paginationData.SortBy[i])
				}
			}
		}
		offset := (paginationData.Page - 1) * paginationData.ItemsPerPage
		return db.Offset(offset).Limit(paginationData.ItemsPerPage)
	}
}
