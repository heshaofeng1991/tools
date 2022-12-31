package repository

import (
	inter "combo/internal"
)

func GetPageSizeAndCurrent(reqCurrent, reqPageSize *int) (pageSize, current int) {
	// 默认情况.
	current = 1
	pageSize = inter.DefaultPageSize

	// 当给值的时候.
	if reqPageSize != nil {
		pageSize = *reqPageSize
	}

	if reqCurrent != nil {
		current = *reqCurrent
	}

	return
}
