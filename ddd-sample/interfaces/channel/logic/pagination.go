/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    pagination
	@Date    2022/5/16 12:29
	@Desc
*/

package logic

const (
	DefaultPageSize = 20
)

func getPageSizeAndCurrent(reqCurrent, reqPageSize *int) (pageSize, current int) {
	// 默认情况.
	current = 1
	pageSize = DefaultPageSize

	// 当给值的时候.
	if reqPageSize != nil {
		pageSize = *reqPageSize
	}

	if reqCurrent != nil {
		current = *reqCurrent
	}

	return
}
