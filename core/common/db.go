package common

import (
	"core/types"

	"gorm.io/gorm"
)

func Paginate(pages types.Pages) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case pages.PageSize > 100:
			pages.PageSize = 100
		case pages.PageSize <= 0:
			pages.PageSize = 10
		}

		offset := (pages.Page - 1) * pages.PageSize
		return db.Offset(offset).Limit(pages.PageSize)
	}
}
